// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build windows
// +build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

import (
	"fmt"
)

// Event is an event stored in windows event log.
type Event struct {
	handle uintptr
}

// RenderSimple will render the event as EventXML without formatted info.
func (e *Event) RenderSimple(buffer Buffer) (EventXML, error) {
	if e.handle == 0 {
		return EventXML{}, fmt.Errorf("event handle does not exist")
	}

	bufferUsed, _, err := evtRender(0, e.handle, EvtRenderEventXML, buffer.SizeBytes(), buffer.FirstByte())
	if err == ErrorInsufficientBuffer {
		buffer.UpdateSizeBytes(*bufferUsed)
		return e.RenderSimple(buffer)
	}

	if err != nil {
		return EventXML{}, fmt.Errorf("syscall to 'EvtRender' failed: %w", err)
	}

	bytes, err := buffer.ReadBytes(*bufferUsed)
	if err != nil {
		return EventXML{}, fmt.Errorf("failed to read bytes from buffer: %w", err)
	}

	return unmarshalEventXML(bytes)
}

// RenderFormatted will render the event as EventXML with formatted info.
func (e *Event) RenderFormatted(buffer Buffer, publisher Publisher) (EventXML, error) {
	if e.handle == 0 {
		return EventXML{}, fmt.Errorf("event handle does not exist")
	}

	var bufferUsed uint32
	err := evtFormatMessage(publisher.handle, e.handle, 0, 0, 0, EvtFormatMessageXML, buffer.SizeWide(), buffer.FirstByte(), &bufferUsed)
	if err == ErrorInsufficientBuffer {
		buffer.UpdateSizeWide(bufferUsed)
		return e.RenderFormatted(buffer, publisher)
	}

	if err != nil {
		return EventXML{}, fmt.Errorf("syscall to 'EvtFormatMessage' failed: %w", err)
	}

	bytes, err := buffer.ReadWideChars(bufferUsed)
	if err != nil {
		return EventXML{}, fmt.Errorf("failed to read bytes from buffer: %w", err)
	}

	return unmarshalEventXML(bytes)
}

// Close will close the event handle.
func (e *Event) Close() error {
	if e.handle == 0 {
		return nil
	}

	if err := evtClose(e.handle); err != nil {
		return fmt.Errorf("failed to close event handle: %w", err)
	}

	e.handle = 0
	return nil
}

func (e *Event) RenderRaw(buffer Buffer) (EventRaw, error) {
	if e.handle == 0 {
		return EventRaw{}, fmt.Errorf("event handle does not exist")
	}

	bufferUsed, _, err := evtRender(0, e.handle, EvtRenderEventXML, buffer.SizeBytes(), buffer.FirstByte())
	if err == ErrorInsufficientBuffer {
		buffer.UpdateSizeBytes(*bufferUsed)
		return e.RenderRaw(buffer)
	}
	bytes, err := buffer.ReadWideChars(*bufferUsed)
	if err != nil {
		return EventRaw{}, fmt.Errorf("failed to read bytes from buffer: %w", err)
	}

	return unmarshalEventRaw(bytes)
}

// NewEvent will create a new event from an event handle.
func NewEvent(handle uintptr) Event {
	return Event{
		handle: handle,
	}
}
