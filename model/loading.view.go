//
//  loading.view.go
//  model
//
//  Created by d-exclaimation on 11:52 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import (
	"github.com/d-exclaimation/lemonade/future"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

// LoadingView
//
// Progress indicator that wait for a process to finish
type LoadingView struct {
	progress progress.Model // Inner animated progress bar
	await    *future.Await  // Await to check progress on process
	scale    float64        // Scala in which percentage increases
}

// NewLoadingView
//
// Construct a new LoadingView
func NewLoadingView(op func()) LoadingView {
	return LoadingView{
		progress: progress.NewModel(progress.WithDefaultGradient()),
		await:    future.Wait(op),
		scale:    0.25,
	}
}

// Init
//
// Initial state for the bubble-tea cli
func (m LoadingView) Init() tea.Cmd {
	return ticking()
}

// Update
//
// Render update for bubble-tea
func (m LoadingView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	case time.Time:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		var cmd tea.Cmd

		if m.await.IsDone() {
			cmd = m.progress.SetPercent(1.0)
		} else if m.progress.Percent() < 1.0-(m.scale+0.1) {
			cmd = m.progress.IncrPercent(m.scale)
		}

		m.scale *= 0.75

		return m, tea.Batch(ticking(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

// View
//
// Render TextField as string
func (m LoadingView) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

// Ticking function
func ticking() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return t
	})
}
