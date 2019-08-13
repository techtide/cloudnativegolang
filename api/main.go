package main

import (
	"encoding/json"
	"godomicroservice/pb"
	"net/http"

	"github.com/gorilla/mux"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("choose-service:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	chooseClient := pb.NewChooseServiceClient(conn)

	routes := mux.NewRouter()
	routes.HandleFunc("/", indexHandler).Methods("GET")

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode("Server is running")
}
