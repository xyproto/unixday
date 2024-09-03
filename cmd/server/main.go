package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/xyproto/unixday"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Current UNIX Day</title>
    <style>
        body {
            background-color: black;
            color: white;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            font-family: Arial, sans-serif;
            font-size: 10vw;
        }
        #countdown {
            text-align: center;
        }
    </style>
</head>
<body>
    <div id="countdown">{{.UnixDay}}</div>
    <script>
        setInterval(function() {
            fetch('/day')
                .then(response => response.text())
                .then(day => {
                    document.getElementById('countdown').innerText = day;
                });
        }, 1000);
    </script>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index").Parse(htmlTemplate)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			UnixDay int64
		}{
			UnixDay: unixday.CurrentUnixDay(),
		}

		tmpl.Execute(w, data)
	})

	http.HandleFunc("/day", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", unixday.CurrentUnixDay())
	})

	port := ":8080"
	fmt.Printf("Server started at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
