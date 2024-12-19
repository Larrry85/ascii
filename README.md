# Art Interface

This is my third task: Converting ASCII art in webbrowser and styling the page with CSS.


## How To Use

1. Program starts a server:

```
go run main.go
```

or

```
go build main.go
./main
```

2.  Web browser should open automatically, but if not, type in address bar:

```
localhost:8080
```

3. Index.html page opens with status code 200.

*  In index page user can press two links: decode page and encode page.
*  Decode page opens with status code 200. Then user can decode art by pushing the decode button.
*  Encode page opens with status code 200. Then user can encode art by pushing the encode button.
*  Invalid input will print status code 400.
*  When everything goes well program prints the converter art and status code 202.

4. Close the server:

```
Ctrl+C
```


## Art Interface Directory Tree

```
art/
├── decodelink/                  // decode page
│   └── decodelink.html
├── encodelink/                  // encode page
│   └── encodelink.html
├── index/    
│   └── index.html               // index page
├── oldArtDecoderTask/           // decode and encode functions from task two
│   └── oldArtDecoder.go
│   └── oldArtEncoder.go
├── pics/
│   └── artinterfacepic.png      // Header pics
│   └── decodepic.png
|   └── encodepic.png
│── server/ 
│   └── createBadRequest.go      // struct for bad request
│   └── handleDecodeLink.go      // this is called when decode page is opened
│   └── handleDecoder.go         // this is called when the decode button is pushed
│   └── handleEncodeLink.go      // this is called when encode page is opened
│   └── handleEncoder.go         // this is called when the encode button is pushed
│   └── renderTemplate.go        // this will send the converter art and status to html
│   └── server.go                // the server  
|── styles/                      // CSS
|   └── styles.css
└── go.mod
└── (main                        // if you build the program)
└── main.go                      // start the program with main.go
└── README.md                
```

### Coder

Laura Levistö 5/24
