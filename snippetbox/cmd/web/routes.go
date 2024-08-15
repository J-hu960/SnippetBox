package main 

import "net/http"

func (app *application) routes() *http.ServeMux{
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	 mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	 mux.HandleFunc("GET /{$}",app.home) // restrict the route only to '/' and nothing more (no wild cards)
	 mux.HandleFunc("GET /snippet/view/{id}",app.snippetView)
	 mux.HandleFunc("GET /snippet/create",app.snippetCreate)
	 mux.HandleFunc("POST /snippet/create",app.snippetCreatePost)


	 return mux
     
}