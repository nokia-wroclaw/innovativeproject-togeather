package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/nokia-wroclaw/innovativeproject-togeather/backend/pkg/core"
)

type lobbyMiddleware struct {
	lobbyService core.LobbyService
	userService core.UserService
}

var CookieUserIDKey = "user-id"
var UserKey = "user"

func authMiddleware(next http.HandlerFunc, m lobbyMiddleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userCookie, err := r.Cookie("user-id")
		if err != nil {
			respondError(w, http.StatusBadRequest,
				fmt.Errorf("unauthorized: user not logged in"))
			return
		}

		if userCookie == nil {
			respondError(w, http.StatusBadRequest,
				errors.New("unauthorized: no user found"))
			return
		}

		userID, err := strconv.Atoi(userCookie.Value)
		if err != nil {
			respondError(w, http.StatusBadRequest,
				errors.New("incorrect cookie value: not a number"))
			return
		}

		user, err := m.userService.Get(ctx, userID)
		if err != nil {
			respondError(w, http.StatusBadRequest,
				fmt.Errorf("error retrieving given user " +
					"from database: %v", userID))
			return
		}

		newCtx := context.WithValue(ctx, UserKey, user)
		next.ServeHTTP(w, r.WithContext(newCtx))
	}
}

//func authMiddleware(m lobbyMiddleware, h http.Handler) http.Handler {
//	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
//		ctx := r.Context()
//
//		userCookie, err := r.Cookie("user-id")
//		if err != nil {
//			h.ServeHTTP(w, r)
//			return
//		}
//
//		if userCookie == nil {
//			respondError(w, http.StatusBadRequest,
//				errors.New("unauthorized: no user found"))
//			return
//		}
//
//		userID, err := strconv.Atoi(userCookie.Value)
//		if err != nil {
//			respondError(w, http.StatusBadRequest,
//				 errors.New("incorrect cookie value: not a number"))
//			return
//		}
//
//		user, err := m.userService.Get(ctx, userID)
//		if err != nil {
//			respondError(w, http.StatusBadRequest,
//				fmt.Errorf("error retrieving given user " +
//					"from database: %v", userID))
//			return
//		}
//
//		newCtx := context.WithValue(ctx, UserKey, user)
//		h.ServeHTTP(w, r.WithContext(newCtx))
//	})
//}
//
func (m *lobbyMiddleware) cookiesMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user-id")
		if err != nil && r.Method == "GET"{
			respondError(w, http.StatusBadRequest, errors.New("user not authorized"))
			return
		} else if err != nil{
			h.ServeHTTP(w, r)
			return
		}

		clientID, err := strconv.Atoi(cookie.Value)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		lobbyID, err := strconv.Atoi(chi.URLParam(r, "lobbyID"))
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		ctx := r.Context()
		exists, err := m.lobbyService.BelongsToLobby(ctx, clientID, lobbyID)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		if (exists && r.Method == "GET") || (!exists && r.Method == "POST"){
			h.ServeHTTP(w, r)
		} else if exists && r.Method == "POST" {
			respondError(w, http.StatusBadRequest, errors.New("given user already belongs to this lobby"))
			return

		} else {
			respondError(w, http.StatusBadRequest, errors.New("given user does not belong to this lobby"))
			return
		}

	})
}
