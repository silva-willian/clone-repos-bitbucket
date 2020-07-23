package main

import (
	"fmt"
	"log"
	"os"

	"github.com/silva-willian/clone-repos-bitbucket/bitbucket"
	"github.com/silva-willian/clone-repos-bitbucket/utils"
)

var baseFolder string = os.Getenv("BASE_FOLDER")
var owner string = os.Getenv("BITBUCKET_OWNER")

func main() {

	log.Println("Starting execution")

	err := utils.ValidateEnvs()

	if err != nil {
		utils.ReturnError(err)
	}

	err = utils.RemoveFolder(baseFolder)

	if err != nil {
		utils.ReturnError(err)
	}

	repositories, err := bitbucket.GetAllProjects(owner)

	if err != nil {
		utils.ReturnError(err)
	}

	log.Printf("Existem %d repositórios mapeados", len(repositories))

	for _, repository := range repositories {
		projectPath := fmt.Sprintf("%s/%s/%s/", baseFolder, repository.Project.Key, repository.Name)

		log.Printf("Iniciando o projeto %s no path %s", repository.Name, projectPath)

		err = utils.CreateFolder(projectPath)

		if err != nil {
			log.Printf("Erro ao criar a pasta fisica  do projeto %s", repository.Name)
			utils.ReturnError(err)
		}

		log.Print("Recuperando as branches do projeto")

		branches, err := bitbucket.GetAllBranches(owner, repository.Name)

		if err != nil {
			utils.ReturnError(err)
		}

		log.Printf("Existem %d branches associadas ao repositório %s", len(branches), repository.Name)

		for _, branch := range branches {

			err = utils.Clone(repository.Links.Clone[0].Href, branch.Name, projectPath)

			if err != nil {
				utils.ReturnError(err)
			}
		}

		log.Printf("Projeto %s clonado com sucesso", repository.Name)
	}

	if err != nil {
		utils.ReturnError(err)
	}

}
