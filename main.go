package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"sync"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/countries", GetAllCountries),
		rest.Get("/countries/:code", GetCountry),
		rest.Post("/countries", PostCountry),
		rest.Delete("/countries/:code", DeleteCountry),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", api.MakeHandler()))

	//api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
	//	w.WriteJson(map[string]string{"Body": "Hello World!"})
	//}))
	//log.Fatal(http.ListenAndServe("127.0.0.1:8080", api.MakeHandler()))
}

type Country struct {
	Code string
	Name string
}

// メモリストアとロック
var store = map[string]*Country{}
var lock = sync.RWMutex{}

// index
func GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	lock.RLock()

	countries := make([]Country, len(store)) // 固定長配列の作成
	i := 0
	for _, country := range store {
		countries[i] = *country
		i++
	}

	lock.RUnlock()
	w.WriteJson(&countries)
}

//show
func GetCountry(w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")

	lock.RLock()
	var country *Country // 初期化
	if store[code] != nil {
		country = &Country{}
		*country = *store[code]
	}
	lock.RUnlock()

	if country == nil {
		rest.NotFound(w, r)
		return
	}

	w.WriteJson(country)
}

// create
func PostCountry(w rest.ResponseWriter, r *rest.Request) {
	country := Country{}
	err := r.DecodeJsonPayload(&country) // requestからbodyのjsonデータをデコードしてstructに割り当て
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if country.Code == "" {
		rest.Error(w, "country code required", 400)
	}
	if country.Name == "" {
		rest.Error(w, "country name required", 400)
		return
	}
	lock.Lock()
	store[country.Code] = &country
	lock.Unlock()
	w.WriteJson(&country)
}

func DeleteCountry(w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")
	lock.Lock()
	delete(store, code)
	lock.Unlock()
	w.WriteHeader(http.StatusOK)
}
