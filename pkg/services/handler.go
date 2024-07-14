package services

import "net/http"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := &TemplateData{}
	app.render(w, r, "index.page.html", data)
}
