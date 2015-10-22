package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
	//Greetings string `json:"greetings"`
}
type Response struct {
	Greetings string `json:"greetings"`
}

func main() {
	ap := httprouter.New()
	ap.GET("/hello:name", func(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
		//hello/:name1
		fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
	})
	ap.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		x := Request{}
		json.NewDecoder(r.Body).Decode(&x)
		x_marshal, _ := json.Marshal(x)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)

		var res Response
		//	var req Request
		if err := json.Unmarshal(x_marshal, &res); err != nil {
			fmt.Println(err)
			return
		}
		res.Greetings = "Hello, " + x.Name + "!"
		// Marshal provided request into JSON Response
		res_marshal, _ := json.Marshal(res)
		fmt.Fprintf(w, "%s", res_marshal)
		/*
			var req Request
			//	var req Request
			if err := json.Unmarshal(x_marshal, &req); err != nil {
				fmt.Println(err)
				return
			}
			req.Greetings = "Hello, " + x.Name + "!"
			// Marshal provided request into JSON Response
			req_marshal, _ := json.Marshal(req.Greetings)

			fmt.Fprintf(w, "%s", req_marshal)
		*/
	})
	http.ListenAndServe("localhost:6666", ap)
}

//$ go run post_demo.go
