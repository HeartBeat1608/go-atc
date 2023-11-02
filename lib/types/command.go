package types

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Command struct {
	ID         string
	Request    string
	Response   string
	RequestAt  int64
	ResponseAt int64
}

func NewCommand(request string) (Command, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return Command{}, err
	}
	return Command{
		ID:         id.String(),
		Request:    request,
		Response:   "",
		ResponseAt: -1,
		RequestAt:  time.Now().UnixMilli(),
	}, nil
}

func (c *Command) Respond(response string) {
	c.Response = response
	c.ResponseAt = time.Now().UnixMilli()
}

func (c *Command) CopyFrom(cmd Command) {
	c.Request = cmd.Request
	c.Response = cmd.Response
	c.ResponseAt = cmd.ResponseAt
	c.RequestAt = cmd.RequestAt
}

func (c Command) Prettify() string {
	return fmt.Sprintf(
		"Command\nID:\t\t%s\nRequest:\t%s\nRequested At:\t%d\nResponse:\t%s\nResponded At:\t%d\n---",
		c.ID, c.Request, c.RequestAt, c.Response, c.ResponseAt,
	)
}
