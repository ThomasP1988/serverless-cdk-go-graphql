// +build !lambda

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"motif/shared/config"
	common "motif/shared/repositories"
	"net/http"
)

var conf *config.Config = config.GetConfig(&config.DEV)

var _ error = common.SetAWSConfig(&conf.Region, &conf.LocalProfile)

func main() {

	schema, err := GetGraphQLSchema()

	if err != nil {
		println("err", err.Error())
		errorAPI := errors.New("error in the server GraphQL Schema")
		panic(errorAPI)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		println("GRAPHQL Request")

		var query string

		switch r.Method {
		case http.MethodGet:
			query = r.URL.Query().Get("query")
		case http.MethodPost:

			// Try to decode the request body into the struct. If there is an error,
			// respond to the client with the error message and a 400 status code.
			bdy, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(errors.New("can't read body").Error())
			}
			query = string(bdy)
		case http.MethodPut:
			fallthrough
		case http.MethodDelete:
			fallthrough
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
			panic(errors.New("only POST/GET accepted").Error())
		}

		result, err := HandleGraphQL(query, &schema)

		if err != nil {
			println("err HandleGraphQL", err)
		}
		println("response", string(result))
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		// json.NewEncoder(w).Encode(string(result))

	})
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Access the graphql app via insomnia/postman at 'http://localhost:8080/graphql'")

	http.ListenAndServe(":8080", nil)
}
