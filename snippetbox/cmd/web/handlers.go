package main


import (
	"fmt"
	"net/http"
	"strconv"
	 "html/template"
	 "log"
)


func (app *application) home(w http.ResponseWriter, r *http.Request){
	//Afegim header
	w.Header().Add("Server","GO")
		
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w,r,err)
	}
    for _ ,snippet := range snippets{
		fmt.Fprintf(w,"%+v",snippet)

	}

    // // Creem slice de cadenes de strings
	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// 	"./ui/html/partials/nav.tmpl",

	// }
    // // fem parse de les files
	// ts,err := template.ParseFiles(files...) //... to pass the contents of the slice files as variadic arguments
	// if err != nil {
	// 	app.logger.Error(err.Error())
	// 	app.serverError(w,r,err)
	// 	return
	// }
    // // les enviem si no hi ha cap error, amb la template base, que invocarà a les altres que fa servir
	// err = ts.ExecuteTemplate(w,"base",nil)

	// if err != nil {
	// 	log.Println(err.Error())
	// 	app.serverError(w,r,err)
	// 	return
	// }
}

func  (app *application) snippetView(w http.ResponseWriter, r *http.Request){
	id,err :=strconv.Atoi(r.PathValue("id"))
	fmt.Println(id)
 
	if err != nil || id < 1 {
		http.NotFound(w,r)
		app.logger.Error(err.Error())

		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		app.serverError(w,r,err)
	}

	 // Creem slice de cadenes de strings
	 files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/view.tmpl",
		"./ui/html/partials/nav.tmpl",

	}
    // fem parse de les files
	ts,err := template.ParseFiles(files...) //... to pass the contents of the slice files as variadic arguments
	if err != nil {
		app.logger.Error(err.Error())
		app.serverError(w,r,err)
		return
	}

	data := templateData{
          Snippet:snippet,
	}
    // les enviem si no hi ha cap error, amb la template base, que invocarà a les altres que fa servir
	err = ts.ExecuteTemplate(w,"base",data)

	if err != nil {
		log.Println(err.Error())
		app.serverError(w,r,err)
		return
	}





    
	fmt.Fprintf(w,"%+v",snippet)
}

func  (app *application) snippetCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a form for creating a snippet..."))

}

func  (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("saving a new snippet.."))

    title := "0 Snail"
	content := "0 snail Climb Mount fugi"
	expires := 7

	id,err := app.snippets.Insert(title,content,expires)

	if err != nil {
		app.serverError(w,r,err)
		return 
	}

	http.Redirect(w,r,fmt.Sprintf("/snippet/view/%d",id),http.StatusSeeOther)


}



