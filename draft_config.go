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
		Rights: []draft.DocAccess{
			{
				ID:   draft.Access.Auth,
				Name: "Auth",
				Extra: []draft.DocAccessExtra{
					{
						Name: "mPOP",
						Headers: struct {
							Cookie string `json:"cookie" required:"true" comment:"mPOP-куки"`
						}{"Mpop=...;"},
						Params: struct {
							Token string `json:"token" required:"true" comment:"mPOP-токен"`
						}{"854724ce05861c2ce336e279039444a9%3A5441407e0..."},
					},
					{
						Name: "OAuth",
						Params: struct {
							AccessToken string `json:"access_token" required:"true" comment:"OAuth-токен"`
						}{"36ee693610a344929218133291cd27ca..."},
					},
				},
			},
		},
	})
}
