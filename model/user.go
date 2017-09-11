package model

import (
	"database/sql"
	"log"
)

// 全て値があるときに利用する
func InsertUser(db *sql.DB, userId string, region Region) error {
	log.Println("InsertUser::", " userId: ", userId, " region: ", region)
	q := `insert into users (user_id, region) values ($1,$2)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userId, region)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUserId(db *sql.DB, userId string) (User, error) {
	log.Println("GettUser::", " userId: ", userId)
	q := `select * from users where user_id=$1`

	var user User

	stmt, err := db.Prepare(q)
	if err != nil {
		return user, err
	}

	//TODO: Regionはstring? int?
	err = stmt.QueryRow(userId).Scan(&user.ID, &user.UserID, &user.Region, &user.Created)
	if err != nil {
		return user, err
	}

	return user, nil
}

func IsUserExists(db *sql.DB, userId string)(bool, error) {
	q := `select count(*) from users where user_id=$1`

	stmt, err := db.Prepare(q)
	if err != nil {
		return false, err
	}

	var count int64
	err = stmt.QueryRow(userId).Scan(&count)
	return count==1, nil
}
