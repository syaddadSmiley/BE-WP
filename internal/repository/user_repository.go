package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"waroeng_pgn1/domain"
)

type userRepository struct {
	database   *sql.DB
	collection string
}

func NewUserRepository(db *sql.DB, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	// sqlStatement := `INSERT INTO products (id_jenis, gambar, judul, deskripsi, ukuran, kondisi, harga, kuantitas) VALUES (?, ?, ?, ?,? ,? ,? ,?);`

	// stmt, err := a.db.Prepare(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }

	// defer stmt.Close()

	// result, err := stmt.Exec(Idjenis, Gambar, Judul, Deskripsi, Ukuran, Kondisi, Harga, Kuantitas)
	// if err != nil {
	// 	panic(err)
	// }

	// return result.LastInsertId()

	stmt, err := ur.database.Prepare(`INSERT INTO users (id, email, password, name, phone_number, role) VALUES (?, ?, ?, ?, ?, ?);`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(user.ID, user.Email, user.Password, user.Name, user.PhoneNumber, "user_wp")
	if err != nil {
		return err
	} else if result != nil {
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
	// rows, err := a.db.Query(`
	// SELECT
	// 	products.Id,
	// 	jenis_products.Jenis AS jenis_products,
	// 	products.gambar,
	// 	products.Judul,
	// 	products.deskripsi,
	// 	products.ukuran,
	// 	products.harga
	// FROM products
	// INNER JOIN jenis_products
	// ON products.Id_jenis = jenis_products.Id`)
	// if err != nil {
	// 	return []Task{}, err
	// }
	// query := ``
	fmt.Println("MASUK SINI GK SIH")
	// err := ur.database.Query(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.PhoneNumber)
	userRow, err := ur.database.Query(`SELECT id, email, password, name, phone_number FROM users WHERE email = ?`, email)
	if err != nil {
		return user, errors.New("error get user by email")
	}
	// defer userRow.Close()
	for userRow.Next() {
		err = userRow.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.PhoneNumber)
		if err != nil {
			return user, err
		}
	}
	// result := []Task{}
	// for rows.Next() {
	// 	admin := Task{}
	// 	err = rows.Scan(&admin.Id, &admin.JenisProducts, &admin.Gambar, &admin.Judul, &admin.Deskripsi, &admin.Ukuran, &admin.Harga)
	// 	if err != nil {
	// 		return []Task{}, err
	// 	}
	// 	result = append(result, admin)
	// }

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
