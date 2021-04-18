package repository

import (
	"challenge.haraj.com.sa/kraicklist/entity"
	"github.com/stretchr/testify/mock"
)

type RecordRepositoryMock struct {
	Mock mock.Mock
}

func (repository *RecordRepositoryMock) Load(filepath string) *entity.Searcher {
	arguments := repository.Mock.Called(filepath)
	if arguments.Get(0) == nil {
		return nil
	} else {
		result := arguments.Get(0).(entity.Searcher)
		return &result
	}
}
