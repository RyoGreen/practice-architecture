package controller

import (
	"clean-architecture/controller/in"
	"clean-architecture/repository"
	"clean-architecture/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type StaffController interface {
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type staffController struct {
	staffUseCase usecase.StaffUsecase
}

func NewStaffController() StaffController {
	return &staffController{
		staffUseCase: usecase.NewStaffUseCase(),
	}
}

func (c *staffController) List(w http.ResponseWriter, r *http.Request) {
	staffs, err := c.staffUseCase.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(staffs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *staffController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest)
		return
	}
	staff, err := c.staffUseCase.Get(idInt)
	if err != nil {
		if errors.Is(err, repository.ErrStaffNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(staff); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *staffController) Create(w http.ResponseWriter, r *http.Request) {
	var request in.StaffRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	staff, err := c.staffUseCase.Create(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(staff); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *staffController) Update(w http.ResponseWriter, r *http.Request) {
	var request in.StaffRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	staff, err := c.staffUseCase.Update(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(staff); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *staffController) Delete(w http.ResponseWriter, r *http.Request) {
	var request in.DeleteStaffRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.staffUseCase.Delete(&request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
