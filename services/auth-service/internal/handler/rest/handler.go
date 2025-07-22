package rest

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rhythin/bookspot/auth-service/internal/service"
)

type Handler interface {
	Register(w http.ResponseWriter, r *http.Request) (err error)
	Login(w http.ResponseWriter, r *http.Request) (err error)
	Logout(w http.ResponseWriter, r *http.Request) (err error)
	GetUsers(w http.ResponseWriter, r *http.Request) (err error)
	GetUser(w http.ResponseWriter, r *http.Request) (err error)
	UpdateUser(w http.ResponseWriter, r *http.Request) (err error)
	DeleteUser(w http.ResponseWriter, r *http.Request) (err error)
	ForgotPassword(w http.ResponseWriter, r *http.Request) (err error)
	ResetPassword(w http.ResponseWriter, r *http.Request) (err error)
	GetToken(w http.ResponseWriter, r *http.Request) (err error)
	RevokeToken(w http.ResponseWriter, r *http.Request) (err error)
	RefreshToken(w http.ResponseWriter, r *http.Request) (err error)
}

type handler struct {
	service   service.Service
	validator *validator.Validate
}

func New(service service.Service, validator *validator.Validate) Handler {
	return &handler{
		service:   service,
		validator: validator,
	}
}
