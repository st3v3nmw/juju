// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package fortress

import (
	"context"
	"sync"

	"gopkg.in/tomb.v2"
)

// fortress coordinates between clients that access it as a Guard and as a Guest.
type fortress struct {
	tomb         tomb.Tomb
	guardTickets chan guardTicket
	guestTickets chan guestTicket
}

// newFortress returns a new, locked, fortress. The caller is responsible for
// ensuring it somehow gets Kill()ed, and for handling any error returned by
// Wait().
func newFortress() *fortress {
	f := &fortress{
		guardTickets: make(chan guardTicket),
		guestTickets: make(chan guestTicket),
	}
	f.tomb.Go(f.loop)
	return f
}

// Kill is part of the worker.Worker interface.
func (f *fortress) Kill() {
	f.tomb.Kill(nil)
}

// Wait is part of the worker.Worker interface.
func (f *fortress) Wait() error {
	return f.tomb.Wait()
}

// Unlock is part of the Guard interface.
func (f *fortress) Unlock(ctx context.Context) error {
	return f.allowGuests(ctx, true)
}

// Lockdown is part of the Guard interface.
func (f *fortress) Lockdown(ctx context.Context) error {
	return f.allowGuests(ctx, false)
}

// Visit is part of the Guest interface.
func (f *fortress) Visit(ctx context.Context, visit Visit) error {
	result := make(chan error)
	select {
	case <-f.tomb.Dying():
		return ErrShutdown
	case <-ctx.Done():
		return ErrAborted
	case f.guestTickets <- guestTicket{
		ctx:    ctx,
		visit:  visit,
		result: result,
	}:
		return <-result
	}
}

// allowGuests communicates Guard-interface requests to the main loop.
func (f *fortress) allowGuests(ctx context.Context, allowGuests bool) error {
	result := make(chan error)
	select {
	case <-f.tomb.Dying():
		return ErrShutdown
	case f.guardTickets <- guardTicket{
		ctx:         ctx,
		allowGuests: allowGuests,
		result:      result,
	}:
		return <-result
	}
}

// loop waits for a Guard to unlock the fortress, and then runs visit funcs in
// parallel until a Guard locks it down again; at which point, it waits for all
// outstanding visits to complete, and reverts to its original state.
func (f *fortress) loop() error {
	var active sync.WaitGroup
	defer active.Wait()

	// guestTickets will be set on Unlock and cleared at the start of Lockdown.
	var guestTickets <-chan guestTicket
	for {
		select {
		case <-f.tomb.Dying():
			return tomb.ErrDying
		case ticket := <-guestTickets:
			active.Add(1)
			go ticket.complete(active.Done)
		case ticket := <-f.guardTickets:
			// guard ticket requests are idempotent; it's not worth building
			// the extra mechanism needed to (1) complain about abuse but
			// (2) remain comprehensible and functional in the face of aborted
			// Lockdowns.
			if ticket.allowGuests {
				guestTickets = f.guestTickets
			} else {
				guestTickets = nil
			}
			go ticket.complete(active.Wait)
		}
	}
}

// guardTicket communicates between the Guard interface and the main loop.
type guardTicket struct {
	ctx         context.Context
	allowGuests bool
	result      chan<- error
}

// complete unconditionally sends a single value on ticket.result; either nil
// (when the desired state is reached) or ErrAborted (when the ticket's ctx is
// done). It should be called on its own goroutine.
func (ticket guardTicket) complete(waitLockedDown func()) {
	var result error
	defer func() {
		ticket.result <- result
	}()

	done := make(chan struct{})
	go func() {
		// If we're locking down, we should wait for all Visits to complete.
		// If not, Visits are already being accepted and we're already done.
		if !ticket.allowGuests {
			waitLockedDown()
		}
		close(done)
	}()
	select {
	case <-done:
	case <-ticket.ctx.Done():
		result = ErrAborted
	}
}

// guestTicket communicates between the Guest interface and the main loop.
type guestTicket struct {
	ctx    context.Context
	visit  Visit
	result chan<- error
}

// complete unconditionally sends any error returned from the Visit func, then
// calls the finished func. It should be called on its own goroutine.
func (ticket guestTicket) complete(finished func()) {
	defer finished()

	select {
	case <-ticket.ctx.Done():
		ticket.result <- ErrAborted
	default:
		ticket.result <- ticket.visit()
	}
}
