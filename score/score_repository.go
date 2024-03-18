package score

import (
	"database/sql"
	"fmt"
)

type Repository struct {
}

func NewScoreRepository() *Repository {
	return &Repository{}
}

func (r Repository) FindByStudentId(tx *sql.Tx, id int64) ([]Score, error) {
	var scores = make([]Score, 0)

	rows, err := tx.Query("SELECT * FROM Scores WHERE student_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("findByStudentId %q: %v", id, err)
	}
	defer rows.Close()

	for rows.Next() {
		var score Score
		if err != nil {
			return nil, fmt.Errorf("findByStudentId %q: %v", id, err)
		}
		scores = append(scores, score)
	}
	if err != nil {
		return nil, fmt.Errorf("findByStudentId %q: %v", id, err)
	}
	return scores, nil
}

func (r Repository) Insert(tx *sql.Tx, score *Score) error {
	result, err := tx.Exec("INSERT INTO Scores VALUES (?, ?, ?, ?)", score.ID, score.Semester, score.StudentId, score.Score)
	if err != nil {
		return fmt.Errorf("insert : %v", score)
	}
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	return nil
}
