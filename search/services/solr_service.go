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

	logger "github.com/sirupsen/logrus"
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

func (s *SolrService) Add(id string) e.ApiError {
	var propertyDto dto.PropertyDto
	resp, err := http.Get("http://localhost:8090/properties/" + id + "/id")
	if err != nil {
		logger.Debugf("error getting property %s: %v", id, err)
		return e.NewBadRequestApiError("error getting property " + id)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Debugf("bad status code %d", resp.StatusCode)
		return e.NewBadRequestApiError("bad status code " + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Debugf("error reading response body: %v", err)
		return e.NewBadRequestApiError("error reading response body")
	}

	err = json.Unmarshal(body, &propertyDto)

	if err != nil {
		logger.Debugf("error in unmarshal of property %s: %v", id, err)
		return e.NewBadRequestApiError("error in unmarshal of property")
	}

	err = s.solr.AddClient(propertyDto)
	if err != nil {
		logger.Debugf("error adding to solr: %v", err)
		return e.NewInternalServerApiError("Adding to Solr error", err)
	}
	return nil
}

/*func (s *SolrService) GetQuery(query string) (dto.PropertiesArrayDto, e.ApiError) {
	var propertiesArrayDto dto.PropertiesArrayDto
	queryParams := strings.Split(query, "_")
	field, query := queryParams[0], queryParams[1]
	propertiesArrayDto, err := s.solr.GetQuery(query, field)
	if err != nil {
		return propertiesArrayDto, e.NewBadRequestApiError("Solr failed")
	}
	return propertiesArrayDto, nil
}*/

func (s *SolrService) GetQuery(query string) (dto.PropertiesArrayDto, e.ApiError) {
	var propertiesArrayDto dto.PropertiesArrayDto
	queryParams := strings.Split(query, "_")
	field, query := queryParams[0], queryParams[1]

	// Replace spaces with %20 in the query string
	query = strings.Replace(query, " ", "%20", -1)

	propertiesArrayDto, err := s.solr.GetQuery(query, field)
	if err != nil {
		return propertiesArrayDto, e.NewBadRequestApiError("Solr failed")
	}

	return propertiesArrayDto, nil
}

/*func (s *SolrService) GetQueryAllFields(query string) (dto.PropertiesArrayDto, e.ApiError) {
	var propertiesArrayDto dto.PropertiesArrayDto
	queryParams := strings.Split(query, "_")
	var q []string
	for i := 0; i < len(queryParams); i++ {
		q = append(q, queryParams[i])
	}
	propertiesArrayDto, err := s.solr.GetQueryAllFields(q)
	if err != nil {
		return propertiesArrayDto, e.NewBadRequestApiError("Solr failed")
	}
	return propertiesArrayDto, nil
}*/

func (s *SolrService) GetQueryAllFields(query string) (dto.PropertiesArrayDto, e.ApiError) {
	var propertiesArrayDto dto.PropertiesArrayDto
	queryParams := strings.Split(query, "_")
	var q []string
	for i := 0; i < len(queryParams); i++ {
		// Reemplazar los espacios por una cadena vacía en cada parámetro
		q = append(q, strings.ReplaceAll(queryParams[i], " ", "%20"))
	}
	propertiesArrayDto, err := s.solr.GetQueryAllFields(q)
	if err != nil {
		return propertiesArrayDto, e.NewBadRequestApiError("Solr failed")
	}
	return propertiesArrayDto, nil
}
