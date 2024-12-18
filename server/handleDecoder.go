// handleDecoder.go
package server

import (
	"art/oldArtDecoderTask" // decode and encode functions
	"net/http"
)

const ( // decoding page path
	decodeLinkDir  = "decodelink"
	decodeLinkFile = "decodelink.html"
)

// when the decode button is pushed..
func handleDecoder(w http.ResponseWriter, r *http.Request) {
	// default status data
	var data = StatusData{
		ResultString: "", 
		StatusPhrase: "",
		StatusCode:   0,
	}
	if r.Method != http.MethodPost { // if request is not POST
		data = createBadRequestStatus() // 400
		renderTemplate(w, decodeLinkDir, decodeLinkFile, data)
	} else { 
		input := r.FormValue("decodeInputString") // user input value from textarea

		decoded, err := oldArtDecoderTask.Decode(input) // input to decode()
		if err != nil {
			data = createBadRequestStatus() // 400 if error
			renderTemplate(w, decodeLinkDir, decodeLinkFile, data)
		} else {
			// struct: converter art, status data
			data = StatusData{
				ResultString: decoded,
				StatusPhrase: http.StatusText(http.StatusAccepted), // 202
				StatusCode:   http.StatusAccepted,
			}
			renderTemplate(w, decodeLinkDir, decodeLinkFile, data)
		}
	}
}  // handleDecoder() END
///////////////////////////////////////////////////////////
