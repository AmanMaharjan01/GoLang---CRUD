package repositories

import (
	"database/sql"
	"fmt"
	"go-postgres/database"
	"go-postgres/models"
	"log"
)

func InsertUser(user models.User) int64 {

	// create the postgres db connection
	db := database.CreateConnection();

	// close the db connection
	defer db.Close()

	sqlStatement := `INSERT INTO customer (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

	var id int64

	err := db.QueryRow(sqlStatement, user.Name, user.Address).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func GetUser(id int64) (models.User, error) {
	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	var user models.User

	sqlStatement := `SELECT * FROM customer WHERE userid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Name, &user.Address)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, err
}

func GetAllUsers() ([]models.User, error) {
	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	var users []models.User

	sqlStatement := `SELECT * FROM customer`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Address)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		users = append(users, user)

	}

	return users, err
}

func UpdateUser(id int64, user models.User) int64 {

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	sqlStatement := `UPDATE customer SET name=$2, address=$3 WHERE userid=$1`

	res, err := db.Exec(sqlStatement, id, user.Name, user.Address)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteUser(id int64) int64 {

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	sqlStatement := `DELETE FROM customer WHERE userid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
