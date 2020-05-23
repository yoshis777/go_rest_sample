package main

import (
	"github.com/ant0ine/go-json-rest/rest"

	"log"
	"net/http"

	// 仮）ここらへんはmodelsとかになるはず
	"github.com/go_rest_sample/countries"
	"github.com/go_rest_sample/users"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	// routes
	router, err := rest.MakeRouter(
		// メモリストア
		rest.Get("/countries", countries.GetAllCountries),
		rest.Get("/countries/:code", countries.GetCountry),
		rest.Post("/countries", countries.PostCountry),
		rest.Delete("/countries/:code", countries.DeleteCountry),

		// DBストア
		rest.Get("/users", users.GetAllUsers),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", api.MakeHandler()))
}
