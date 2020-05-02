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

// AuthLoginResponse — ответ на запрос
type AuthLoginResponse struct {
	// инлайн комментарий с наследованием родительского
	UserID types.UserID `comment:"{super}, авторизованного пользователя"`
}

// InitEndpointScheme - инициализация описания «конца»
func (a *AuthLogin) InitEndpointScheme(s *draft.Scheme) {
	// Проект, которому принадлежит «конец»
	s.Project("auth")

	// Метод доступен всем
	s.Access(draft.Access.All)

	// Базовый метод для всех кейсов
	s.Method(draft.Method.POST)

	// Название метода
	s.Name("«Вход»")

	// URL по которому будет доступен метод
	s.URL("/api/v1/auth/login")

	// Базовые параметры запроса для всех кейсов
	s.Params(AuthLoginParams{
		Login:    types.GenLogin(),
		Password: types.GenPassword(),
	})

	// 200 OK
	s.Case(draft.Status.OK, "Успешная авторизация", func() {
		s.Body(AuthLoginResponse{
			UserID: types.GenUserID(),
		})
	})

	// 403 OK
	s.Case(draft.Status.Denied, "Неправильный Логин или Пароль", func() {
		// Переопределяем базовые параметры запроса
		s.Params(AuthLoginParams{
			Login:    "not-exists-login",
			Password: types.GenPassword(),
		})
	})
}
