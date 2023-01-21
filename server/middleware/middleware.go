package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/snehadeep-wagh/go-todo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClient() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file!")
	}

	url := os.Getenv("MONGO_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	client.Connect(ctx)

	return client
}

var Client = getClient()
var mongo_db = os.Getenv("MONGO_DB")
var mongo_col = os.Getenv("MONGO_COLLECTION")
var taskCollection = Client.Database(mongo_db).Collection(mongo_col)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	cur, err := taskCollection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		log.Fatal("Problem in getting the documents!")
		return
	}

	var listOfTasks []primitive.M

	for cur.Next(ctx) {
		var item bson.M
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal("Problem in decoding the data!")
			return
		}

		listOfTasks = append(listOfTasks, item)
	}

	cur.Close(ctx)

	// send data to the user
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listOfTasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	ctx := context.Background()

	var t model.ToDoStruct
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	t.ID = primitive.NewObjectID()
	t.Status = false

	_, err = taskCollection.InsertOne(ctx, t)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", id}}
	update := bson.M{"$set": bson.M{"status": true}}

	_, err = taskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", false}}}}

	_, err = taskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	res, err := taskCollection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Delete count: " + string(res.DeletedCount))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.DeletedCount)
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", id}}
	_, err = taskCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
