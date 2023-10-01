package db

import (
	"database/sql"
	"fmt"

	"github.com/leandro-portugal/MyTask.git/settings"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := settings.GetDataBase()

	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Hostname, conf.Port, conf.Username, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", connect)

	if err != nil {
		return nil, fmt.Errorf("erro ao abrir a conexão com o banco de dados: %w", err)
	}

	return conn, nil
}
