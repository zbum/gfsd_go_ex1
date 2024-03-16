package student

import (
	"database/sql"
	"fmt"
)

type Repository struct {
}

func NewStudentRepository() *Repository {
	return &Repository{}
}

func (r Repository) FindById(db *sql.DB, id int64) (*Student, error) {
	row := db.QueryRow("SELECT * FROM Students WHERE id = ?", id)

	var student Student
	if err := row.Scan(&student.ID, &student.Name); err != nil {
		return nil, fmt.Errorf("FindById %q: %v", id, err)
	}
	return &student, nil
}

func (r Repository) Insert(tx *sql.Tx, student *Student) error {
	result, err := tx.Exec("INSERT INTO Students VALUES (?, ?)", student.ID, student.Name)
	if err != nil {
		return fmt.Errorf("Insert : %v", student)
	}
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	return nil
}
