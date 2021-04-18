package repository

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"

	"challenge.haraj.com.sa/kraicklist/entity"
)

type RecordRepositoryInterface interface {
	Load(filepath string) (*entity.Searcher, error)
}

type RecordRepository struct {
	searcher *entity.Searcher
}

func PrepareRecordRepository(searcher entity.Searcher) *RecordRepository {
	return &RecordRepository{searcher: &searcher}
}

func (record *RecordRepository) Load(filepath string) (*entity.Searcher, error) {
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to open source file due: %v", err)
	}
	defer file.Close()
	// read as gzip
	reader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize gzip reader due: %v", err)
	}
	// read the reader using scanner to contstruct records
	var records []entity.Record
	cs := bufio.NewScanner(reader)
	for cs.Scan() {
		var r entity.Record
		err = json.Unmarshal(cs.Bytes(), &r)
		if err != nil {
			continue
		}
		records = append(records, r)
	}
	record.searcher.Records = records

	return record.searcher, nil
}
