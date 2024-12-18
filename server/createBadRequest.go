// createBadRequest.go
package server

import ( 
	"net/http"
)

// creates struct for "bad request"
func createBadRequestStatus() StatusData {
	return StatusData{
		ResultString: "",
		StatusPhrase: http.StatusText(http.StatusBadRequest), // 400
		StatusCode:   http.StatusBadRequest,
	}
}// createBadRequestStatus() END
///////////////////////////////////////////////////////////