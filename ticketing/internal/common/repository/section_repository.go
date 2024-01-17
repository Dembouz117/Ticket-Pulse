package repository

import (
	"context"
	"log"
	"ticketing/ent"
	"ticketing/ent/concertsession"
	"ticketing/ent/section"
	"ticketing/ent/ticket"
	"ticketing/internal/common/cache"
	"time"

	"github.com/google/uuid"
)

type SectionRepository interface {
	GetAllSections(ctx context.Context) ([]*ent.Section, error)
	GetSectionByID(ctx context.Context, sectionID uuid.UUID) (*ent.Section, error)
	CreateSection(ctx context.Context, sessionID uuid.UUID, name string, capacity int, bought int, reserved int, category section.Category, price int) (*ent.Section, error)
	UpdateSection(ctx context.Context, sectionID uuid.UUID, name string, capacity int, reserved int, bought int, category section.Category, price int) (*ent.Section, error)
	DeleteSection(ctx context.Context, sectionID uuid.UUID) error
	GetAvailableTicketsBySectionID(ctx context.Context, sectionID uuid.UUID) ([]*ent.Ticket, error)
	GetTicketsBySection(ctx context.Context, sectionID uuid.UUID) ([]*ent.Ticket, error)
	DeleteAllSections(ctx context.Context) error
	ReleaseExpiredTicketsBySectionID(ctx context.Context, sectionID uuid.UUID) error
}

type SectionRepositoryImpl struct {
	Client      *ent.Client
	RedisClient *cache.RedisCache
}

func NewSectionRepository(client *ent.Client, redisClient *cache.RedisCache) SectionRepository {
	return &SectionRepositoryImpl{Client: client, RedisClient: redisClient}
}

func (r *SectionRepositoryImpl) GetAllSections(ctx context.Context) ([]*ent.Section, error) {
	sections, err := r.Client.Section.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return sections, nil
}

func (r *SectionRepositoryImpl) GetSectionByID(ctx context.Context, sectionID uuid.UUID) (*ent.Section, error) {
	log.Printf("%s", sectionID.String())
	section, err := r.Client.Section.
		Query().
		Where(section.IDEQ(sectionID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return section, nil
}

func (r *SectionRepositoryImpl) CreateSection(ctx context.Context, sessionID uuid.UUID, name string, capacity int, bought int, reserved int, category section.Category, price int) (*ent.Section, error) {

	concertSession, err := r.Client.ConcertSession.
		Query().
		Where(concertsession.IDEQ(sessionID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	createdSection, err := r.Client.Section.
		Create().
		SetName(name).
		SetCapacity(capacity).
		SetCategory(category).
		SetReserved(reserved).
		SetBought(bought).
		SetPrice(price).
		SetAtConcertSession(concertSession).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return createdSection, nil
}

func (r *SectionRepositoryImpl) UpdateSection(ctx context.Context, sectionID uuid.UUID, name string, capacity int, reserved int, bought int, category section.Category, price int) (*ent.Section, error) {

	section, err := r.Client.Section.
		UpdateOneID(sectionID).
		SetName(name).
		SetCapacity(capacity).
		SetReserved(reserved).
		SetBought(bought).
		SetCategory(category).
		SetPrice(price).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return section, nil
}

func (r *SectionRepositoryImpl) DeleteSection(ctx context.Context, sectionID uuid.UUID) error {

	_, err := r.Client.Ticket.
		Delete().
		Where(ticket.HasWithinSectionWith(section.IDEQ(sectionID))).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Delete the Section entities
	_, err = r.Client.Section.
		Delete().
		Where(section.IDEQ(sectionID)).
		Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func (r *SectionRepositoryImpl) GetAvailableTicketsBySectionID(ctx context.Context, sectionID uuid.UUID) ([]*ent.Ticket, error) {
	tickets, err := r.Client.Ticket.
		Query().
		Where(ticket.HasWithinSectionWith(section.IDEQ(sectionID)),
			ticket.StatusEQ(ticket.StatusAVAILABLE)).
		WithWithinSection().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *SectionRepositoryImpl) ReleaseExpiredTicketsBySectionID(ctx context.Context, sectionID uuid.UUID) error {

	// Define the duration for which a reservation is valid.
	reservationValidDuration := 10 * time.Minute

	// Calculate the cutoff time.
	cutoffTime := time.Now().Add(-reservationValidDuration)
	log.Printf("cut off time: %d", cutoffTime.Unix())

	// Find all tickets that have a ReservedAt timestamp older than the cutoff time.
	expiredTickets, err := r.Client.Ticket.
		Query().
		Where(
			ticket.ReservedAtLT(int(cutoffTime.Unix())),
			ticket.StatusEQ(ticket.StatusRESERVED),
		).
		All(ctx)

	if err != nil {
		return err
	}

	log.Printf("expire tickets: %s", expiredTickets)
	for _, expiredTicket := range expiredTickets {
		log.Printf("expire ticket: %s", expiredTicket.ID)
		_, err := r.Client.Ticket.
			UpdateOneID(expiredTicket.ID).
			SetStatus(ticket.StatusAVAILABLE).
			SetUserId(uuid.Nil).
			ClearReservedAt().
			Save(ctx)
		if err != nil {
			log.Printf("Cannot release ticket: %s", expiredTicket.ID)
		}
	}

	return nil
}

func (r *SectionRepositoryImpl) GetTicketsBySection(ctx context.Context, sectionID uuid.UUID) ([]*ent.Ticket, error) {
	log.Printf("get tickets by section: %s", sectionID.String())
	tickets, err := r.Client.Section.
		Query().
		Where(section.IDEQ(sectionID)).
		QueryHasTickets().
		All(ctx)
	log.Printf("get tickets by section: %d, %s", len(tickets), err)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *SectionRepositoryImpl) DeleteAllSections(ctx context.Context) error {

	if _, err := r.Client.Section.
		Delete().
		Exec(ctx); err != nil {
		log.Fatalf("failed to delete session: %v", err)
		return err
	}

	return nil
}
