// handleEncodeLink.go
package server

import ( 
	"net/http"
)

// when encodelink is opened render page with status OK 200
func handleEncodeLink(w http.ResponseWriter, r *http.Request) {
	data := StatusData{
		ResultString: "",
		StatusPhrase: http.StatusText(http.StatusOK), // 200
		StatusCode:   http.StatusOK,
	}
	renderTemplate(w, encodeLinkDir, encodeLinkFile, data)
}