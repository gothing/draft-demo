package main

import "github.com/gothing/draft"

func init() {
	// Инициализируем GODRAFT :: DOC
	draft.SetupDoc(draft.DocConfig{
		FrontURL: "https://gothing.github.io/draft-front/",
	})
}
