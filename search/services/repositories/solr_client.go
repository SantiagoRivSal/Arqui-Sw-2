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

func (sc *SolrClient) GetQuery(query string, field string) (dto.PropertiesArrayDto, e.ApiError) {
	var response dto.SolrResponseDto
	var propertiesArrayDto dto.PropertiesArrayDto

	q, err := http.Get(fmt.Sprintf("http://localhost:8983/solr/property/select?q=%s%s%s&rows=100000", field, ":", query))

	if err != nil {
		return propertiesArrayDto, e.NewBadRequestApiError("error getting from solr")
	}

	defer q.Body.Close()

	err = json.NewDecoder(q.Body).Decode(&response)

	if err != nil {
		log.Debug("error: ", err)
		return propertiesArrayDto, e.NewBadRequestApiError("error in unmarshal")
	}

	for _, doc := range response.Response.Docs {
		propertyArrayDto := doc
		propertiesArrayDto = append(propertiesArrayDto, propertyArrayDto)
	}

	return propertiesArrayDto, nil
}

func (sc *SolrClient) GetQueryAllFields(query []string) (dto.PropertiesArrayDto, e.ApiError) {
	var response dto.SolrResponseDto
	var propertiesArrayDto dto.PropertiesArrayDto

	//q, err := http.Get("http://localhost:8983/solr/property/select?q=service" + ":" + query[0] + "%0A" + "country" + ":" + query[1] + "%0A" + "city" + ":" + query[2])
	q, err := http.Get("http://localhost:8983/solr/property/select?q=service" + ":" + query[0] + "%20AND%20" + "country" + ":" + query[1] + "%20AND%20" + "city" + ":" + query[2] + "&rows=100000")
	if err != nil {
		log.Debug("error: ", err)
		return propertiesArrayDto, e.NewBadRequestApiError("error getting from solr")

	}

	var body []byte
	body, err = io.ReadAll(q.Body)
	if err != nil {
		log.Debug("error: ", err)
		return propertiesArrayDto, e.NewBadRequestApiError("error reading body")
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Debug("error: ", err)
		return propertiesArrayDto, e.NewBadRequestApiError("error in unmarshal")
	}

	for _, doc := range response.Response.Docs {
		propertyArrayDto := doc
		propertiesArrayDto = append(propertiesArrayDto, propertyArrayDto)
	}
	return propertiesArrayDto, nil
}
