package repository

import (
	"context"
	"log"
	"ticketing/ent"
	"ticketing/ent/concert"
	"ticketing/ent/concertsession"
	"ticketing/ent/section"
	"ticketing/ent/ticket"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// SessionRepository defines the interface for ticket-related data operations.
type SessionRepository interface {
	GetAllSessions(ctx context.Context) ([]*ent.ConcertSession, error)
	GetSessionByID(ctx context.Context, sessionID uuid.UUID) (*ent.ConcertSession, error)
	CreateConcertSession(ctx context.Context, concertID uuid.UUID, datetime int) (*ent.ConcertSession, error)
	UpdateConcertSession(ctx context.Context, sessionID uuid.UUID, sessionDateTime int) (*ent.ConcertSession, error)
	DeleteConcertSession(ctx context.Context, sessionID uuid.UUID) error
	GetSectionsOfAConcertSession(ctx context.Context, sessionID uuid.UUID) ([]*ent.Section, error)
	DeleteAllSessions(ctx context.Context) error
}

// SessionRepositoryImpl is an implementation of SessionRepository.
type SessionRepositoryImpl struct {
	Client      *ent.Client
	RedisClient *redis.Client
}

// NewSessionRepository creates a new instance of SessionRepositoryImpl.
func NewSessionRepository(client *ent.Client) SessionRepository {
	return &SessionRepositoryImpl{Client: client}
}

func (r *SessionRepositoryImpl) GetAllSessions(ctx context.Context) ([]*ent.ConcertSession, error) {
	sessions, err := r.Client.ConcertSession.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (r *SessionRepositoryImpl) GetSessionByID(ctx context.Context, sessionID uuid.UUID) (*ent.ConcertSession, error) {
	session, err := r.Client.ConcertSession.
		Query().
		Where(concertsession.IDEQ(sessionID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *SessionRepositoryImpl) CreateConcertSession(ctx context.Context, concertID uuid.UUID, datetime int) (*ent.ConcertSession, error) {
	concert, err := r.Client.Concert.
		Query().
		Where(concert.IDEQ(concertID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	createdSession, err := r.Client.ConcertSession.
		Create().
		SetSessionDateTime(datetime).
		AddOfConcert(concert).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return createdSession, nil
}

func (r *SessionRepositoryImpl) UpdateConcertSession(ctx context.Context, sessionID uuid.UUID, sessionDateTime int) (*ent.ConcertSession, error) {
	existingConcerSession, err := r.Client.ConcertSession.
		Get(ctx, sessionID)

	if err != nil {
		return nil, err
	}
	concerSession, err := existingConcerSession.
		Update().
		SetSessionDateTime(int(sessionDateTime)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return concerSession, nil
}

func (r *SessionRepositoryImpl) DeleteConcertSession(ctx context.Context, sessionID uuid.UUID) error {
	// Fetch related Section entities
	sections, err := r.Client.Section.
		Query().
		Where(section.HasAtConcertSessionWith(concertsession.IDEQ(sessionID))).
		All(ctx)
	if err != nil {
		return err
	}

	// Delete related Ticket entities for each Section
	for _, selectedSection := range sections {
		_, err := r.Client.Ticket.
			Delete().
			Where(ticket.HasWithinSectionWith(section.IDEQ(selectedSection.ID))).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	// Delete the Section entities
	_, err = r.Client.Section.
		Delete().
		Where(section.HasAtConcertSessionWith(concertsession.IDEQ(sessionID))).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Finally, delete the ConcertSession entity
	_, err = r.Client.ConcertSession.
		Delete().
		Where(concertsession.IDEQ(sessionID)).
		Exec(ctx)
	return err
}

func (r *SessionRepositoryImpl) GetSectionsOfAConcertSession(ctx context.Context, sessionID uuid.UUID) ([]*ent.Section, error) {
	session, err := r.Client.ConcertSession.
		Query().
		Where(concertsession.IDEQ(sessionID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	sections, err := session.
		QueryHasSections().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return sections, nil
}

func (r *SessionRepositoryImpl) DeleteAllSessions(ctx context.Context) error {
	if _, err := r.Client.ConcertSession.
		Delete().
		Exec(ctx); err != nil {
		// Handle the error if the deletion fails.
		log.Fatalf("failed to delete session: %v", err)
		return err
	}
	return nil
}
