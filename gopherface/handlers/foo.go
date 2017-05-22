package handlers

import (
	"fmt"
	"net/http"
)

func FooHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("test")
	fooID := r.Context().Value("fooID").(string)
	w.Write([]byte(fooID))

}
