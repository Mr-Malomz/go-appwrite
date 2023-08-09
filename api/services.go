package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-appwrite/data"
	"io/ioutil"
	"net/http"
)

//get details from environment variable
var projectId = GetEnvVariable("PROJECT_ID")
var databaseId = GetEnvVariable("DATABASE_ID")
var collectionId = GetEnvVariable("COLLECTION_ID")
var apiKey = GetEnvVariable("PROJECT_ID")

func (app *Config) createProject(newProject *data.ProjectRequest) (*data.ProjectResponse, error) {
	url := fmt.Sprintf("https://cloud.appwrite.io/v1/databases/%s/collections/%s/documents", databaseId, collectionId)

	createdProject := data.ProjectResponse{}
	jsonData := data.JsonAPIBody{
		DocumentId: "unique()",
		Data:       newProject,
	}
	postBody, _ := json.Marshal(jsonData)
	bodyData := bytes.NewBuffer(postBody)

	//making the request
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bodyData)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Appwrite-Key", apiKey)
	req.Header.Add("X-Appwrite-Project", projectId)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &createdProject)
	if err != nil {
		return nil, err
	}
	return &createdProject, nil
}

func (app *Config) getProject(documentId string) (*data.Project, error) {
	url := fmt.Sprintf("https://cloud.appwrite.io/v1/databases/%s/collections/%s/documents/%s", databaseId, collectionId, documentId)

	projectDetail := data.Project{}

	//making the request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Appwrite-Key", apiKey)
	req.Header.Add("X-Appwrite-Project", projectId)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &projectDetail)
	if err != nil {
		return nil, err
	}
	return &projectDetail, nil
}

func (app *Config) updateProject(updatedProject *data.ProjectRequest, documentId string) (*data.ProjectResponse, error) {
	url := fmt.Sprintf("https://cloud.appwrite.io/v1/databases/%s/collections/%s/documents/%s", databaseId, collectionId, documentId)

	updates := data.ProjectResponse{}
	jsonData := data.JsonAPIBody{
		Data: updatedProject,
	}
	postBody, _ := json.Marshal(jsonData)
	bodyData := bytes.NewBuffer(postBody)

	//making the request
	client := &http.Client{}
	req, _ := http.NewRequest("PATCH", url, bodyData)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Appwrite-Key", apiKey)
	req.Header.Add("X-Appwrite-Project", projectId)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &updates)
	if err != nil {
		return nil, err
	}
	return &updates, nil
}

func (app *Config) deleteProject(documentId string) (string, error) {
	url := fmt.Sprintf("https://cloud.appwrite.io/v1/databases/%s/collections/%s/documents/%s", databaseId, collectionId, documentId)

	//making the request
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Appwrite-Key", apiKey)
	req.Header.Add("X-Appwrite-Project", projectId)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return documentId, nil
}
