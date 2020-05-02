package main

import "github.com/gothing/draft"

func init() {
	// Инициализируем GODRAFT :: DOC
	draft.SetupDoc(draft.DocConfig{
		FrontURL:    "https://gothing.github.io/draft-front/",
		ActiveGroup: "demo",
		Groups: []draft.DocGroup{
			{"demo", "Demo", []string{"http://localhost:2047/godraft:scheme/"}},
		},
		Projects: []draft.DocProject{
			{
				ID:      "auth",
				Name:    "Auth",
				Host:    "auth.mail.ru",
				HostRC:  "test.auth.mail.ru",
				HostDEV: "auth.devmail.ru",
			},
		},
	})
}
