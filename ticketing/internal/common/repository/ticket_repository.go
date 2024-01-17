package repository

import (
	"context"
	"log"
	"ticketing/ent"
	"ticketing/ent/section"
	"ticketing/ent/ticket"
	"ticketing/internal/common/cache"
	"time"

	"github.com/google/uuid"
)

const (
	ExpiryMin = 10
)

// TicketRepository defines the interface for ticket-related data operations.
type TicketRepository interface {
	GetAllTickets(ctx context.Context) ([]*ent.Ticket, error)
	GetTicketById(ctx context.Context, ticketID uuid.UUID) (*ent.Ticket, error)
	GetTicketsByUserID(ctx context.Context, userId uuid.UUID) ([]*ent.Ticket, error)
	CreateTicket(ctx context.Context, sectionID uuid.UUID, seatNumber int, status ticket.Status) (*ent.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *ent.Ticket) error
	UpdateTicketStatus(ctx context.Context, ticketID uuid.UUID, status ticket.Status, userID uuid.UUID) (*ent.Ticket, error)
	DeleteTicket(ctx context.Context, ticketID uuid.UUID) error
	DeleteAllTickets(ctx context.Context) error
	ReserveTicket(ctx context.Context, sectionID uuid.UUID, seatID uuid.UUID, userID uuid.UUID) error
}

// TicketRepositoryImpl is an implementation of TicketRepository.
type TicketRepositoryImpl struct {
	Client      *ent.Client
	RedisClient *cache.RedisCache
}

// NewTicketRepository creates a new instance of TicketRepositoryImpl.
func NewTicketRepository(client *ent.Client, redisClient *cache.RedisCache) TicketRepository {
	return &TicketRepositoryImpl{Client: client, RedisClient: redisClient}
}

// GetAllTickets retrieves all tickets from the database.
func (r *TicketRepositoryImpl) GetAllTickets(ctx context.Context) ([]*ent.Ticket, error) {

	tickets, err := r.Client.Ticket.Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *TicketRepositoryImpl) GetTicketById(ctx context.Context, ticketID uuid.UUID) (*ent.Ticket, error) {

	ticket, err := r.Client.Ticket.
		Query().
		Where(ticket.IDEQ(ticketID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *TicketRepositoryImpl) GetTicketsByUserID(ctx context.Context, userId uuid.UUID) ([]*ent.Ticket, error) {

	tickets, err := r.Client.Ticket.
		Query().
		Where(ticket.UserIdEQ(userId)).
		WithWithinSection(func(q *ent.SectionQuery) {
			q.WithAtConcertSession(func(sq *ent.ConcertSessionQuery) {
				sq.WithOfConcert() // Load associated Concert for each Session
			})
		}).
		All(ctx)

	r.Client.Section.
		Query().
		Where()

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *TicketRepositoryImpl) CreateTicket(ctx context.Context, sectionID uuid.UUID, seatNumber int, status ticket.Status) (*ent.Ticket, error) {

	section, err := r.Client.Section.
		Query().
		Where(section.IDEQ(sectionID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	createdTicket, err := r.Client.Ticket.
		Create().
		SetSeatNumber(seatNumber).
		SetStatus(status).
		SetWithinSection(section).
		SetUserId(uuid.Nil).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return createdTicket, nil
}

func (r *TicketRepositoryImpl) UpdateTicket(ctx context.Context, ticket *ent.Ticket) error {
	_, err := r.Client.Ticket.
		UpdateOne(ticket).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *TicketRepositoryImpl) UpdateTicketStatus(ctx context.Context, ticketID uuid.UUID, status ticket.Status, userID uuid.UUID) (*ent.Ticket, error) {

	ticket, err := r.Client.Ticket.
		UpdateOneID(ticketID).
		SetStatus(status).
		SetUserId(userID).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *TicketRepositoryImpl) ReserveTicket(ctx context.Context, sectionID uuid.UUID, seatID uuid.UUID, userID uuid.UUID) error {

	_, err := r.Client.Ticket.
		UpdateOneID(seatID).
		SetStatus(ticket.StatusRESERVED).
		SetUserId(userID).
		SetReservedAt(int(time.Now().Unix())).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *TicketRepositoryImpl) DeleteTicket(ctx context.Context, ticketID uuid.UUID) error {
	return r.Client.Ticket.
		DeleteOneID(ticketID).
		Exec(ctx)
}

func (r *TicketRepositoryImpl) DeleteAllTickets(ctx context.Context) error {
	if _, err := r.Client.Ticket.
		Delete().
		Exec(ctx); err != nil {
		// Handle the error if the deletion fails.
		log.Fatalf("failed to delete tickets: %v", err)
		return err
	}
	return nil
}
