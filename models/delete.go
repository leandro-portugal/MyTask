package models

import "github.com/leandro-portugal/MyTask.git/db"

func Delete(id int64) (int64, error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec(`DELETE FROM WHERE id=$1`, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
