package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"log"
	"net/http"

	"github.com/go_rest_sample/countries"
)

func main() {
	// db（仮）
	db := gormConnect()
	defer db.Close()

	var allUsers []User
	db.Find(&allUsers)
	fmt.Println(allUsers)

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/countries", countries.GetAllCountries),
		rest.Get("/countries/:code", countries.GetCountry),
		rest.Post("/countries", countries.PostCountry),
		rest.Delete("/countries/:code", countries.DeleteCountry),
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

type User struct {
	ID       int64 `gorm:"primary_key"`
	Username string
}

//func (u *User) TableName() string {
//	return "auth_user"
//}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "go_rest_sample_development"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error()) // deferのみ処理して、強制終了
	}
	return db
}
