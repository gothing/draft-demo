package auth

import (
	"github.com/gothing/draft"
	"github.com/gothing/draft-demo/auth/login"
	"github.com/gothing/draft-demo/auth/logout"
)

// Service — api отвечающее за авторизацию
var Service = draft.Compose(
	"Авторизация",
	login.Endpoint,
	logout.Endpoint,
)
