package logout

import (
	"github.com/gothing/draft"
	"github.com/gothing/draft/types"
)

// AuthLogout — сткрутура описания endpoint
type AuthLogout struct {
	draft.Endpoint
}

// AuthLogoutParams — параметры запроса
type AuthLogoutParams struct {
}

// AuthLogoutResponse — ответ на запрос
type AuthLogoutResponse struct {
	UserID types.UserID
}

// InitEndpointScheme - инициализация описания «конца»
func (a *AuthLogout) InitEndpointScheme(s *draft.Scheme) {
	// Проект, которому принадлежит «конец»
	s.Project("auth")

	// Только для авторизованного юзера
	s.Access(draft.Access.Auth)

	// Базовый метод для всех кейсов
	s.Method(draft.Method.POST)

	// Название метода
	s.Name("«Выход»")

	// URL по которому будет доступен метод
	s.URL("/api/v1/auth/logout")

	// 200 OK
	s.Case(draft.Status.OK, "Успех", func() {
		s.Body(AuthLogoutResponse{
			UserID: types.GenUserID(),
		})
	})
}
