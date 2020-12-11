package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// Global variables related to app info
const appName = "2bitprogrammers/api_mockup"
const appVersion = "2018.31a"
const appPort = "1234"

// A single key/value pair
type kvPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// The definition which specifies how the response should be formed
type endpointResponse struct {
	Headers []kvPair `json:"headers"`
	Payload string   `json:"payload"`
}

// This will determine if the endpointResponse is an empty struct (uninitialized with no data)
func (er endpointResponse) IsEmpty() bool {
	return reflect.DeepEqual(er, endpointResponse{})
}

// Global variables related to the config file
const defaultConfigFilename = "config_api_mockup.json"

var configFileData map[string]map[string]endpointResponse

// This will load the config file into the global variable.
func loadConfigFile() {
	fmt.Printf("Loading config file:  %s\n", defaultConfigFilename)
	cFile, err := ioutil.ReadFile(defaultConfigFilename)
	if err != nil {
		s := fmt.Sprintf("[ERROR] LoadConfigFile() - failed to load config file: %s\n%s", defaultConfigFilename, err)
		log.Fatal(s)
	}
	err = json.Unmarshal(cFile, &configFileData)
	if err != nil {
		s := fmt.Sprintf("[ERROR] LoadConfigFile() - invalid config file\n%s", err)
		log.Fatal(s)
	}
}

// This will return 404 not found for all endpoints which don't have a definition
func handleAPIInvalidRequest(w http.ResponseWriter, uri string, method string) {
	log.Printf("%d\t%s\t%s", http.StatusNotFound, method, uri)
	responsePayload := fmt.Sprintf(`{ "errors": [{ "status": 404, "message": "Resource Not Found", "uri": "%s", "method": "%s" }] }`, uri, method)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(responsePayload))
}

// This will handle all incoming API requests
func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	method := r.Method
	cURIMethod := configFileData[uri][method]
	if cURIMethod.IsEmpty() { // No endpoint definition in the configuration
		handleAPIInvalidRequest(w, uri, method)
	} else { // We found an endpoint definition in the configuration
		log.Printf("%d\t%s\t%s", http.StatusOK, method, uri)
		if len(cURIMethod.Headers) <= 0 { // No headers, use default
			w.Header().Set("Content-Type", "text/plain")
		} else { // Headers found, add each to the response
			for k := range cURIMethod.Headers {
				h := cURIMethod.Headers[k].Key
				hv := cURIMethod.Headers[k].Value
				w.Header().Set(h, hv)
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cURIMethod.Payload))
	}
}

func main() {
	// Load the configuration file
	loadConfigFile()

	// Output app info
	fmt.Printf("%s v%s\n", appName, appVersion)
	fmt.Println("www.2BitProgrammers.com\nCopyright (C) 2020. All Rights Reserved.\n")
	log.Printf("Starting App on Port %s", appPort)

	// Listen for all endpoints
	http.HandleFunc("/", handleAPIRequest)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
