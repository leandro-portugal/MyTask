package models

import "github.com/leandro-portugal/MyTask.git/db"

func read(id int64) (task Task, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM tasks WHERE ID=$1`, id)

	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.CompletedAt, &task.Done)

	return
}
