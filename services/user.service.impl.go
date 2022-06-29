package services

import (
	"context"
	"errors"

	"github.com/TomasHansut/gin_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	// Pointer to user collection
	usercollection *mongo.Collection
	ctx            context.Context
}

// Constructor
func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

// Pass user and insert it to DB return any error
func (u *UserServiceImpl) CreateUser(user *models.User) error {
	// if error return error if success nothing
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

// Pass user name get user object or error
func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	// Build query for db
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	// Get user from Db
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

// Get all users as slice from Db return any error
func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

// Pass user obejct return any arry
func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	// Build filter by witch user will be find
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	// Build updated user
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_address", Value: user.Address}, bson.E{Key: "user_age", Value: user.Age}}}}
	// Update Db
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched document found for update")
	}
	return nil
}

// Pass user name as string return any error
func (u *UserServiceImpl) DeleteUser(name *string) error {
	// Build fiilter by witch user will be find
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	// Delete user from Db
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No matched document found for update")
	}
	return nil
}
