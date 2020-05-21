package main

import (
	"fmt"
	"log"

	"github.com/gothing/draft"
	"github.com/gothing/draft-demo/auth"
)

const srv = "localhost:2047"

func main() {
	// Создаём API-контейнер
	api := draft.Create(draft.Config{
		DevMode: true,
	})

	// Добавляем «сервисы»
	api.Add(auth.Service)

	// Принтим всё, что получилось
	fmt.Printf("Server: http://%s\n", srv)
	fmt.Printf(" - http://%s/godraft:docs/\n", srv)
	fmt.Printf(" - http://%s/godraft:scheme/\n\n", srv)
	fmt.Println("API:")

	for _, u := range api.URLs() {
		fmt.Printf(" --> http://%s%s\n", srv, u)
	}

	fmt.Println("")

	// Стартуем сервер
	err := api.ListenAndServe(srv)
	if err != nil {
		log.Fatal(err)
	}
}
