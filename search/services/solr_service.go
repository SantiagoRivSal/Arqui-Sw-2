package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"search/dto"
	client "search/services/repositories"
	e "search/utils/errors"
	"strings"

	log "github.com/sirupsen/logrus"
)

type SolrService struct {
	solr *client.SolrClient
}

func NewSolrServiceImpl(
	solr *client.SolrClient,
) *SolrService {
	return &SolrService{
		solr: solr,
	}
}

func (s *SolrService) GetQuery(query string) (dto.PropertiesDto, e.ApiError) {
	var propertiesDto dto.PropertiesDto
	queryParams := strings.Split(query, "_")
	field, query := queryParams[0], queryParams[1]
	propertiesDto, err := s.solr.GetQuery(query, field)
	if err != nil {
		return propertiesDto, e.NewBadRequestApiError("Solr failed")
	}
	return propertiesDto, nil
}

func (s *SolrService) GetQueryAllFields(query string) (dto.PropertiesDto, e.ApiError) {
	var propertiesDto dto.PropertiesDto
	propertiesDto, err := s.solr.GetQueryAllFields(query)
	if err != nil {
		return propertiesDto, e.NewBadRequestApiError("Solr failed")
	}
	return propertiesDto, nil
}

func (s *SolrService) AddFromId(id string) e.ApiError {
	var propertyDto dto.PropertyDto
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/properties/%s", "properties", 8090, id))
	if err != nil {
		log.Debugf("error getting property %s", id)
		return e.NewBadRequestApiError("error getting property " + id)
	}
	var body []byte
	body, _ = io.ReadAll(resp.Body)
	log.Debugf("%s", body)
	err = json.Unmarshal(body, &propertyDto)
	if err != nil {
		log.Debugf("error in unmarshal of property %s", id)
		return e.NewBadRequestApiError("error in unmarshal of property")
	}
	er := s.solr.Add(propertyDto)
	if er != nil {
		log.Debugf("error adding to solr")
		return e.NewInternalServerApiError("Adding to Solr error", err)
	}
	return nil
}

func (s *SolrService) Delete(id string) e.ApiError {
	err := s.solr.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
