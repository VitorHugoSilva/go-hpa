package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x := 0.0001
		for i := 0; i <= 1; i++ {
			x = Calc(x)
		}
		fmt.Println("Code.education Rocks!")
    })

	http.ListenAndServe(":8000", nil);
}