package repository

import (
	"context"
	"log"
	"ticketing/ent"
	"ticketing/ent/concert"
	"ticketing/internal/common/cache"

	"github.com/google/uuid"
)

// ConcertRepository defines the interface for ticket-related data operations.
type ConcertRepository interface {
	GetAllConcerts(ctx context.Context) ([]*ent.Concert, error)
	GetConcertByID(ctx context.Context, concertID uuid.UUID) (*ent.Concert, error)
	GetConcertsByArtist(ctx context.Context, artistName string) ([]*ent.Concert, error)
	CreateConcert(ctx context.Context, title string, artist string, imageUrl string, description string, headline string, featured bool) (*ent.Concert, error)
	DeleteConcert(ctx context.Context, concertID uuid.UUID) error
	UpdateConcert(ctx context.Context, id uuid.UUID, title string, artist string, imageUrl string, description string, headline string, featured bool) (*ent.Concert, error)
	GetSessionsOfConcert(ctx context.Context, concertID uuid.UUID) ([]*ent.ConcertSession, error)
	DeleteAllConcerts(ctx context.Context) error
	GetFeaturedConcerts(ctx context.Context) ([]*ent.Concert, error)
}

// ConcertRepositoryImpl is an implementation of ConcertRepository.
type ConcertRepositoryImpl struct {
	Client      *ent.Client
	RedisClient *cache.RedisCache
}

// NewConcertRepository creates a new instance of ConcertRepositoryImpl.
func NewConcertRepository(client *ent.Client, redisClient *cache.RedisCache) ConcertRepository {
	return &ConcertRepositoryImpl{Client: client, RedisClient: redisClient}
}

func (r *ConcertRepositoryImpl) GetAllConcerts(ctx context.Context) ([]*ent.Concert, error) {
	// Define a key for caching
	cacheKey := "concerts"

	// Attempt to retrieve concerts from the cache
	var concerts []*ent.Concert
	err := r.RedisClient.GetEntity(ctx, cacheKey, &concerts)

	// If there's a cache miss, fetch from the database
	if err != nil {
		log.Printf("cache miss")
		// Query the database to fetch the concerts
		concerts, err = r.Client.Concert.
			Query().
			All(ctx)

		if err != nil {
			return nil, err
		}

		// Cache the fetched concerts
		err = r.RedisClient.SetEntityExpireIn(ctx, cacheKey, 30, concerts)
		if err != nil {
			// Handle the error (e.g., log it), but don't let it affect the response to the client
			log.Printf("Failed to cache concerts")
		}
	}
	log.Printf("cache hit")

	return concerts, nil
}

func (r *ConcertRepositoryImpl) GetConcertByID(ctx context.Context, concertID uuid.UUID) (*ent.Concert, error) {
	concert, err := r.Client.Concert.
		Query().
		Where(concert.IDEQ(concertID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return concert, nil
}

func (r *ConcertRepositoryImpl) GetConcertsByArtist(ctx context.Context, artistName string) ([]*ent.Concert, error) {
	concerts, err := r.Client.Concert.
		Query().
		Where(concert.ArtistEQ(artistName)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return concerts, nil
}

func (r *ConcertRepositoryImpl) CreateConcert(ctx context.Context, title string, artist string, imageUrl string, description string, headline string, featured bool) (*ent.Concert, error) {
	createdConcert, err := r.Client.Concert.
		Create().
		SetTitle(title).
		SetArtist(artist).
		SetImageUrl(imageUrl).
		SetDescription(description).
		SetHeadline(headline).
		SetFeatured(featured).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	err = r.RedisClient.DeleteEntity(ctx, "concerts")
	if err != nil {
		log.Printf("Failed to delete concerts cache")
	}

	return createdConcert, nil
}

func (r *ConcertRepositoryImpl) UpdateConcert(ctx context.Context, id uuid.UUID, title string, artist string, imageUrl string, description string, headline string, featured bool) (*ent.Concert, error) {
	// First, retrieve the existing concert by ID
	existingConcert, err := r.Client.Concert.
		Get(ctx, id)

	if err != nil {
		return nil, err
	}
	// Save the updated concert
	concert, err := existingConcert.
		Update().
		SetTitle(title).
		SetArtist(artist).
		SetImageUrl(imageUrl).
		SetDescription(description).
		SetHeadline(headline).
		SetFeatured(featured).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	err = r.RedisClient.DeleteEntity(ctx, "concerts")
	if err != nil {
		log.Printf("Failed to delete concerts cache")
	}

	return concert, nil
}

func (r *ConcertRepositoryImpl) DeleteConcert(ctx context.Context, concertID uuid.UUID) error {

	err := r.RedisClient.DeleteEntity(ctx, "concerts")
	if err != nil {
		log.Printf("Failed to delete concerts cache")
	}

	return r.Client.Concert.
		DeleteOneID(concertID).
		Exec(ctx)
}

func (r *ConcertRepositoryImpl) GetSessionsOfConcert(ctx context.Context, concertID uuid.UUID) ([]*ent.ConcertSession, error) {
	// Query the Concert by its ID
	concert, err := r.Client.Concert.
		Query().
		Where(concert.IDEQ(concertID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	// Query the sessions associated with the concert
	sessions, err := concert.
		QueryHasConcertSessions().
		WithOfConcert().
		All(ctx)

	return sessions, nil
}

func (r *ConcertRepositoryImpl) DeleteAllConcerts(ctx context.Context) error {
	if _, err := r.Client.Concert.
		Delete().
		Exec(ctx); err != nil {
		// Handle the error if the deletion fails.
		log.Fatalf("failed to delete session: %v", err)
		return err
	}
	return nil
}

func (r *ConcertRepositoryImpl) GetFeaturedConcerts(ctx context.Context) ([]*ent.Concert, error) {
	return r.Client.Concert.
		Query().
		Where(concert.Featured(true)).
		All(ctx)
}
