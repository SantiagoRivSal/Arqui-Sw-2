package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"search/dto"
	e "search/utils/errors"

	log "github.com/sirupsen/logrus"
	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func (sc *SolrClient) AddClient(PropertyDto dto.PropertyDto) e.ApiError {
	var addPropertyDto dto.AddDto
	addPropertyDto.Add = dto.DocDto{Doc: PropertyDto}
	data, err := json.Marshal(addPropertyDto)

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

func (sc *SolrClient) GetQuery(query string, field string) (dto.PropertiesDto, e.ApiError) {
	var response dto.SolrResponseDto
	var propertiesDto dto.PropertiesDto

	q, err := http.Get(fmt.Sprintf("http://localhost:8983/solr/property/select?q=%s%s%s", field, ":", query))

	if err != nil {
		return propertiesDto, e.NewBadRequestApiError("error getting from solr")
	}

	defer q.Body.Close()

	err = json.NewDecoder(q.Body).Decode(&response)

	if err != nil {
		log.Debug("error: ", err)
		return propertiesDto, e.NewBadRequestApiError("error in unmarshal")
	}

	for _, doc := range response.Response.Docs {
		propertyDto := doc
		propertiesDto = append(propertiesDto, propertyDto)
	}

	return propertiesDto, nil
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
