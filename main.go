package main

import (
	"attendance/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"net/http"
)



var  client *mongo.Client

var employees []model.Employee

func TheDayListEndpoint(response http.ResponseWriter, request *http.Request)  {

	response.Header().Add("content-type", "application/json")
	json.NewDecoder(request.Body).Decode(&employees)
	collection:=client.Database("employers").Collection("employee")
	ctx, _:=context.WithTimeout(context.Background(), 10*time.Second)
	result, _:=collection.InsertOne(ctx, employees)
	json.NewEncoder(response).Encode(result)
}
func Hello(w http.ResponseWriter, r *http.Request)  {
	err:=json.NewDecoder(r.Body).Decode(&employees)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "` +err.Error()+`"}`))

}

func main()  {
	fmt.Println("Server starting...")
	ctx, _:=context.WithTimeout(context.Background(), 10*time.Second)
	client, _=mongo.Connect(ctx)
	router:=mux.NewRouter()
	router.HandleFunc("/", Hello)


	http.ListenAndServe(":9999", router)

}
