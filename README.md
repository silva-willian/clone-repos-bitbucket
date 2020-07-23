# clone-repos-bitbucket

Project created to clone all projects and branches of an owner in bitbucket

## Building and running

### Local Build

    export BITBUCKET_HOST=https://api.bitbucket.org/2.0
    export BITBUCKET_USER=your_user
    export BITBUCKET_PASSWORD=your_pass
    export BITBUCKET_OWNER=your_owner
    export BASE_FOLDER=files
    export APP_DIR=/Users/william/go/src/github.com/silva-willian/clone-repos-bitbucket #folder where the application is running locally 
    
    go run main.go

Keep an eye on the output to recover the projects

### Building with Docker

    docker build --tag clone-repos-bitbucket -f devops/application/build/Dockerfile .

    docker run \
        -e BITBUCKET_HOST=https://api.bitbucket.org/2.0 \
        -e BITBUCKET_USER=your_user \
        -e BITBUCKET_PASSWORD=your_pass \
        -e BITBUCKET_OWNER=your_owner \
        -e BASE_FOLDER=files \
        -e APP_DIR=/Users/william/go/src/github.com/silva-willian/clone-repos-bitbucket \
        clone-repos-bitbucket

## Releases

### 1.0.0 (01.04.2020)

* Initial release