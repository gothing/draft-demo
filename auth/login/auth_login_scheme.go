package login

import (
	"github.com/gothing/draft"
	"github.com/gothing/draft/types"
)

// AuthLogin — сткрутура описания endpoint
type AuthLogin struct {
	draft.Endpoint
}

// AuthLoginParams — параметры запроса
type AuthLoginParams struct {
	Login    types.Login    `required:"true"`
	Password types.Password `required:"true"`
	Remember bool           `comment:"Запомнить сессию"` // инлайн комментарий
}
