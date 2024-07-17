package endpoint

import "github.com/DangerZombie/golang-gin-boilerplate/helper/auth"

type endpointImpl struct {
	authHelper auth.AuthHelper
}

func NewEndpoint(ah auth.AuthHelper) Endpoint {
	return &endpointImpl{ah}
}
