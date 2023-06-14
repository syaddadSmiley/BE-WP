package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"waroeng_pgn1/domain"
	db "waroeng_pgn1/internal/database"
	"waroeng_pgn1/internal/repository"
	"waroeng_pgn1/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {

	var databaseHelper *sql.DB
	var collectionHelper *mocks.Collection

	databaseHelper, _ = db.ConnectToDB()
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionUser

	mockUser := &domain.User{
		ID:       primitive.NewObjectID().String(),
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "password",
	}

	mockEmptyUser := &domain.User{}
	mockUserID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {

		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockUserID, nil).Once()

		databaseHelper.Query("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", mockUser.ID, mockUser.Name, mockUser.Email, mockUser.Password)

		ur := repository.NewUserRepository(databaseHelper, collectionName)

		err := ur.Create(context.Background(), mockUser)

		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockEmptyUser, errors.New("Unexpected")).Once()

		databaseHelper.Query("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", mockUser.ID, mockUser.Name, mockUser.Email, mockUser.Password)

		ur := repository.NewUserRepository(databaseHelper, collectionName)

		err := ur.Create(context.Background(), mockEmptyUser)

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
	})

}
