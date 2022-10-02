package alert

import (
	"signin-go/global/config"
	"signin-go/global/logger"
	"signin-go/internal/errors"
	"signin-go/internal/mail"
	"signin-go/internal/proposal"

	"go.uber.org/zap"
)

// NotifyHandler 告警通知
func NotifyHandler() func(msg *proposal.AlertMessage) {

	return func(msg *proposal.AlertMessage) {
		if config.Notify.Host == "" || config.Notify.Port == 0 || config.Notify.User == "" || config.Notify.Pass == "" || config.Notify.To == "" {
			logger.Logger.Error("Mail config error")
			return
		}

		subject, body, err := newHTMLEmail(
			msg.Method,
			msg.HOST,
			msg.URI,
			msg.TraceID,
			msg.ErrorMessage,
			msg.ErrorStack,
		)
		if err != nil {
			logger.Logger.Error("email template error", zap.Error(err))
			return
		}

		options := &mail.Options{
			MailHost: config.Notify.Host,
			MailPort: config.Notify.Port,
			MailUser: config.Notify.User,
			MailPass: config.Notify.Pass,
			MailTo:   config.Notify.To,
			Subject:  subject,
			Body:     body,
		}
		if err := mail.Send(options); err != nil {
			logger.Logger.Error("发送告警通知邮件失败", zap.Error(errors.WithStack(err)))
		}
	}
}
