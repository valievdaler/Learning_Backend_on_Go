package main

import (
	"log"
	"net/http"
)

type api struct{
	addr string
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Get Method"))
}
func (a *api) postUsersHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Post Method"))
}

func main(){
api:=&api{addr:":8080"}
mux:=http.NewServeMux()
srv:=&http.Server{
	Addr:api.addr,
	Handler:mux,
}

mux.HandleFunc("GET /users",api.getUsersHandler)
mux.HandleFunc("POST /users",api.postUsersHandler)
log.Printf("Server is running on %s",api.addr)
if err:=srv.ListenAndServe();
err!=nil{
	log.Fatal(err)
}

}
