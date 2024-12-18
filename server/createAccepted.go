// createAccepted.go
package server

import (
	"net/http"
)

// creates a status struct for "accepted" for printing the status
func createAcceptedStatus() StatusData {
	return StatusData{
		StatusPhrase: http.StatusText(http.StatusAccepted), // 202
		StatusCode:   http.StatusAccepted,
	}
}// createAcceptedStatus() END
///////////////////////////////////////////////////////////