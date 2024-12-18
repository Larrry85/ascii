// renderTemplate.go
package server

import ( 
	"html/template"
	"net/http"
	"path/filepath"
)

// renders HTML template
func renderTemplate(w http.ResponseWriter, tmplDir, tmplFile string, data StatusData) {
    tmplPath := filepath.Join(tmplDir, tmplFile) // path to html file

    // reads the contents of the html file
    t, err := template.ParseFiles(tmplPath) // parses it into template
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError) // 500
        return
    }
    // writes parsed template with data to http.ResponseWriter
    err = t.Execute(w, data) // sends the struct data to html placeholders {{ }}
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError) // 500
        return
    }
} // renderTemplate() END
///////////////////////////////////////////////////////////