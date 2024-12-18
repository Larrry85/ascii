// handleDecodeLink.go
package server

import (
	"net/http"
) 

// when decodelink is opened render page with status OK 200
func handleDecodeLink(w http.ResponseWriter, r *http.Request) {
	data := StatusData{
		ResultString: "",
		StatusPhrase: http.StatusText(http.StatusOK), // 200
		StatusCode:   http.StatusOK,
	}
	renderTemplate(w, decodeLinkDir, decodeLinkFile, data)
}