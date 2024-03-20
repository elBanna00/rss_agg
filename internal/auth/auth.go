package auth

import (
	"errors"
	"net/http"
	"strings"
)

//Gets the API key from the Header of the http req

func GetApiKey(header http.Header) (string,error) {
	val := header.Get("Authorization");
	if val == "" {
		return "", errors.New("No Authrization");
	}
	vals := strings.Split(val," ");
	if len(vals) != 2 {
		return "",errors.New("Authrization header Doesn't match schema")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Authrization header Doesn't match schema")
	}
	return vals[1],nil
}