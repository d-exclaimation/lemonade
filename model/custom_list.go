//
//  custom_list.go.go
//  model
//
//  Created by d-exclaimation on 10:30 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
)

// NewModel returns a new model with sensible defaults.
func NewModel(items []list.Item) list.Model {
	styles := list.DefaultStyles()
	l := list.NewModel(items, list.NewDefaultDelegate(), 0, 0)
	l.Paginator.Type = paginator.Dots
	l.Paginator.ActiveDot = styles.ActivePaginationDot.String()
	l.Paginator.InactiveDot = styles.InactivePaginationDot.String()
	l.Paginator.PerPage = 10
	return l
}
