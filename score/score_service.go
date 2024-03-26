package score

import (
	"context"
	"database/sql"
	"gfsd_go_ex1/common"
)

type Service struct {
	datasource      *common.DataSource
	scoreRepository *Repository
}

func NewScoreService(datasource *common.DataSource, studentRepository *Repository) *Service {
	return &Service{datasource: datasource, scoreRepository: studentRepository}
}

var (
	txOptions         = &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false}
	readOnlyTxOptions = &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}
)

func (s Service) GetScores(context context.Context, id int64) []Score {
	db := s.datasource.GetDB()

	tx, err := db.BeginTx(context, readOnlyTxOptions)
	scores, err := s.scoreRepository.FindByStudentId(tx, id)
	if err != nil {
		return nil
	}
	return scores
}

func (s Service) RegisterStudent(context context.Context, score *Score) (*Score, error) {
	db := s.datasource.GetDB()

	tx, err := db.BeginTx(context, txOptions)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = s.scoreRepository.Insert(tx, score)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return score, nil
}
