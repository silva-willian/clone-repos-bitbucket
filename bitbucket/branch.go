package bitbucket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetAllBranches retorna todos as branchs de um repositorio do bitbucket
func GetAllBranches(owner, repository string) ([]Branch, error) {
	host := fmt.Sprintf("%s/repositories/%s/%s/refs/branches", baseHost, owner, repository)

	result, err := getBranch(host)

	if err != nil {
		return nil, err
	}

	if result.Page == result.TotalPages {
		return result.Values, nil
	}

	found := result.Values
	nextPage := result.NextPage

	for nextPage != "" {

		newResult, err := getBranch(nextPage)

		if err != nil {
			return nil, err
		}

		nextPage = newResult.NextPage

		found = append(found, newResult.Values...)
	}

	return found, nil
}

func getBranch(host string) (BranchResult, error) {

	client := http.Client{}
	request, err := http.NewRequest("GET", host, nil)
	request.SetBasicAuth(user, password)

	res, err := client.Do(request)

	if err != nil {
		log.Panicf("Error in GET %s %v", host, err)
		return BranchResult{}, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicf("ReadError in GET %s %v", host, err)
		return BranchResult{}, nil
	}

	var result BranchResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Panicf("JsonError in GET %s %v", host, err)
		return BranchResult{}, err
	}

	return result, nil
}
