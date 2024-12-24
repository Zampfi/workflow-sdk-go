package activity

import (
	temporalactivity "go.temporal.io/sdk/activity"
)

type Activity struct {
	Function        interface{}
	RegisterOptions RegisterActivityOptions
}

type RegisterActivityOptions = temporalactivity.RegisterOptions
