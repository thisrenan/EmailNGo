package main

import (
	"EmailNGo/internal/contract"
	"EmailNGo/internal/domain/campaign"
	"EmailNGo/internal/infrastructure/database"
	internalerrors "EmailNGo/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type product struct {
	Id   int
	Name string
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaign
		render.DecodeJSON(r.Body, &request)
		id, err := service.Create(request)

		if err != nil {

			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, 500)
				render.JSON(w, r, map[string]string{"error": err.Error()})
			} else {
				render.Status(r, 400)
				render.JSON(w, r, map[string]string{"error": err.Error()})
			}

			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", r)

}

/*
	r.Use(myMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	// localhost:3000?name=Renan     if more than one param u use localhost:3000?name=Renan&anotherparam=X
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")
		println("exec into endpoint")
		if param != "" {
			w.Write([]byte("Hello World! " + param))
		} else {
			w.Write([]byte("test"))
		}

	})
	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte("Hello World! " + param))
	})
	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		obj := map[string]string{"message": "success"}
		b, _ := json.Marshal(obj)
		w.Write(b)
	})
	//same as localhost/json but with render package, so you dont need to put the content-type and convert with json.marshal
	r.Get("/jsonRender", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "success"}
		render.JSON(w, r, obj)
	})
	//PUT or POST is the same result
	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product
		render.DecodeJSON(r.Body, &product)
		product.Id = 5
		render.JSON(w, r, product)
	})
	http.ListenAndServe(":3000", r)
*/

/*
func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("request", r.Method, " url ", r.RequestURI)
		next.ServeHTTP(w, r)
		println("exec after")
	})
}
*/
