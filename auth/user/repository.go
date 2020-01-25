package user

import (
	"crypto/sha256"
	"database/sql"
	pb "github.com/ArtGooner/test-microservice/config"
)

type Repository interface {
	//Create(user *User) error
	Get(user *pb.Account) (*pb.User, error)
}
type repository struct {
	db *sql.DB
}

//// Create the user record
//func (r repository) Create(user *pb.User) error {
//	_, err := r.db.Exec("insert into Users (Email, PasswordHash) values ($1,$2);", user.Email, user.PasswordHash)
//	return err
//}

//Get the user record
func (r repository) Get(acc *pb.Account) (*pb.User, error) {
	h := sha256.New()
	h.Write([]byte(acc.GetPassword()))
	pwd := h.Sum(nil)
	rows, err := r.db.Query("SELECT * from Users where Email=$1 AND PasswordHash = $2;", acc.GetEmail(), pwd)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int32
		var email string
		var name string
		var surname string
		var age int32
		var passwordHash []byte

		err := rows.Scan(&id, &email, &name, &surname, &age, &passwordHash)

		if err != nil {
			return nil, err
		}

		return &pb.User{Id: id, Email: email, Name: name, Surname: surname, Age: age, PasswordHash: passwordHash}, nil
	}
	return nil, nil
}

func NewRepository() (Repository, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=213612458 dbname=therion sslmode=disable")

	if err != nil {
		return nil, err
	}

	sqlStmt := `create table if not exists users (Id serial primary key, PasswordHash bytea, Name text, Surname text, Age integer)`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		return nil, err
	}

	return &repository{db: db}, nil
}
