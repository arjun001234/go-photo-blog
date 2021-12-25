package middleware

import (
	"clean-architecture/entity"
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Middleware interface {
	Auth() Adapter
}

type appmiddleware struct {
	userService entity.UserService
}

func NewMiddleware(s entity.UserService) Middleware {
	return &appmiddleware{s}
}

func (ap appmiddleware) Auth() Adapter {
	return func(h httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			u := entity.User{}

			ck, err := r.Cookie("session")

			if err != nil {
				http.Error(w, "User not logged in", http.StatusNotFound)
				return
			}

			s := entity.Session{
				Session: ck.Value,
				User:    &u,
			}

			err = ap.userService.IsUserLoggedIn(&s)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			ctx := context.WithValue(r.Context(), USER, u)

			h(w, r.WithContext(ctx), p)
		}
	}
}

// func (appmiddleware) SetHeaders() Adapter {
// 	return func(h httprouter.Handle) httprouter.Handle {
// 		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 			h(w, r, p)

// 		}
// 	}
// }
