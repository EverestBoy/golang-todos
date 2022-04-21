/*
 * Copyright (c) 2022.
 * TO-DO Project Application
 * You can use this as a starter project or for your reference
 *
 */

package impl

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"todoGo/database"
	"todoGo/model"
	"todoGo/repository"
)

/*
@author everestboy
*/

// getting the db
var (
	dbConnectionUser = database.NewMongoDbConnection()
)

type authRepo struct {
}

func NewAuthRepository() repository.AuthRepository {
	return &authRepo{}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a authRepo) RegisterUser(user *model.User) (*model.User, error) {

	err, ctx, client, collection, cancel := dbConnectionUser.GetCollection("test", "user")
	defer cancel()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}

	// checking if user exists or not
	count, err := collection.CountDocuments(ctx, bson.D{
		{
			"$or",
			[]bson.D{
				bson.D{{"email", user.Email}},
				bson.D{{"phone", user.Phone}},
				bson.D{{"username", user.Username}},
			},
		},
	})

	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Errorf("User already exists")
	}
	// creating a user
	cur, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	insertedId := cur.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("Inserted user %s", insertedId)

	return user, err
}

func (a authRepo) UserDetail(email *string, username *string) (*model.User, error) {
	err, ctx, client, collection, cancel := dbConnectionUser.GetCollection("test", "user")
	defer cancel()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}
	// checking if user exists or not
	var user model.User
	userResult := collection.FindOne(ctx, bson.D{
		{
			"$or",
			[]bson.D{
				bson.D{{"email", email}},
				bson.D{{"username", username}},
			},
		},
	})
	err = userResult.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a authRepo) UserDetailById(userId *string) (*model.User, error) {
	err, ctx, client, collection, cancel := dbConnectionUser.GetCollection("test", "user")
	defer cancel()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}
	// checking if user exists or not
	var user model.User
	id, err := primitive.ObjectIDFromHex(*userId)
	if err != nil {
		return nil, err
	}
	userResult := collection.FindOne(ctx, bson.M{
		"_id": id,
	})
	err = userResult.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
