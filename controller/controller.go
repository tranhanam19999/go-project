package controller

import (
	"context"
	helper "echo-project/helper"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, "Home")
}

func GetUsers(c echo.Context) error {
	client := helper.GetMongoClient()
	collection := client.Database("GolangTestDB").Collection("user")

	if pingErr := client.Ping(context.TODO(), readpref.Primary()); pingErr != nil {
		print("Cant ping\n ", pingErr.Error())
	}

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, results)
}

func UpdateUser(c echo.Context) error {
	my_data := echo.Map{}
	if err := c.Bind(&my_data); err != nil {
		return err
	} else {
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid Id")
		}

		if len(id) > 0 {
			client := helper.GetMongoClient()
			collection := client.Database("GolangTestDB").Collection("user")

			var foundedUser bson.M
			if err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&foundedUser); err != nil {
				print("err", err.Error())
				return c.JSON(http.StatusBadGateway, "Can't find any user with that id")
			}

			// var updatedUser bson.M
			update := bson.D{
				primitive.E{Key: "$set", Value: bson.M{
					"username": "hanam123",
					"password": "abc12333",
				}},
			}
			newUpdatedUser, err := collection.UpdateByID(context.Background(), id, update)
			if err != nil {
				print(err.Error())
				return c.JSON(http.StatusBadRequest, "Something wrong")
			}
			//collection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, bson.M{"username": "hanam123", "password": "abc123"}).Decode(&updatedUser)
			return c.JSON(http.StatusOK, newUpdatedUser)
		}
		return c.JSON(http.StatusBadRequest, "You must fill in id field")
	}
}

func CreateUser(c echo.Context) error {
	my_data := echo.Map{}
	if err := c.Bind(&my_data); err != nil {
		return err
	} else {
		username := fmt.Sprintf("%v", my_data["username"])
		password := fmt.Sprintf("%v", my_data["password"])

		if username == "<nil>" || password == "<nil>" {
			return c.JSON(http.StatusBadRequest, "Must have two field username and password")
		}
		if len(username) > 0 && len(password) > 0 {
			client := helper.GetMongoClient()
			collection := client.Database("GolangTestDB").Collection("user")
			createdUser, err := collection.InsertOne(context.Background(), bson.M{"username": username, "password": password})

			if err != nil {
				log.Fatal(err)
				return c.JSON(http.StatusBadGateway, "Failed to create")
			}

			return c.JSON(http.StatusOK, createdUser)
		}
		return c.JSON(http.StatusBadRequest, "You must fill in of two field")
	}
}
func LogIn(c echo.Context) error {
	my_data := echo.Map{}
	if err := c.Bind(&my_data); err != nil {
		return err
	} else {
		username := fmt.Sprintf("%v", my_data["username"])
		password := fmt.Sprintf("%v", my_data["password"])

		if username == "<nil>" || password == "<nil>" {
			return c.JSON(http.StatusBadRequest, "Must have two field username and password")
		}

		if len(username) > 0 && len(password) > 0 {
			client := helper.GetMongoClient()
			collection := client.Database("GolangTestDB").Collection("user")

			var foundedUser bson.M
			if err = collection.FindOne(context.Background(), bson.M{"username": username, "password": password}).Decode(&foundedUser); err != nil {
				return c.JSON(http.StatusBadGateway, "Can't find any user with that username and password")
			}
			return c.JSON(http.StatusOK, foundedUser)
		}
		return c.JSON(http.StatusBadRequest, "You must fill in of two field")
	}
}
