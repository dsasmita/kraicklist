package service

import (
	"log"
	"strings"

	"challenge.haraj.com.sa/kraicklist/entity"
	"challenge.haraj.com.sa/kraicklist/repository"
)

type RecordServiceInterface interface {
	Search(query string) ([]entity.Record, error)
}

type RecordService struct {
	recordRepository repository.RecordRepositoryInterface
}

func PrepareRecordServices(record repository.RecordRepositoryInterface) *RecordService {
	return &RecordService{recordRepository: record}
}

func (s *RecordService) Search(query string) ([]entity.Record, error) {
	var result []entity.Record
	records, err := s.recordRepository.Load("data.gz")
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records.Records {
		if strings.Contains(record.Title, query) || strings.Contains(record.Content, query) {
			result = append(result, record)
		}
	}
	return result, nil
}
