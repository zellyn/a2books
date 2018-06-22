package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vektah/gqlgen/handler"
	"github.com/zellyn/transcriber/graph"
)

func main() {
	/*
			utaiie := data.Book{
				ID:               "utaiie",
				Title:            "Understanding the Apple IIe",
				Authors:          []string{"James Fielding Sather"},
				URL:              "https://archive.org/details/Understanding_the_Apple_IIe",
				ISBN:             "0-8359-8019-7",
				ImageURLTemplate: `https://ia600909.us.archive.org/BookReader/BookReaderImages.php?zip=/30/items/Understanding_the_Apple_IIe/Understanding_the_Apple_IIe_jp2.zip&file=Understanding_the_Apple_IIe_jp2/Understanding_the_Apple_IIe_{{printf "%04d" $page}}.jp2&scale={{$scale}}&rotate=0`,
			}

		bb, err := xml.MarshalIndent(utaiie, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		err = ioutil.WriteFile("./bookdata/utaiie/book.xml", bb, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	*/

	if len(os.Args) > 1 {
		return
	}

	app := graph.NewApp("bookdata")
	http.Handle("/", handler.Playground("Transcriber", "/query"))
	http.Handle("/query", handler.GraphQL(graph.MakeExecutableSchema(app)))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
