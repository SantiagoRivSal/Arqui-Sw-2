package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"search/dto"

	//logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func (s *SolrClient) AddClient(propertyDto dto.PropertyDto) error {
	data, err := json.Marshal(propertyDto)
	if err != nil {
		return fmt.Errorf("error marshalling propertyDto: %v", err)
	}

	reader := bytes.NewReader(data)
	if _, err := s.Client.Update(context.TODO(), "properties", solr.JSON, reader); err != nil {
		return fmt.Errorf("error in solr update: %v", err)
	}

	if err := s.Client.Commit(context.TODO(), "properties"); err != nil {
		return fmt.Errorf("error committing to solr: %v", err)
	}
	return nil
}

/*
func (sc *SolrClient) Delete(id string) e.ApiError {
	var deleteDto dto.DeleteDto
	deleteDto.Delete = dto.DeleteDoc{Query: fmt.Sprintf("id:%s", id)}
	data, err := json.Marshal(deleteDto)
	reader := bytes.NewReader(data)
	if err != nil {
		return e.NewBadRequestApiError("Error getting json")
	}
	resp, err := sc.Client.Update(context.TODO(), sc.Collection, solr.JSON, reader)
	logger.Debug(resp)
	if err != nil {
		return e.NewBadRequestApiError("Error in solr")
	}

	er := sc.Client.Commit(context.TODO(), sc.Collection)
	if er != nil {
		logger.Debug("Error committing load")
		return e.NewInternalServerApiError("Error committing to solr", er)
	}
	return nil
}
*/
