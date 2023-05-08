package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

//Go use interfaces implecity so.
// just by using the method in the interface
// Go will implement the interface an make use of it.
// in this case we use handler interface. which have serveHTTP
// methos in it.

// This class (file) will take care of all route of our API
// We need to register all of them here.

func RegisterControllers() {
	uc := newUserController()
	//now we regisster it and also
	//prive the pattern to make sure the regex in our constructor
	//function returns the correct type.
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
