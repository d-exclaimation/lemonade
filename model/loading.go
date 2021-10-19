//
//  loading.go
//  model
//
//  Created by d-exclaimation on 11:52 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import (
	"github.com/d-exclaimation/lemonade/utils"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type tickMsg time.Time

type Loading struct {
	progress progress.Model
	await    *utils.Await
}

func NewLoading(op func()) Loading {
	return Loading{
		progress: progress.NewModel(progress.WithDefaultGradient()),
		await:    utils.Wait(op),
	}
}

func (m Loading) Init() tea.Cmd {
	return tickCmd()
}

func (m Loading) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		os.Exit(1)
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		var cmd tea.Cmd

		if m.await.IsDone() {
			cmd = m.progress.SetPercent(1.0)
		} else if m.progress.Percent() < 1.0-0.3 {
			cmd = m.progress.IncrPercent(0.25)
		}

		return m, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m Loading) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
