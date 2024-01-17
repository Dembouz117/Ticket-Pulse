package types

import "github.com/google/uuid"

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type Concert struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Artist   string    `json:"artist"`
	ImageURL string    `json:"imageUrl"`
}

type Session struct {
	ID              uuid.UUID `json:"id"`
	SessionDateTime int       `json:"sessionDateTime"`
}

type Section struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Capacity int       `json:"capacity"`
	Reserved int       `json:"reserved"`
	Bought   int       `json:"bought"`
	Category string    `json:"category"`
	Price    int       `json:"price"`
}

type Ticket struct {
	ID         uuid.UUID `json:"id"`
	SeatNumber int       `json:"seatNumber"`
	Status     string    `json:"status"`
	UserID     uuid.UUID `json:"userId"`
}

// Stores the concerts list only
type ConcertListResponse struct {
	Concert []Concert `json:"concerts"`
}

// Stores the concert only
type ConcertResponse struct {
	Concert Concert `json:"concert"`
}

type ConcertWithSessionListResponse struct {
	Concert  string    `json:"concert"`
	Sessions []Session `json:"sessions"`
}

type SessionListResponse struct {
	Session []Session `json:"sessions"`
}

type SessionResponse struct {
	Session Session `json:"session"`
}

type SessionWithSectionListResponse struct {
	Session  string    `json:"session"`
	Sections []Section `json:"sections"`
}

type SectionListResponse struct {
	Section []Section `json:"sections"`
}

type SectionResponse struct {
	Section Section `json:"section"`
}

type SectionWithTicketsListResponse struct {
	Section string   `json:"section"`
	Tickets []Ticket `json:"tickets"`
}
