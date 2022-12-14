package core

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"

	"gsteps-go/global/config"
	"gsteps-go/global/logger"
	"gsteps-go/global/mysql"
	"gsteps-go/global/time"
	"gsteps-go/internal/code"
	"gsteps-go/internal/proposal"
	"gsteps-go/internal/trace"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	cors "github.com/rs/cors/wrapper/gin"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/time/rate"
)

type Option func(*option)

type option struct {
	disablePProf      bool
	disableSwagger    bool
	disablePrometheus bool
	enableCors        bool
	enableRate        bool
	alertNotify       proposal.NotifyHandler
	recordHandler     proposal.RecordHandler
}

// WithDisablePProf 禁用 pprof
func WithDisablePProf() Option {
	return func(opt *option) {
		opt.disablePProf = true
	}
}

// WithDisableSwagger 禁用 swagger
func WithDisableSwagger() Option {
	return func(opt *option) {
		opt.disableSwagger = true
	}
}

// WithDisablePrometheus 禁用prometheus
func WithDisablePrometheus() Option {
	return func(opt *option) {
		opt.disablePrometheus = true
	}
}

// WithAlertNotify 设置告警通知
func WithAlertNotify(notifyHandler proposal.NotifyHandler) Option {
	return func(opt *option) {
		opt.alertNotify = notifyHandler
	}
}

// WithRecordMetrics 设置记录接口指标
func WithRecordMetrics(recordHandler proposal.RecordHandler) Option {
	return func(opt *option) {
		opt.recordHandler = recordHandler
	}
}

// WithEnableCors 设置支持跨域
func WithEnableCors() Option {
	return func(opt *option) {
		opt.enableCors = true
	}
}

// WithEnableRate 设置支持限流
func WithEnableRate() Option {
	return func(opt *option) {
		opt.enableRate = true
	}
}

// DisableTraceLog 禁止记录日志
func DisableTraceLog(ctx Context) {
	ctx.disableTrace()
}

// DisableRecordMetrics 禁止记录指标
func DisableRecordMetrics(ctx Context) {
	ctx.disableRecordMetrics()
}

// AliasForRecordMetrics 对请求路径起个别名，用于记录指标。
// 如：Get /user/:username 这样的路径，因为 username 会有非常多的情况，这样记录指标非常不友好。
func AliasForRecordMetrics(path string) HandlerFunc {
	return func(ctx Context) {
		ctx.setAlias(path)
	}
}

// RouterGroup 包装gin的RouterGroup
type RouterGroup interface {
	Group(string, ...HandlerFunc) RouterGroup
	IRoutes
}

var _ IRoutes = (*router)(nil)

// IRoutes 包装gin的IRoutes
type IRoutes interface {
	Any(string, ...HandlerFunc)
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	OPTIONS(string, ...HandlerFunc)
	HEAD(string, ...HandlerFunc)
}

type router struct {
	group *gin.RouterGroup
}

func (r *router) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	group := r.group.Group(relativePath, wrapHandlers(handlers...)...)
	return &router{group: group}
}

