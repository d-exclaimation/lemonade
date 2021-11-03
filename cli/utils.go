//
//  utils.go.go
//  cli
//
//  Created by d-exclaimation on 9:06 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package cli

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"os/exec"
)

// Tea
//
// Execute a model in Bubble-tea CLI
func Tea(model tea.Model) {
	p1 := tea.NewProgram(model, tea.WithAltScreen())
	if err := p1.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// Run
//
// Execute `/bin/sh` command but not wait for it.
func Run(s string, arg ...string) *exec.Cmd {
	cmd := exec.Command(s, arg...)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err.Error())
		return nil
	}
	return cmd
}

// RunUnder
//
// Execute `/bin/sh` command but not wait for it.
func RunUnder(directory, s string, arg ...string) *exec.Cmd {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}
	cmd := exec.Command(s, arg...)
	cmd.Dir = dir + "/" + directory
	if err := cmd.Run(); err != nil {
		log.Fatalln(err.Error())
		return nil
	}
	return cmd
}

func Write(name string, content string) error {
	f, err := os.Create(name)
	defer f.Close()
	_, err = f.Write([]byte(content))
	return err
}

func Move(file1, file2 string) error {
	return os.Rename(file1, file2)
}
