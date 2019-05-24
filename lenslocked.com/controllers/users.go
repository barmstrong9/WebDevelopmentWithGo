package controllers

import (
	"net/http"
	"github.com/barmstrong9/WebDevelopmentWithGo/lenslocked.com/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct{
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}