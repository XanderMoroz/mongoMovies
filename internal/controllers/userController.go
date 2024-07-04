package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/XanderMoroz/mongoMovies/configs"
	"github.com/XanderMoroz/mongoMovies/internal/models"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

// @Summary        create new user
// @Description    Creating User in DB with given request body
// @Tags           Users
// @ID				create-new-user
// @Accept         json
// @Produce        json
// @Param          request         		body        models.CreateUserBody    true    "Enter user data"
// @Success        201              	{string}    string
// @Failure        400              	{string}    string    "Bad Request"
// @Router         /api/users 			[post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	log.Println("Поступил запрос на создание новой записи в БД...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("При извлечении тела запроса - Произошла ошибка: <%v>\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		log.Println("...успешно")
		// log.Printf("Тело запроса: %+v", user)
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		log.Printf("При добавлении новой записи - Произошла ошибка: <%v>\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		log.Println("Новая запись успешно добавлена:")
		log.Printf("ID новой записи: %v", result.InsertedID)
	}

	json.NewEncoder(w).Encode(newUser)
}

// @Summary		get a user by ID
// @Description Get a user by ID
// @Tags 		Users
// @ID			get-user-by-id
// @Produce		json
// @Param		id					path		string			true	"UserID"
// @Success		200					{object}	models.User
// @Failure		404					{object}	[]string
// @Router		/api/users/{id} 	[get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	log.Printf("Поступил запрос на извлечение записи по ID: <%s>\n", params["id"])

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := params["id"]
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		log.Printf("При извлечении записи -произошла ошибка: <%v>\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		log.Printf("Запись успешно извлечена: <%+v>\n", user)
	}
	json.NewEncoder(w).Encode(user)
}

// @Summary		get all users
// @Description Get all users from db
// @Tags 		Users
// @ID			get-all-users
// @Produce		json
// @Success		200		{object}	[]models.User
// @Router		/api/users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Printf("При извлечении списка записей - произошла ошибка: <%v>\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			log.Printf("При обработке списка записей -произошла ошибка: <%v>\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		users = append(users, singleUser)
	}

	json.NewEncoder(w).Encode(users)
}

// @Summary			update user by ID
// @Description 	Update user by ID
// @ID				update-user-by-id
// @Tags 			Users
// @Produce			json
// @Param			id					path		string									true	"UserID"
// @Param           request         	body        models.CreateUserBody    true    	"Введите новые данные пользователя"
// @Success			200	{object}		[]string
// @Failure			404	{object}		[]string
// @Router			/api/users/{id} 	[put]
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	log.Printf("Поступил запрос на обновление записи по ID: <%s>\n", params["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := params["id"]
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	//validate the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("При извлечении тела запроса - Произошла ошибка: <%v>\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		log.Println("...успешно")
	}

	update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}

	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		log.Printf("При обновлении записи -произошла ошибка: <%v>\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get updated user details
	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
		if err != nil {
			log.Printf("При извлечении записи -произошла ошибка: <%v>\n", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	json.NewEncoder(w).Encode(updatedUser)
}

// @Summary		delete a user by ID
// @Description Delete a user by ID
// @ID			delete-user-by-id
// @Tags 		Users
// @Produce		json
// @Param		id					path		string		true	"UserID"
// @Success		200					{object}	[]string
// @Failure		404					{object}	[]string
// @Router		/api/users/{id} 	[delete]
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	log.Printf("Поступил запрос на удаление записи по ID: <%s>\n", params["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := params["id"]
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		log.Printf("При удалении записи - произошла ошибка: <%v>\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount < 1 {
		log.Printf("При извлечении тела запроса - Произошла ошибка: <%v>\n", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("User successfully deleted!")
}

//MongoDB helpers
// func checkNilError(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func deleteAllMovie() int64 {
// 	delCount, err := userCollection.DeleteMany(context.Background(), bson.D{{}}, nil)
// 	checkNilError(err)
// 	fmt.Println("No of movies deleted:", delCount.DeletedCount)
// 	return delCount.DeletedCount
// }