package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alastairruhm/guidor/client/context"
)

const (
	instancesBasePath = "/api/v1/instances"

	// ActionInProgress is an in progress action status
	ActionInProgress = "in-progress"

	//ActionCompleted is a completed action status
	ActionCompleted = "completed"
)

// InstancesService handles communction with action related methods of the
type InstancesService interface {
	List(context.Context, *ListOptions) ([]Instance, *Response, error)
	Get(context.Context, int) (*Instance, *Response, error)
	Register(context.Context, Instance) (*Instance, *Response, error)
}

// InstancesServiceOp handles communition with the image action related methods of the
type InstancesServiceOp struct {
	client *Client
}

var _ InstancesService = &InstancesServiceOp{}

type instancesRoot struct {
	Instances []Instance `json:"instances"`
}

type instanceRoot struct {
	Instance *Instance `json:"data"`
}

// Instance represents a database instance
type Instance struct {
	ID        string    `json:"id"`
	Token     string    `json:"token" bson:"token"`
	IP        string    `json:"ip" bson:"ip"`
	Hostname  string    `json:"hostname" bson:"hostname"`
	DbType    string    `json:"db_type" bson:"db_type"`
	DbVersion string    `json:"db_version" bson:"db_version"`
	DbName    string    `json:"db_name" bson:"db_name"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

// List all instances
func (s *InstancesServiceOp) List(ctx context.Context, opt *ListOptions) ([]Instance, *Response, error) {
	path := instancesBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(instancesRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Instances, resp, err
}

// Get an instance by ID.
func (s *InstancesServiceOp) Get(ctx context.Context, id int) (*Instance, *Response, error) {
	if id < 1 {
		return nil, nil, NewArgError("id", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s/%d", instancesBasePath, id)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(instanceRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Instance, resp, err
}

// Register an instance by ID.
func (s *InstancesServiceOp) Register(ctx context.Context, i Instance) (*Instance, *Response, error) {
	// if id < 1 {
	// 	return nil, nil, NewArgError("id", "cannot be less than 1")
	// }

	path := fmt.Sprintf("%s", instancesBasePath)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(instanceRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Instance, resp, err
}

func (a Instance) String() string {
	return Stringify(a)
}
