package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var ctx context.Context
var err error

var client *mongo.Client

func init() {

	recipes := make([]Recipe, 0)

	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	var listOfRecipes []interface{}

	for _, recipe := range recipes {

		listOfRecipes = append(listOfRecipes, recipe)
	}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	insertManyResult, err := collection.InsertMany(
		ctx, listOfRecipes)

	if err != nil {

		log.Fatal(err)
	}

	log.Println("Inserted recipes: ",

		len(insertManyResult.InsertedIDs))

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// status := redisClient.Ping()
	// log.Println(status)

	// recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)
}

func main() {
	router := gin.Default()
	// router.POST("/recipes", recipesHandler.NewRecipeHandler)
	// router.GET("/recipes", recipesHandler.ListRecipesHandler)
	// router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	// router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	// router.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	/*
		router.GET("/recipes/search", SearchRecipesHandler)*/
	router.Run()
}
