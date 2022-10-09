package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// DefaultLevel the default log level
	DefaultLevel = zapcore.InfoLevel

	// DefaultTimeLayout the default time layout;
	DefaultTimeLayout = time.RFC3339

	JSON   EncoderType = "JSON"
	NORMAL EncoderType = "NORMAL"
)

type EncoderType string

// Option custom setup config
type Option func(*option)

type option struct {
	level          zapcore.Level
	fields         map[string]string
	file           io.Writer
	timeLayout     string
	disableConsole bool
	encoderType    EncoderType
	levelLog       bool   //分日志级别记录日志
	levelLogDir    string //分日志级别记录日志所在的目录
}

// WithDebugLevel only greater than 'level' will output
func WithDebugLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.DebugLevel
	}
}

// WithInfoLevel only greater than 'level' will output
func WithInfoLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.InfoLevel
	}
}

// WithWarnLevel only greater than 'level' will output
func WithWarnLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.WarnLevel
	}
}

// WithErrorLevel only greater than 'level' will output
func WithErrorLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.ErrorLevel
	}
}

// WithField add some field(s) to log
func WithField(key, value string) Option {
	return func(opt *option) {
		opt.fields[key] = value
	}
}

// WithFileP write log to some file
func WithFileP(file string) Option {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
	if err != nil {
		panic(err)
	}

	return func(opt *option) {
		opt.file = zapcore.Lock(f)
	}
}

// WithTimeLayout custom time format
func WithTimeLayout(timeLayout string) Option {
	return func(opt *option) {
		opt.timeLayout = timeLayout
	}
}

// WithDisableConsole WithEnableConsole write log to os.Stdout or os.Stderr
func WithDisableConsole() Option {
	return func(opt *option) {
		opt.disableConsole = true
	}
}

func WithJSONEncoder() Option {
	return func(opt *option) {
		opt.encoderType = JSON
	}
}

func WithNORMALEncoder() Option {
	return func(opt *option) {
		opt.encoderType = NORMAL
	}
}

func WithLevelLog(fileDir string) Option {
	dir := filepath.Dir(fileDir)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}
	return func(opt *option) {
		opt.levelLog = true
		opt.levelLogDir = fileDir
	}
}

// NewJSONLogger return a json-encoder zap logger,
func NewLogger(opts ...Option) (*zap.Logger, error) {
	opt := &option{level: DefaultLevel, fields: make(map[string]string)}
	for _, f := range opts {
		f(opt)
	}

	timeLayout := DefaultTimeLayout
	if opt.timeLayout != "" {
		timeLayout = opt.timeLayout
	}

	// similar to zap.NewProductionEncoderConfig()
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger", // used by logger.Named(key); optional; useless
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace", // use by zap.AddStacktrace; optional; useless
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(timeLayout))
		},
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
	}

	var encoder zapcore.Encoder
	switch opt.encoderType {
	case JSON:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case NORMAL:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// lowPriority usd by info\debug\warn
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= opt.level && lvl < zapcore.ErrorLevel
	})

	// highPriority usd by error\panic\fatal
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= opt.level && lvl >= zapcore.ErrorLevel
	})

	stdout := zapcore.Lock(os.Stdout) // lock for concurrent safe
	stderr := zapcore.Lock(os.Stderr) // lock for concurrent safe

	core := zapcore.NewTee()

	if !opt.disableConsole {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder,
				zapcore.NewMultiWriteSyncer(stdout),
				lowPriority,
			),
			zapcore.NewCore(encoder,
				zapcore.NewMultiWriteSyncer(stderr),
				highPriority,
			),
		)
	}

	if opt.levelLog {
		debugF, err := os.OpenFile(opt.levelLogDir+"debug.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
		if err != nil {
			panic(err)
		}

		infoF, err := os.OpenFile(opt.levelLogDir+"info.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
		if err != nil {
			panic(err)
		}

		warnF, err := os.OpenFile(opt.levelLogDir+"warn.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
		if err != nil {
			panic(err)
		}

		errorF, err := os.OpenFile(opt.levelLogDir+"error.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
		if err != nil {
			panic(err)
		}

		fatalF, err := os.OpenFile(opt.levelLogDir+"fatal.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
		if err != nil {
			panic(err)
		}

		// DEBUG
		lvlDebug := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return opt.level <= zapcore.DebugLevel && lvl == zapcore.DebugLevel
		})
		// INFO
		lvlInfo := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return opt.level <= zapcore.InfoLevel && lvl == zapcore.InfoLevel
		})
		// WARN
		lvlWarn := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return opt.level <= zapcore.WarnLevel && lvl == zapcore.WarnLevel
		})
		// ERROR
		lvlError := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return opt.level <= zapcore.ErrorLevel && lvl == zapcore.ErrorLevel
		})
		// FATAL
		lvlFatal := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return opt.level <= zapcore.FatalLevel && lvl >= zapcore.FatalLevel
		})
		core = zapcore.NewTee(core,
			zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(debugF)), lvlDebug),
			zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(infoF)), lvlInfo),
			zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(warnF)), lvlWarn),
			zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(errorF)), lvlError),
			zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(fatalF)), lvlFatal),
		)

	} else if opt.file != nil {
		core = zapcore.NewTee(core,
			zapcore.NewCore(encoder,
				zapcore.AddSync(opt.file),
				zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					return lvl >= opt.level
				}),
			),
		)
	}

	logger := zap.New(core,
		zap.AddCaller(),
		zap.ErrorOutput(stderr),
	)

	for key, value := range opt.fields {
		logger = logger.WithOptions(zap.Fields(zapcore.Field{Key: key, Type: zapcore.StringType, String: value}))
	}
	return logger, nil
}

var _ Meta = (*meta)(nil)

// Meta key-value
type Meta interface {
	Key() string
	Value() interface{}
	meta()
}

type meta struct {
	key   string
	value interface{}
}

func (m *meta) Key() string {
	return m.key
}

func (m *meta) Value() interface{} {
	return m.value
}

func (m *meta) meta() {}

// NewMeta create meat
func NewMeta(key string, value interface{}) Meta {
	return &meta{key: key, value: value}
}

// WrapMeta wrap meta to zap fields
func WrapMeta(err error, metas ...Meta) (fields []zap.Field) {
	capacity := len(metas) + 1 // namespace meta
	if err != nil {
		capacity++
	}

	fields = make([]zap.Field, 0, capacity)
	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	fields = append(fields, zap.Namespace("meta"))
	for _, meta := range metas {
		fields = append(fields, zap.Any(meta.Key(), meta.Value()))
	}

	return
}
