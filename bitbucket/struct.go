package bitbucket

// RepositoryResult representa o retorno exato do serviço
type RepositoryResult struct {
	TotalPages int32        `json:"pagelen"`
	Size       int32        `json:"size"`
	Page       int32        `json:"page"`
	Values     []Repository `json:"values"`
	NextPage   string       `json:"next"`
}

// Repository representa o respositorio no bitbucket
type Repository struct {
	Name    string `json:"name"`
	Project struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	} `json:"project"`
	Links struct {
		Clone []struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"clone"`
	} `json:"links,omitempty"`
}

// BranchResult representa o retorno exato do serviço
type BranchResult struct {
	TotalPages int32    `json:"pagelen"`
	Size       int32    `json:"size"`
	Page       int32    `json:"page"`
	Values     []Branch `json:"values"`
	NextPage   string   `json:"next"`
}

// Branch representa o retorno exato do serviço
type Branch struct {
	Name string `json:"name"`
}
