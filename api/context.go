package api

import (
	"time"
)

type EnvironmentVariable struct {
	Variable string
	ContextID string
	CreatedAt string // TODO: make this a time.Time
}

type Context struct{
	CreatedAt time.Time `json:"created_at"`
	ID string `json:"id"`
	Name string `json:"name"`
}

type ClientInterface interface {
	Contexts(vcs, org string) (*[]Context, error)
	ContextByName(vcs, org, name string) (*Context, error)
	DeleteContext(contextID string) error
	CreateContext(vcs, org, name string) (error)

	EnvironmentVariables(contextID string) (*[]EnvironmentVariable, error)
	CreateEnvironmentVariable(contextID, variable, value string) error
	DeleteEnvironmentVariable(contextID, variable string) error
}
