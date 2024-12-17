package models

import (
	"github.com/jinzhu/copier"
	"go.temporal.io/sdk/client"
)

func (c *ConnectClientParams) ToTemporalClientOptions() client.Options {
	options := client.Options{}
	copier.Copy(&options, &c.Options)
	return options
}

func (s *StartWorkflowParams) ToTemporalStartWorkflowOptions() client.StartWorkflowOptions {
	options := client.StartWorkflowOptions{}
	copier.Copy(&options, s.Options)
	return options
}
