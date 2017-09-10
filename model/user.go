package model

import (
	"database/sql"
)

// 全て値があるときに利用する
func InsertUser(db *sql.DB, userId string, region Region) error {
	q := `insert into users (user_id, region) values ('?',?)`
	_, err := db.Exec(q, userId, region)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUserId(db *sql.DB, userId string) (User, error) {
	q := `select * from users where user_id='?''`
	var user User
	//TODO: Regionはstring? int?
	err := db.QueryRow(q, userId).Scan(&user.ID, &user.UserID, &user.Region)
	if err != nil {
		return user, err
	}

	return user, nil
}
