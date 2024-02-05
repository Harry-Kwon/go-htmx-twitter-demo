package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Starting server...")
	fmt.Println(quote.Go())

	// read index.html as a string and serve on "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := os.ReadFile("index.html")
		fmt.Fprint(w, string(bytes))
	})

	var quotes [4]string
	quotes[0] = quote.Glass()
	quotes[1] = quote.Hello()
	quotes[2] = quote.Go()
	quotes[3] = quote.Opt()

	http.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		content := quotes[rand.Intn(4)]
		component := post(content)
		fmt.Println("New post with content:", content)
		component.Render(context.Background(), w)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		content := r.Form.Get("content")
		fmt.Println("New post with content:", content)

		component := post(content)
		component.Render(context.Background(), w)
	})

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
