package service

import (
	"testing"

	"challenge.haraj.com.sa/kraicklist/entity"
	"challenge.haraj.com.sa/kraicklist/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var RecordRepositoryMock = &repository.RecordRepositoryMock{Mock: mock.Mock{}}
var RecordServiceMock = RecordService{recordRepository: RecordRepositoryMock}

func TestSearch(t *testing.T) {
	query := "iPhone"
	records := []entity.Record{
		{
			ID:      63983811,
			Title:   "M attorney and legal advisor for drafting objection regulations",
			Content: "We do our best iPhone to achieve justice and strive to achieve the goal. To communicate and consider any case or advice\n##### We serve all regions of the Kingdom",
		},
		{
			ID:      63983811,
			Title:   "Basin covers for tundra and ford",
			Content: "Peace, mercy and blessings of God iPhone",
		},
		{
			ID:      63983811,
			Title:   "Stallion free trunk variety",
			Content: "A free solution is a kind of a stump that is lost in the halal.",
		},
	}
	expected := []entity.Record{
		{
			ID:      63983811,
			Title:   "M attorney and legal advisor for drafting objection regulations",
			Content: "We do our best iPhone to achieve justice and strive to achieve the goal. To communicate and consider any case or advice\n##### We serve all regions of the Kingdom",
		},
		{
			ID:      63983811,
			Title:   "Basin covers for tundra and ford",
			Content: "Peace, mercy and blessings of God iPhone",
		},
	}
	searcher := entity.Searcher{}
	searcher.Records = records

	RecordRepositoryMock.Mock.On("Load", "data.gz").Return(searcher).Once()
	RecordRepositoryMock.Mock.On("Load", "data.gz").Return(nil).Once()
	result, _ := RecordServiceMock.Search(query)

	assert.Equal(t, result, expected)
}

func TestSearchCompare(t *testing.T) {
	sampleData := []struct {
		request  string
		query    string
		record   entity.Record
		expected bool
	}{
		{
			request: "search found on title",
			query:   "legal",
			record: entity.Record{
				ID:      63983811,
				Title:   "M attorney and legal advisor for drafting objection regulations",
				Content: "We do our best to achieve justice and strive to achieve the goal. To communicate and consider any case or advice\n##### We serve all regions of the Kingdom",
			},
			expected: true,
		},
		{
			request: "search found on content",
			query:   "blessings",
			record: entity.Record{
				ID:      63983811,
				Title:   "Basin covers for tundra and ford",
				Content: "Peace, mercy and blessings of God",
			},
			expected: true,
		},
		{
			request: "search not found",
			query:   "iPhone",
			record: entity.Record{
				ID:      63983811,
				Title:   "Stallion free trunk variety",
				Content: "A free solution is a kind of a stump that is lost in the halal.",
			},
			expected: false,
		},
	}

	for _, data := range sampleData {
		t.Run(data.request, func(t *testing.T) {
			record, status := searchCompare(data.query, data.record)
			assert.Equal(t, data.expected, status)
			assert.Equal(t, data.record, record)
		})
	}
}
