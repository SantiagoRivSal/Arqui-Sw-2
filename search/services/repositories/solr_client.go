package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"search/dto"
	e "search/utils/errors"

	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func (sc *SolrClient) GetQuery(query string, field string) (dto.PropertiesDto, e.ApiError) {
	var response dto.SolrResponseDto
	var propertiesDto dto.PropertiesDto
	q, err := http.Get(fmt.Sprintf("http://%s:%d/solr/properties/select?q=%s%s%s", "solr", 8983, field, "%3A", query))

	if err != nil {
		return propertiesDto, e.NewBadRequestApiError("error getting from solr")
	}

	defer q.Body.Close()
	err = json.NewDecoder(q.Body).Decode(&response)
	if err != nil {
		return propertiesDto, e.NewBadRequestApiError("error in unmarshal")
	}
	propertiesDto = response.Response.Docs
	return propertiesDto, nil
}

func (sc *SolrClient) GetQueryAllFields(query string) (dto.PropertiesDto, e.ApiError) {
	var response dto.SolrResponseDto
	var propertiesDto dto.PropertiesDto

	q, err := http.Get(
		fmt.Sprintf("http://%s:%d/solr/properties/select?q=city%s%s%scountry%s%s%sservice%s%s%s",
			"solr", 8983,
			"%3A", query, "%0A",
			"%3A", query, "%0A",
			"%3A", query, "%0A",
		))
	if err != nil {
		return propertiesDto, e.NewBadRequestApiError("error getting from solr")
	}

	var body []byte
	body, err = io.ReadAll(q.Body)
	if err != nil {
		return propertiesDto, e.NewBadRequestApiError("error reading body")
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return propertiesDto, e.NewBadRequestApiError("error in unmarshal")
	}

	propertiesDto = response.Response.Docs
	return propertiesDto, nil
}

func (sc *SolrClient) Add(propertyDto dto.PropertyDto) e.ApiError {
	var addpropertieDto dto.AddDto
	addpropertieDto.Add = dto.DocDto{Doc: propertyDto}
	data, err := json.Marshal(addpropertieDto)

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
