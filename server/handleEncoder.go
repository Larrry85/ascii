// handleEncoder.go
package server

import ( 
	"art/oldArtDecoderTask" // decode and encode functions
	"net/http"
)

const ( // encoding page path
	encodeLinkDir  = "encodelink"
	encodeLinkFile = "encodelink.html"
)

// when the encode button is pushed..
func handleEncoder(w http.ResponseWriter, r *http.Request) {
	// default status data
	var data = StatusData{
		ResultString: "",
		StatusPhrase: "",
		StatusCode:   0,
	}
	if r.Method != http.MethodPost { // if request is not POST
		data = createBadRequestStatus() // 400
		renderTemplate(w, encodeLinkDir, encodeLinkFile, data)
	} else { 
		input := r.FormValue("encodeInputString") // user input value from textarea

		encoded, err := oldArtDecoderTask.Encode(input) // input to encode()
		if err != nil {
			data = createBadRequestStatus() // 400 if error
			renderTemplate(w, encodeLinkDir, encodeLinkFile, data)
		} else {
			// struct: converter art, status data
			data = StatusData{
				ResultString: encoded,
				StatusPhrase: http.StatusText(http.StatusAccepted), // 202
				StatusCode:   http.StatusAccepted,
			}
			renderTemplate(w, encodeLinkDir, encodeLinkFile, data)
		}
	}
}   // handleEncoder() END
///////////////////////////////////////////////////////////

