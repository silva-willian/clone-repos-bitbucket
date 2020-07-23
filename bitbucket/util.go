package bitbucket

import "os"

var baseHost string = os.Getenv("BITBUCKET_HOST")
var user string = os.Getenv("BITBUCKET_USER")
var password string = os.Getenv("BITBUCKET_PASSWORD")
