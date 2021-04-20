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
		result_search, status_search := searchCompare(query, record)
		if status_search {
			result = append(result, result_search)
		}
	}
	return result, nil
}

func searchCompare(query string, record entity.Record) (entity.Record, bool) {
	if strings.Contains(strings.ToLower(record.Title), strings.ToLower(query)) || strings.Contains(strings.ToLower(record.Content), strings.ToLower(query)) {
		return record, true
	} else {
		return record, false
	}
}
