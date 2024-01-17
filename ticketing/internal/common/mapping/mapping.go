package mapping

import (
	"ticketing/ent"
	"ticketing/internal/common/types"
)

func FromEntTicket(e *ent.Ticket) (*types.Ticket, error) {
	return &types.Ticket{
		ID:         e.ID,
		SeatNumber: e.SeatNumber,
		Status:     e.Status.String(),
		UserID:     e.UserId,
	}, nil

}

func FromEntTicketList(e []*ent.Ticket) ([]*types.Ticket, error) {
	responseTickets := make([]*types.Ticket, len(e))
	for i, ticket := range e {
		t, err := FromEntTicket(ticket)
		if err != nil {
			return nil, err // Return immediately on error. Adjust as needed.
		}
		responseTickets[i] = t
	}
	return responseTickets, nil
}
