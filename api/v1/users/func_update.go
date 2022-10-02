package users

import (
	"signin-go/internal/core"
)

func (h *handler) update(c core.Context) {

	h.userService.Update(c, 3352)

	c.Payload("success")
}
