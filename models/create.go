package models

import "github.com/leandro-portugal/MyTask.git/db"

func Create(task Task) (id int64, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO tasks (title, description, createdAt, completedAt, done) VALUES ($1,$2,$3,$4,$5) RETURNING id`

	err = conn.QueryRow(sql, task.Title, task.Description, task.CreatedAt, task.CompletedAt, task.Done).Scan(&id)

	return

}
