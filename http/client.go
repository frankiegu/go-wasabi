package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/inloco/go-wasabi/assignments"
	"github.com/inloco/go-wasabi/experiments"
)

type HttpClient struct {
	http.Client

	address string

	applicationName string

	login    string
	password string
}

func NewHttpClient(address, applicationName, login, password string) *HttpClient {
	return &HttpClient{
		address:         address,
		applicationName: applicationName,
		login:           login,
		password:        password,
	}
}

func (c *HttpClient) GenerateAssignment(ctx context.Context, experimentLabel string, userID string) (*assignments.Assignment, error) {

	url := c.address + generateAssignmentPath(c.applicationName, experimentLabel, userID)

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	assignment := &assignments.Assignment{}
	err = json.Unmarshal(body, assignment)

	return assignment, err
}

func (c *HttpClient) CreateExperiment(ctx context.Context, experiment *experiments.Experiment) (*experiments.Experiment, error) {
	url := c.address + createExperimentPath()

	payload, err := json.Marshal(experiment)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.login, c.password)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	experimentCreated := &experiments.Experiment{}
	err = json.Unmarshal(body, experimentCreated)
	return experimentCreated, err
}

func (c *HttpClient) CreateBucket(ctx context.Context, bucket *experiments.Bucket) (*experiments.Bucket, error) {
	url := c.address + createBucketPath(bucket.ExperimentID)

	payload, err := json.Marshal(bucket)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.login, c.password)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	bucketCreated := &experiments.Bucket{}
	err = json.Unmarshal(body, bucketCreated)
	return bucketCreated, err
}

func (c *HttpClient) GetExperiments(ctx context.Context) ([]*experiments.Experiment, error) {
	return nil, nil
}

func (c *HttpClient) GetExperimentByID(ctx context.Context, experimentID string) (*assignments.Assignment, error) {
	return nil, nil
}
