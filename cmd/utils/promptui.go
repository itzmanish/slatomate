package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type PromptType int

const (
	TextPrompt PromptType = iota
	PasswordPrompt
)

// PromptContent defines content for prompt ui
type PromptContent struct {
	ErrorMsg string
	Label    string
	Type     PromptType
}

// PromptGetInput prompts for input
func PromptGetInput(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label + " :",
		Templates: templates,
		Validate:  validate,
	}
	if pc.Type == PasswordPrompt {
		prompt.Mask = '*'
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
