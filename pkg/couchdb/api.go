package couchdb

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/thymesave/funnel/pkg/config"
)

type createUserRequest struct {
	// ID contains the user id for couchdb
	ID string `json:"_id"`
	// Name is the plain username
	Name string `json:"name"`
	// Type is the type of user to create
	Type string `json:"type"`
	// Roles is a list with roles to assign to the user
	Roles []string `json:"roles"`
	// Password is the plain text password
	Password string `json:"password"`
}

// CreateUserResponse transports relevant information about the user and userdb creation
type CreateUserResponse struct {
	// NewUser states if the user has been created or if it already exists
	NewUser bool `json:"-"`
	// DBName contains the database name
	DBName string `json:"dbName"`
}

// Client for couchdb api operations
type Client struct {
	// HTTPClient used for http requests
	HTTPClient *http.Client
}

// NewClient creates a new couchdb api client
func NewClient() *Client {
	return &Client{HTTPClient: http.DefaultClient}
}

func newCreateUserRequest(username string) createUserRequest {
	rand := sha256.New()
	_, _ = sha256.New().Write([]byte(uuid.New().String()))

	return createUserRequest{
		ID:       fmt.Sprintf("org.couchdb.user:%s", username),
		Name:     username,
		Type:     "user",
		Roles:    []string{},
		Password: hex.EncodeToString(rand.Sum(nil)),
	}
}

// CreateUser creates a new user in couchdb and the couch_per_user plugin creates a new userdb
func CreateUser(c *Client, couchDBConfig *config.CouchDB, username string) (*CreateUserResponse, error) {
	userReq, _ := json.Marshal(newCreateUserRequest(username))

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/_users/org.couchdb.user:%s", couchDBConfig.Endpoint(), username), bytes.NewReader(userReq))
	setAuthHeader(&req.Header, couchDBConfig.AdminUser, []string{"_admin"})
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if res.StatusCode != http.StatusConflict && res.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, errors.New("failed to create user, upstream response: " + string(body))
	}

	return &CreateUserResponse{
		NewUser: res.StatusCode == http.StatusCreated,
		DBName:  fmt.Sprintf("userdb-%s", hex.EncodeToString([]byte(username))),
	}, nil
}
