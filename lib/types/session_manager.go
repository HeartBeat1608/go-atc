package types

import (
	"fmt"
	"go_atc/lib"
	"go_atc/lib/utils"
	"time"

	"github.com/google/uuid"
)

type SessionManager struct {
	ID       string
	Commands []Command
	StartAt  int64
	Prompt   string
}

func NewSessionManager(prompt string) (SessionManager, error) {
	id, err := uuid.NewUUID()
	utils.HandleError(err)

	return SessionManager{
		ID:       id.String(),
		Commands: []Command{},
		StartAt:  time.Now().UnixMilli(),
		Prompt:   prompt,
	}, nil
}

func (s SessionManager) BuildPrompt(m *lib.Metar) string {
	prompt := ""
	metar := m.GetStringifiedWeather()

	for i := range s.Commands {
		cmd := s.Commands[i]
		if i > 0 {
			prompt += fmt.Sprintf("[INST] %s [/INST] %s\n", cmd.Request, cmd.Response)
		} else {
			prompt += fmt.Sprintf("[INST] <<SYS>>\n%s\nMetar Information for the airport is\n%s\n<<SYS>>\n\n%s [/INST] %s\n", s.Prompt, metar, cmd.Request, cmd.Response)
		}
	}

	return prompt
}

func (s *SessionManager) AddCommand(command Command) {
	s.Commands = append(s.Commands, command)
}

func (s *SessionManager) GetLastCommand() Command {
	return s.Commands[len(s.Commands)-1]
}

func (s *SessionManager) UpdateLastCommand(cmd Command) {
	s.Commands[len(s.Commands)-1].CopyFrom(cmd)
}

func (s *SessionManager) RespondToLastCommand(response string) {
	s.Commands[len(s.Commands)-1].Respond(response)
}

func (s *SessionManager) Prettify() string {
	return fmt.Sprintf("SessionManager\nID:\t\t%s\nN Commands:\t%d\nStarted At:\t%d\n---", s.ID, len(s.Commands), s.StartAt)
}

func (s *SessionManager) PrettyCommands() string {
	res := ""

	for i := range s.Commands {
		cmd := s.Commands[i]
		res += fmt.Sprintf("[%d] Pilot: %s\n[%d] ATC  : %s\n", cmd.RequestAt, cmd.Request, cmd.ResponseAt, cmd.Response)
	}

	return res
}
