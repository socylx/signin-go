package strategy

import (
	"log"
	"signin-go/global/config"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	"signin-go/service/strategy"
)

type recommendGenerateRequest struct {
	Token string `form:"token" binding:"required"`
}

func recommendGenerateOfLaxin(c core.Context) {
	request := new(recommendGenerateRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	if request.Token != config.Server.Token {
		c.AbortWithError(core.Error(code.PermissionError, code.Text(code.PermissionError)))
		return
	}

	generr := strategy.GenerateOfLaxin(c.RequestContext())
	if generr != nil {
		c.AbortWithError(generr)
		return
	}
	log.Println("recommendGenerateOfLaxin success...")
	c.Payload("success")
}

func recommendGenerateOfRenew(c core.Context) {
	request := new(recommendGenerateRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	if request.Token != config.Server.Token {
		c.AbortWithError(core.Error(code.PermissionError, code.Text(code.PermissionError)))
		return
	}
	generr := strategy.GenerateOfRenew(c.RequestContext())
	if generr != nil {
		c.AbortWithError(generr)
		return
	}
	log.Println("recommendGenerateOfRenew success...")
	c.Payload("success")
}
