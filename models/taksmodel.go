package models

import (
	"GOLANG-IMAM/config"
	"GOLANG-IMAM/entities"
	"database/sql"
)

type TaksModel struct {
	db *sql.DB
}

func New() *TaksModel {
	db, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &TaksModel{db: db}
}

func (m *TaksModel) FindAll(taks *[]entities.Taks) error {
	rows, err := m.db.Query("select * from tbl_taks ORDER BY id DESC")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Taks
		rows.Scan(
			&data.Id,
			&data.TaksName,
			&data.Assignor,
			&data.Deadline,
			&data.Finish)
		*taks = append(*taks, data)
	}

	return nil
}

func (m *TaksModel) Create(taks *entities.Taks) error {
	result, err := m.db.Exec("insert into tbl_taks (taks_name, assignor, deadline, finish) values(?,?,?,?)",
		taks.TaksName, taks.Assignor, taks.Deadline, taks.Finish)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	taks.Id = lastInsertId
	return nil
}

func (m *TaksModel) Find(id int64, taks *entities.Taks) error {
	return m.db.QueryRow("select * from tbl_taks where id = ?", id).Scan(
		&taks.Id,
		&taks.TaksName,
		&taks.Assignor,
		&taks.Deadline)
}

func (m *TaksModel) Update(taks entities.Taks) error {

	_, err := m.db.Exec("update tbl_taks set taks_name = ?, assignor = ?, deadline = ?",
		taks.TaksName, taks.Assignor, taks.Deadline)

	if err != nil {
		return err
	}

	return nil
}

func (m *TaksModel) Done(id int64) error {
	_, err := m.db.Exec("update tbl_taks set finish = 1 where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (m *TaksModel) Delete(id int64) error {
	_, err := m.db.Exec("delete from tbl_taks where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
