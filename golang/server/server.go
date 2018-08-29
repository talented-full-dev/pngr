package server

import (
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/utils"

	"github.com/karlkeefer/pngr/golang/server/login"
	"github.com/karlkeefer/pngr/golang/server/signup"
	"github.com/karlkeefer/pngr/golang/server/verify"

	"net/http"
)

type server struct {
	env *env.Env
	fs  http.Handler
}

func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head, tail := utils.ShiftPath(r.URL.Path)
	if head == "api" {
		r.URL.Path = tail
		srv.ServeAPI(w, r)
	} else {
		srv.fs.ServeHTTP(w, r)
	}
}

func (srv *server) ServeAPI(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	switch head {
	case "login":
		login.ServeHTTP(srv.env, w, r)
	case "signup":
		signup.ServeHTTP(srv.env, w, r)
	case "verify":
		verify.ServeHTTP(srv.env, w, r)
	default:
		errors.Write(w, errors.RouteNotFound)
	}
}

func New() (*server, error) {
	env, err := env.New()
	if err != nil {
		return nil, err
	}

	return &server{
		env: env,
		// built front-end and static files get copied into the docker
		// container during the production build process
		fs: http.FileServer(http.Dir("/root/front")),
	}, nil
}
