package users

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/go_rest_sample/db"
	"net/http"
)

type User struct {
	ID       int64 `gorm:"primary_key"`
	Username string
}

func GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
	db := db.GormConnect()
	defer db.Close()

	var allUsers []User // スライス
	db.Find(&allUsers)  // 結果を詰める（変数の中身を変更するのでポインタ）

	w.WriteHeader(http.StatusOK)
	w.WriteJson(&allUsers)
}
