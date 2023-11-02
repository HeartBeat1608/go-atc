package main

import (
	"fmt"
	"go_atc/lib"
	"go_atc/lib/types"
	"go_atc/lib/utils"
)

var SYSTEM_PROMPT = "You are an expert Air Traffic Controller, managing flights to and from the London City Airport Tower.  You coordinate various aircrafts and command them with appropriate signals towards their journey into or out of the airport. Make sure to speak numbers in digits. You will hereafter act as the ATC and communicate in aviation terms with me."

func main() {
	metar, err := lib.NewMetar("EGLL")
	utils.HandleError(err)

	if err := metar.Refresh(); err != nil {
		utils.HandleError(err)
	}
	fmt.Println(metar.Pretty())

	deepinfra, err := lib.NewDeepInfra()
	utils.HandleError(err)

	manager, err := types.NewSessionManager(SYSTEM_PROMPT)
	utils.HandleError(err)

	cmd1, err := types.NewCommand("Air Asia 113, Requesting pushback and start.")
	utils.HandleError(err)
	manager.AddCommand(cmd1)

	manager.RespondToLastCommand("Air Asia 113, Cleared to pushback. Start engines. Report when ready to taxi.")

	cmd2, err := types.NewCommand("Air Asia 113, Requesting taxi to runway 21 Left.")
	utils.HandleError(err)
	manager.AddCommand(cmd2)
	prompt := manager.BuildPrompt(metar)
	res, err := deepinfra.Complete(prompt)
	utils.HandleError(err)
	manager.RespondToLastCommand(res.Results[0].Text)

	cmd1, err = types.NewCommand("Air Asia 113, Requesting weather information.")
	utils.HandleError(err)
	manager.AddCommand(cmd1)
	prompt = manager.BuildPrompt(metar)
	res, err = deepinfra.Complete(prompt)
	utils.HandleError(err)
	manager.RespondToLastCommand(res.Results[0].Text)

	fmt.Println(prompt)
	fmt.Println(manager.PrettyCommands())
}
