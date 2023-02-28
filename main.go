package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/machinebox/graphql"
)

var myTemplate *template.Template

func init() {
	myTemplate = template.Must(template.ParseGlob("templates/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	graphqlClient := graphql.NewClient("https://countries.trevorblades.com/")
	graphqlRequest := graphql.NewRequest(`
		{
			countries{
				code,
				name,
				currency
			}
		}
	`)

	var graphqlResponse interface{}
	err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		panic(err)
	}

	// print the graphql response to terminal
	fmt.Printf("Type of %+v", graphqlResponse)
	fmt.Println(graphqlResponse)
	myTemplate.Execute(w, graphqlResponse)
}

func main() {
	fileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fileServer))
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8000", nil)
}
