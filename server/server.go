// server.go
package server

import (
	"fmt"
	"net/http"
	"os"
	"os/exec" // these are for opening..
	"runtime" // ..the web browser windown
)

// struct: converter art, status data
type StatusData struct {
	ResultString string
	StatusPhrase string
	StatusCode int
}

// the server
func Server() {
	// Registers different routes and their respective handlers
	http.HandleFunc("/", handleMain) // sets up a route for the root URL ("/")
	// When a request is made to the root URL, the handleMain function will be called

	http.HandleFunc("/decoder", handleDecoder) // route for /decoder -> call handleDecoder
	http.HandleFunc("/decodelink/decodelink.html", handleDecodeLink)//-> call handleDecodeLink
	http.HandleFunc("/encoder", handleEncoder) // route for /encoder -> call handleEncoder
	http.HandleFunc("/encodelink/encodelink.html", handleEncodeLink) //-> call handleEncodeLink

	http.HandleFunc("/styles.css", handleCSS)  // route for /styles.css -> call handleCSS

	// Serve the static files (html, css, images)
	// when a request is made, it will serve index.html from index directory
	// removes the "/index/" prefix from the URL before serving the files, making it easier to organize and access static files
	http.Handle("/index/", http.StripPrefix("/index/", http.FileServer(http.Dir("index"))))
	http.Handle("/decodelink/", http.StripPrefix("/decodelink/", http.FileServer(http.Dir("decodelink"))))
	http.Handle("/encodelink/", http.StripPrefix("/encodelink/", http.FileServer(http.Dir("encodelink"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/pics/", http.StripPrefix("/pics/", http.FileServer(http.Dir("pics"))))

	fmt.Println("Server listening on port 8080...")

	// Open a web browser
	openBrowser("http://localhost:8080") // send URL to openBrowser()

	err := http.ListenAndServe(":8080", nil) // starts server on port 8080
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

} // Server() END
///////////////////////////////////////////////////////////

// handleCSS is called when a request is made to /styles.css
func handleCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css") // content type header to text/css
	// Serve the CSS file using http.ServeFile()
	http.ServeFile(w, r, "/styles/styles.css") // path to CSS file

} // handleCSS() END
///////////////////////////////////////////////////////////

// handleMain is called when a request is made to the root URL "/"
func handleMain(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet { // if request is GET
		w.WriteHeader(http.StatusOK)
		data := StatusData{
			StatusPhrase: http.StatusText(http.StatusOK), // 200
			StatusCode:   http.StatusOK,
		}
		renderTemplate(w, "index", "index.html", data)
		return
	}
	// If the request method is not GET, return Method Not Allowed
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // 405

} // handleMain() END
///////////////////////////////////////////////////////////

// openBrowser opens the default web browser with the specified URL
func openBrowser(url string) { // takes a URL string as argument
	var err error

	switch runtime.GOOS { // runtime.GOOS to determine the operating system
	case "linux":
		err = exec.Command("xdg-open", url).Start() // open the default web browser
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default: // if not recognized, it returns an error
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil { // in case of an error it prints error message
		fmt.Println("Failed to open web browser:", err)
	}
} // openBrowser() END
///////////////////////////////////////////////////////////
