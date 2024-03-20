package student

import (
	"context"
	"gfsd_go_ex1/common"
)

type Service struct {
	datasource        *common.DataSource
	studentRepository *Repository
}

func NewStudentService(datasource *common.DataSource, studentRepository *Repository) *Service {
	return &Service{datasource: datasource, studentRepository: studentRepository}
}

func (s Service) GetStudent(id int64) *Student {
	db := s.datasource.GetDB()
	student, err := s.studentRepository.FindById(db, id)
	if err != nil {
		return nil
	}
	return student
}

func (s Service) RegisterStudent(context context.Context, student *Student) (*Student, error) {
	db := s.datasource.GetDB()

	tx, err := db.BeginTx(context, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = s.studentRepository.Insert(tx, student)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return student, nil
}
