package controllers

import (
	"fmt"
	"net/http"

	"github.com/barmstrong9/WebDevelopmentWithGo/lenslocked.com/views"
)

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Email is", form.Password)
}
