package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	controllerHttp "refactoring/internal/controller/http"
	"refactoring/internal/repo"
	"refactoring/internal/usecase"
	"time"
)

func Run() error {
	userRepoJSON := repo.NewUserRepoJSON(`users.json`)

	userUsecase := usecase.NewUserUsecase(userRepoJSON)

	userHandler := controllerHttp.NewUserHandler(userUsecase)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", userHandler.SearchUsersHandler)
				r.Post("/", userHandler.CreateUserHandler)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", userHandler.GetUserHandler)
					r.Patch("/", userHandler.UpdateUserHandler)
					r.Delete("/", userHandler.DeleteUserHandler)
				})
			})
		})
	})

	if err := http.ListenAndServe(":3333", r); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

	return nil
}
