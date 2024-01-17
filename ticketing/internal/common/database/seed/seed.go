package seed

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"ticketing/ent/section"
	"ticketing/ent/ticket"
	"ticketing/internal/common/repository"
)

var seedData []struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Description string `json:"description"`
	Headline    string `json:"headline"`
	ImageUrl    string `json:"imageUrl"`
	Featured    bool   `json:"featured"`
	Sessions    []struct {
		SessionDateTime int `json:"sessionDateTime"`
		Sections        []struct {
			Name     string `json:"name"`
			Capacity int    `json:"capacity"`
			Category string `json:"category"`
			Price    int    `json:"price"`
		}
	}
}

func SeedDatabase(ctx context.Context, concertRepo repository.ConcertRepository, sessionRepo repository.SessionRepository,
	sectionRepo repository.SectionRepository, ticketRepo repository.TicketRepository, filePath string) error {
	// Open and read the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open JSON file: %v", err)
		return err
	}
	defer file.Close()

	// Decode the JSON data into the seedData slice
	if err := json.NewDecoder(file).Decode(&seedData); err != nil {
		log.Fatalf("Failed to decode JSON data: %v", err)
	}

	// Seed the concerts
	for _, seedConcert := range seedData {

		concert, err := concertRepo.CreateConcert(ctx, seedConcert.Title, seedConcert.Artist, seedConcert.ImageUrl, seedConcert.Description, seedConcert.Headline, seedConcert.Featured)
		log.Printf("Created concert: %s", concert.Title)
		if err != nil {
			log.Printf("Failed to create concert: %s", err)
			return err
		}

		for _, seedSession := range seedConcert.Sessions {
			session, err := sessionRepo.CreateConcertSession(ctx, concert.ID, seedSession.SessionDateTime)
			log.Printf("Created session: %d", session.SessionDateTime)
			if err != nil {
				log.Printf("Failed to create session: %s", err)
				return err
			}

			for _, seedSection := range seedSession.Sections {
				// Ensure that the provided category value is one of the allowed enum values
				allowedCategories := map[string]section.Category{
					"CAT1": section.CategoryCAT1,
					"CAT2": section.CategoryCAT2,
					"CAT3": section.CategoryCAT3,
					"CAT4": section.CategoryCAT4,
					"CAT5": section.CategoryCAT5,
				}

				category, ok := allowedCategories[seedSection.Category]
				if !ok {
					log.Printf("Failed to parse category: %t", ok)
					return err
				}
				section, err := sectionRepo.CreateSection(ctx, session.ID, seedSection.Name, seedSection.Capacity, 0, 0, category, seedSection.Price)
				log.Printf("Created section: %s - %s", section.Name, section.Category)

				if err != nil {
					log.Printf("Failed to create section: %s", err)
					return err
				}

				for i := 0; i < section.Capacity; i++ {
					_, err := ticketRepo.CreateTicket(ctx, section.ID, i+1, ticket.StatusAVAILABLE)
					if err != nil {
						log.Printf("Failed to create Ticket for Section: %s", err)
						return err
					}
				}
				log.Printf("Created %d tickets", section.Capacity)
			}
		}
	}
	return nil
}
