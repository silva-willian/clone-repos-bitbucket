package bitbucket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetAllProjects retorna todos os projetos do bitbucket
func GetAllProjects(owner string) ([]Repository, error) {
	host := fmt.Sprintf("%s/repositories/%s/", baseHost, owner)

	result, err := getProject(host)

	if err != nil {
		return nil, err
	}

	if result.Page == result.TotalPages {
		return result.Values, nil
	}

	found := result.Values
	nextPage := result.NextPage

	for nextPage != "" {

		newResult, err := getProject(nextPage)

		if err != nil {
			return nil, err
		}

		nextPage = newResult.NextPage

		found = append(found, newResult.Values...)
	}

	return found, nil
}

func getProject(host string) (RepositoryResult, error) {

	client := http.Client{}
	request, err := http.NewRequest("GET", host, nil)
	request.SetBasicAuth(user, password)

	res, err := client.Do(request)

	if err != nil {
		log.Panicf("Error in GET %s %v", host, err)
		return RepositoryResult{}, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicf("ReadError in GET %s %v", host, err)
		return RepositoryResult{}, nil
	}

	var result RepositoryResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Panicf("JsonError in GET %s %v", host, err)
		return RepositoryResult{}, err
	}

	return result, nil
}
