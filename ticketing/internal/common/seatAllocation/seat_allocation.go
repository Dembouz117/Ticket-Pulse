package seatallocation

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type Seat struct {
	SeatID          uuid.UUID `json:"ticketId"`
	SectionID       uuid.UUID `json:"sectionId"`
	SectionCategory string    `json:"sectionCategory"`
	SectionName     string    `json:"sectionName"`
	SeatNumber      int       `json:"seatNumber"`
}

type SeatAvailable struct {
	SectionID uuid.UUID `json:"sectionId"`
	Seats     []Seat    `json:"seats"`
}

type SeatRequirement struct {
	SectionID uuid.UUID `json:"sectionId"`
	Quantity  int       `json:"quantity"`
}

func seatAllocationAlgorithm(section SeatRequirement, seatAvailable SeatAvailable) []Seat {
	var result []Seat

	seats := seatAvailable.Seats
	startSeatIndex := 0
	endSeatIndex := 0
	seatNumber := seats[0].SeatNumber
	consecutiveSeats := 1
	var lock sync.Mutex

	lock.Lock()
	defer lock.Unlock()

	// if we only need 1
	if consecutiveSeats == section.Quantity {
		result = append(result, seats[startSeatIndex:endSeatIndex+1]...)
		return result
	}

	for i := 1; i < len(seats); i++ {
		if seatNumber+1 == seats[i].SeatNumber {
			// Found an available seat
			consecutiveSeats++
			if consecutiveSeats == section.Quantity {
				// Found enough consecutive seats
				startSeatIndex = i - section.Quantity + 1
				endSeatIndex = i
				break
			}
			seatNumber += 1
		} else {
			// Reset consecutive seats count
			consecutiveSeats = 0
		}
	}

	if consecutiveSeats == section.Quantity {
		result = append(result, seats[startSeatIndex:endSeatIndex+1]...)
	}

	return result
}

// takes in requirements of seats and available seats
func AllocateSeatsConcurrently(seatsRequirementList []SeatRequirement, seatAvailableList []SeatAvailable) []Seat {
	var result []Seat

	// Loop seat requirement list
	for _, section := range seatsRequirementList {

		// loop the available seats
		for _, seatAvailable := range seatAvailableList {

			// match the section id with what we have
			if seatAvailable.SectionID == section.SectionID {

				// check if the quantity matches first
				if len(seatAvailable.Seats) < section.Quantity {
					log.Println("Lack of seats")
					return nil
				}

				// sliding window
				seats := seatAllocationAlgorithm(section, seatAvailable)
				if seats == nil {
					log.Println("Seat allocation algo cannot find")
					return nil
				}
				result = append(result, seats...)
			}
		}
	}

	return result
}
