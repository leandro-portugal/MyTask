package models

import "github.com/leandro-portugal/MyTask.git/db"

func Update(id int64, task Task) (int64, error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec(`UPDATE tasks SET title=$1, description=$2, createdAt=$3, completedAt=$4, done=$5 WHERE id=$6`, task.Title, task.Description, task.CreatedAt, task.CompletedAt, task.Done, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
