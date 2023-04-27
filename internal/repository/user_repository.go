package repository

import (
	"context"
	"errors"

	"github.com/gocql/gocql"

	"waroeng_pgn1/domain"
)

type userRepository struct {
	database   *gocql.Session
	collection string
}

func NewUserRepository(db *gocql.Session, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {

	query := ur.database.Query(`INSERT INTO users (id, email, password, name, phone_number) VALUES (?, ?, ?, ?, ?)`, user.ID, user.Email, user.Password, user.Name, user.PhoneNumber).Exec()
	if query == nil {
		return nil
	}

	// query2 := `INSERT INTO users (id, email, password, name, phone_number) VALUES (?, ?, ?, ?, ?)`
	// database.ExecuteQuery(query2, uuid, user.Email, user.Password, user.Name, user.PhoneNumber)

	return errors.New("error while creating user")
}

// func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
// 	collection := ur.database.Collection(ur.collection)

// 	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
// 	cursor, err := collection.Find(c, bson.D{}, opts)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var users []domain.User

// 	err = cursor.All(c, &users)
// 	if users == nil {
// 		return []domain.User{}, err
// 	}

// 	return users, err
// }

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	query := `SELECT id, email, password, name, phone_number FROM users WHERE email = ?`
	err := ur.database.Query(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.PhoneNumber)
	if err != nil {
		return user, errors.New("error get user by email")
	}
	return user, nil
}

// func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
// 	collection := ur.database.Collection(ur.collection)

// 	var user domain.User

// 	idHex, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return user, err
// 	}

// 	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
// 	return user, err
// }
