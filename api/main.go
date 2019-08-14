package main

import (
	"context"
	"encoding/json"
	"fmt"
	"godomicroservice/pb"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	chooseClient := pb.NewChooseServiceClient(conn)

	routes := mux.NewRouter()

	routes.HandleFunc("/", indexHandler).Methods("GET")
	routes.HandleFunc("/choose/{n}/{k}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UFT-8")

		vars := mux.Vars(r)
		n, err := strconv.ParseFloat(vars["n"], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Invalid parameter n")
		}
		k, err := strconv.ParseFloat(vars["k"], 64)
		if err != nil {
			json.NewEncoder(w).Encode("Invalid parameter k")
		}

		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()

		req := &pb.ChooseRequest{N: n, K: k}
		if resp, err := chooseClient.Compute(ctx, req); err == nil {
			msg := fmt.Sprintf("nCk is %f", resp.Result)
			json.NewEncoder(w).Encode(msg)
		} else {
			msg := fmt.Sprintf("Internal server error: %s", err.Error())
			json.NewEncoder(w).Encode(msg)
		}
	}).Methods("GET")

	fmt.Println("Application is running on : 8080 .....")
	err = http.ListenAndServe(":8080", routes)
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode("Server is running")
}
