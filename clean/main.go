package main

import (
	"clean-architecture/controller"
	"clean-architecture/postgres"

	"log/slog"
	"net/http"
)

func main() {
	if err := postgres.Init(); err != nil {
		slog.Error(err.Error())
		return
	}
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