func (r *router) Any(relativePath string, handlers ...HandlerFunc) {
	r.group.Any(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) GET(relativePath string, handlers ...HandlerFunc) {
	r.group.GET(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) POST(relativePath string, handlers ...HandlerFunc) {
	r.group.POST(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) DELETE(relativePath string, handlers ...HandlerFunc) {
	r.group.DELETE(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PATCH(relativePath string, handlers ...HandlerFunc) {
	r.group.PATCH(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PUT(relativePath string, handlers ...HandlerFunc) {
	r.group.PUT(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	r.group.OPTIONS(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) HEAD(relativePath string, handlers ...HandlerFunc) {
	r.group.HEAD(relativePath, wrapHandlers(handlers...)...)
}

func wrapHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	funcs := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handler := handler
		funcs[i] = func(c *gin.Context) {
			ctx := newContext(c)
			defer releaseContext(ctx)

			handler(ctx)
		}
	}

	return funcs
}

var _ Mux = (*mux)(nil)

// Mux http mux
type Mux interface {
	http.Handler
	Group(relativePath string, handlers ...HandlerFunc) RouterGroup
}

type mux struct {
	engine *gin.Engine
}

func (m *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.engine.ServeHTTP(w, req)
}

func (m *mux) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	return &router{
		group: m.engine.Group(relativePath, wrapHandlers(handlers...)...),
	}
}

func New(options ...Option) Mux {
	gin.SetMode(config.Server.Mode)

	mux := &mux{
		engine: gin.New(),
	}

	// withoutTracePaths 这些请求，默认不记录日志
	withoutTracePaths := map[string]bool{
		"/metrics": true,

		"/debug/pprof/":             true,
		"/debug/pprof/cmdline":      true,
		"/debug/pprof/profile":      true,
		"/debug/pprof/symbol":       true,
		"/debug/pprof/trace":        true,
		"/debug/pprof/allocs":       true,
		"/debug/pprof/block":        true,
		"/debug/pprof/goroutine":    true,
		"/debug/pprof/heap":         true,
		"/debug/pprof/mutex":        true,
		"/debug/pprof/threadcreate": true,

		"/favicon.ico": true,

		"/system/health": true,
	}

	opt := new(option)
	for _, f := range options {
		f(opt)
	}

	if !opt.disablePProf {
		if !config.Server.IsRelease {
			pprof.Register(mux.engine) // register pprof to gin
		}
	}

	// if !opt.disableSwagger {
	// 	if !config.Server.IsRelease {
	// 		mux.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // register swagger
	// 	}
	// }

	if !opt.disablePrometheus {
		mux.engine.GET("/metrics", gin.WrapH(promhttp.Handler())) // register prometheus
	}

	if opt.enableCors {
		mux.engine.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
			},
			AllowedHeaders:     []string{"*"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
		}))
	}

	// recover两次，防止处理时发生panic，尤其是在OnPanicNotify中。
	mux.engine.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())
				log.Println(err)
				log.Println(stack)
				logger.Logger.Error(
					"gsteps-go panic",
					zap.String("panic", fmt.Sprintf("%+v", err)),
					zap.String("stack", stack),
				)
				ctx.JSON(
					http.StatusOK,
					&Response{
						Code:    code.ServerError,
						Message: fmt.Sprintf("%v", err),
					},
				)
			}
		}()

		ctx.Next()
	})

	mux.engine.Use(func(ctx *gin.Context) {

		if ctx.Writer.Status() == http.StatusNotFound {
			return
		}

		ts := time.Now()

		context := newContext(ctx)
		defer releaseContext(context)

		context.init()
		context.setDB(mysql.DB.WithContext(ctx))
		context.ableRecordMetrics()

		if !withoutTracePaths[ctx.Request.URL.Path] {
			if traceId := context.GetHeader(trace.Header); traceId != "" {
				context.setTrace(trace.New(traceId))
			} else {
				context.setTrace(trace.New(""))
			}
		}
		context.setLogger(logger.Logger.WithOptions(zap.Fields(zapcore.Field{Key: "traceID", Type: zapcore.StringType, String: context.Trace().ID()})))

		defer func() {
			var (
				response        interface{}
				businessCode    int
				businessCodeMsg string
				abortErr        error
				traceId         string
				graphResponse   interface{}
			)

			if ct := context.Trace(); ct != nil {
				context.SetHeader(trace.Header, ct.ID())
				traceId = ct.ID()
			}

			// region 发生 Panic 异常发送告警提醒
			if err := recover(); err != nil {
				stackInfo := string(debug.Stack())
				log.Println(err)
				log.Println(stackInfo)
				logger.Logger.Error(
					"gsteps-go panic",
					zap.String("panic", fmt.Sprintf("%+v", err)),
					zap.String("stack", stackInfo),
				)
				context.AbortWithError(Error(
					code.ServerError,
					code.Text(code.ServerError)),
				)

				// if notifyHandler := opt.alertNotify; notifyHandler != nil {
				// 	notifyHandler(&proposal.AlertMessage{
				// 		ProjectName:  config.Server.Name,
				// 		Env:          config.Server.Mode,
				// 		TraceID:      traceId,
				// 		HOST:         context.Host(),
				// 		URI:          context.URI(),
				// 		Method:       context.Method(),
				// 		ErrorMessage: err,
				// 		ErrorStack:   stackInfo,
				// 		Timestamp:    time.Now(),
				// 	})
				// }
			}
			// endregion

			// region 发生错误，进行返回
			if ctx.IsAborted() {
				for i := range ctx.Errors {
					multierr.AppendInto(&abortErr, ctx.Errors[i])
				}

				if err := context.abortError(); err != nil { // customer err
					// 判断是否需要发送告警通知
					// if err.IsAlert() {
					// 	if notifyHandler := opt.alertNotify; notifyHandler != nil {
					// 		notifyHandler(&proposal.AlertMessage{
					// 			ProjectName:  config.Server.Mode,
					// 			Env:          config.Server.Mode,
					// 			TraceID:      traceId,
					// 			HOST:         context.Host(),
					// 			URI:          context.URI(),
					// 			Method:       context.Method(),
					// 			ErrorMessage: err.Message(),
					// 			ErrorStack:   fmt.Sprintf("%+v", err.StackError()),
					// 			Timestamp:    time.Now(),
					// 		})
					// 	}
					// }

					multierr.AppendInto(&abortErr, err.StackError())
					ctx.JSON(
						http.StatusOK,
						&Response{
							Code:    err.BusinessCode(),
							Message: err.Message(),
						},
					)
				}
			}
			// endregion

			// region 正确返回
			response = context.getPayload()
			if response != nil {
				ctx.JSON(
					http.StatusOK,
					&Response{
						Code: 0,
						Res:  response,
					},
				)
			}
			// endregion

			// region 记录指标
			if opt.recordHandler != nil && context.isRecordMetrics() {
				path := context.Path()
				if alias := context.Alias(); alias != "" {
					path = alias
				}

				opt.recordHandler(&proposal.MetricsMessage{
					ProjectName:  config.Server.Name,
					Env:          config.Server.Mode,
					TraceID:      traceId,
					HOST:         context.Host(),
					Path:         path,
					Method:       context.Method(),
					HTTPCode:     ctx.Writer.Status(),
					BusinessCode: businessCode,
					CostSeconds:  time.Since(ts).Seconds(),
					IsSuccess:    !ctx.IsAborted() && (ctx.Writer.Status() == http.StatusOK),
				})
			}
			// endregion

			// region 记录日志
			var t *trace.Trace
			if x := context.Trace(); x != nil {
				t = x.(*trace.Trace)
			} else {
				return
			}

			decodedURL, _ := url.QueryUnescape(ctx.Request.URL.RequestURI())

			// ctx.Request.Header，精简 Header 参数
			traceHeader := map[string]string{
				"Content-Type": ctx.GetHeader("Content-Type"),
				// configs.HeaderLoginToken:    ctx.GetHeader(configs.HeaderLoginToken),
				// configs.HeaderSignToken:     ctx.GetHeader(configs.HeaderSignToken),
				// configs.HeaderSignTokenDate: ctx.GetHeader(configs.HeaderSignTokenDate),
			}

			t.WithRequest(&trace.Request{
				TTL:        "un-limit",
				Method:     ctx.Request.Method,
				DecodedURL: decodedURL,
				Header:     traceHeader,
				Body:       string(context.RawData()),
			})

			var responseBody interface{}

			if response != nil {
				responseBody = response
			}

			graphResponse = context.getGraphPayload()
			if graphResponse != nil {
				responseBody = graphResponse
			}

			t.WithResponse(&trace.Response{
				Header:          ctx.Writer.Header(),
				HttpCode:        ctx.Writer.Status(),
				HttpCodeMsg:     http.StatusText(ctx.Writer.Status()),
				BusinessCode:    businessCode,
				BusinessCodeMsg: businessCodeMsg,
				Body:            responseBody,
				CostSeconds:     time.Since(ts).Seconds(),
			})

			t.Success = !ctx.IsAborted() && (ctx.Writer.Status() == http.StatusOK)
			t.CostSeconds = time.Since(ts).Seconds()

			logger.Logger.Info("trace-log",
				zap.Any("method", ctx.Request.Method),
				zap.Any("path", decodedURL),
				zap.Any("http_code", ctx.Writer.Status()),
				zap.Any("business_code", businessCode),
				zap.Any("success", t.Success),
				zap.Any("cost_seconds", t.CostSeconds),
				zap.Any("trace_id", t.Identifier),
				zap.Any("trace_info", t),
				zap.Error(abortErr),
			)
			// endregion
		}()

		ctx.Next()
	})

	if opt.enableRate {
		limiter := rate.NewLimiter(rate.Every(time.Second*1), config.Server.MaxRequestsPerSecond)
		mux.engine.Use(func(ctx *gin.Context) {
			context := newContext(ctx)
			defer releaseContext(context)

			if !limiter.Allow() {
				context.AbortWithError(Error(
					code.TooManyRequests,
					code.Text(code.TooManyRequests)),
				)
				return
			}

			ctx.Next()
		})
	}

	mux.engine.NoMethod(wrapHandlers(DisableTraceLog)...)
	mux.engine.NoRoute(wrapHandlers(DisableTraceLog)...)

	system := mux.Group("/system")
	{
		// 健康检查
		system.GET("/health", func(ctx Context) {
			resp := &struct {
				Timestamp   time.Time `json:"timestamp"`
				Environment string    `json:"environment"`
				Host        string    `json:"host"`
				Status      string    `json:"status"`
			}{
				Timestamp:   time.Now(),
				Environment: config.Server.Mode,
				Host:        ctx.Host(),
				Status:      "ok",
			}
			ctx.Payload(resp)
		})
	}

	return mux
}
