package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sagikazarmark.dev/mga/internal/generate/gentypes"
)

func TestGenerate(t *testing.T) {
	file := File{
		File: gentypes.File{
			Package: gentypes.PackageRef{
				Name: "pkggen",
				Path: "app.dev/pkg/pkggen",
			},
			HeaderText: `// Copyright 2020 Acme Inc.
// All rights reserved.
//
// Licensed under "Only for testing purposes" license.
`,
		},
		EventHandlers: []EventHandler{
			{
				Name: "Event",
				Event: Event{
					Name: "Event",
					Package: gentypes.PackageRef{
						Name: "pkg",
						Path: "app.dev/pkg",
					},
				},
			},
		},
	}

	expected := `// +build !ignore_autogenerated

// Copyright 2020 Acme Inc.
// All rights reserved.
//
// Licensed under "Only for testing purposes" license.

// Code generated by mga tool. DO NOT EDIT.

package pkggen

import (
	"app.dev/pkg"
	"context"
	"emperror.dev/errors"
	"fmt"
)

// EventHandler handles Event events.
type EventHandler interface {
	// Event handles a(n) Event event.
	Event(ctx context.Context, event pkg.Event) error
}

// EventEventHandler handles Event events.
type EventEventHandler struct {
	handler EventHandler
	name    string
}

// NewEventEventHandler returns a new EventEventHandler instance.
func NewEventEventHandler(handler EventHandler, name string) EventEventHandler {
	return EventEventHandler{
		handler: handler,
		name:    name,
	}
}

// HandlerName returns the name of the event handler.
func (h EventEventHandler) HandlerName() string {
	return h.name
}

// NewEvent returns a new empty event used for serialization.
func (h EventEventHandler) NewEvent() interface{} {
	return &pkg.Event{}
}

// Handle handles an event.
func (h EventEventHandler) Handle(ctx context.Context, event interface{}) error {
	e, ok := event.(*pkg.Event)
	if !ok {
		return errors.NewWithDetails("unexpected event type", "type", fmt.Sprintf("%T", event))
	}

	return h.handler.Event(ctx, *e)
}
`

	actual, err := Generate(file)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, string(actual), "the generated code does not match the expected one")
}
