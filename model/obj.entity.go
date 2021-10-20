//
//  obj.entity.go
//  model
//
//  Created by d-exclaimation on 10:48 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

// Object type for the Bubbles' List item
type obj struct {
	title, desc string // Title and description
}

func (i obj) Title() string       { return i.title }
func (i obj) Description() string { return i.desc }
func (i obj) FilterValue() string { return i.title }
