package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"search/dto"
	client "search/services/repositories"
	e "search/utils/errors"
	"strconv"
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
	resp, err := http.Get("http://localhost:8090/properties/" + id + "/id")
	if err != nil {
		log.Debugf("error getting property %s", id)
		return e.NewBadRequestApiError("error getting property " + id)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Debugf("bad status code %d", resp.StatusCode)
		return e.NewBadRequestApiError("bad status code " + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Debugf("error reading response body")
		return e.NewBadRequestApiError("error reading response body")
	}

	err = json.Unmarshal(body, &propertyDto)
	log.Debug("propiedad: ", propertyDto.Tittle)

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
