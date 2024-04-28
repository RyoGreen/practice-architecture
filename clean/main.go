package main

import (
	"architecture/clean/controller"
	"log/slog"
	"net/http"
)

func main() {
	// mysqlを使う場合は最初にInit関数を呼び出す
	// if err := mysql.Init(); err != nil {
	// 	slog.Error(err.Error())
	// 	return
	// }
	controller := controller.NewStaffController()
	http.HandleFunc("/", controller.List)
	http.HandleFunc("/get", controller.Get)
	http.HandleFunc("/create", controller.Create)
	http.HandleFunc("/update", controller.Update)
	http.HandleFunc("/delete", controller.Delete)
	slog.Info("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error(err.Error())
		return
	}
}
