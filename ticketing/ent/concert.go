// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"ticketing/ent/concert"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Concert is the model entity for the Concert schema.
type Concert struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Artist holds the value of the "artist" field.
	Artist string `json:"artist,omitempty"`
	// ImageUrl holds the value of the "imageUrl" field.
	ImageUrl string `json:"imageUrl,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Headline holds the value of the "headline" field.
	Headline string `json:"headline,omitempty"`
	// Featured holds the value of the "featured" field.
	Featured bool `json:"featured,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConcertQuery when eager-loading is set.
	Edges        ConcertEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ConcertEdges holds the relations/edges for other nodes in the graph.
type ConcertEdges struct {
	// HasConcertSessions holds the value of the hasConcertSessions edge.
	HasConcertSessions []*ConcertSession `json:"hasConcertSessions,omitempty"`
	// Genres holds the value of the genres edge.
	Genres []*Genre `json:"genres,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// HasConcertSessionsOrErr returns the HasConcertSessions value or an error if the edge
// was not loaded in eager-loading.
func (e ConcertEdges) HasConcertSessionsOrErr() ([]*ConcertSession, error) {
	if e.loadedTypes[0] {
		return e.HasConcertSessions, nil
	}
	return nil, &NotLoadedError{edge: "hasConcertSessions"}
}

// GenresOrErr returns the Genres value or an error if the edge
// was not loaded in eager-loading.
func (e ConcertEdges) GenresOrErr() ([]*Genre, error) {
	if e.loadedTypes[1] {
		return e.Genres, nil
	}
	return nil, &NotLoadedError{edge: "genres"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Concert) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case concert.FieldFeatured:
			values[i] = new(sql.NullBool)
		case concert.FieldTitle, concert.FieldArtist, concert.FieldImageUrl, concert.FieldDescription, concert.FieldHeadline:
			values[i] = new(sql.NullString)
		case concert.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Concert fields.
func (c *Concert) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case concert.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case concert.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case concert.FieldArtist:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field artist", values[i])
			} else if value.Valid {
				c.Artist = value.String
			}
		case concert.FieldImageUrl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field imageUrl", values[i])
			} else if value.Valid {
				c.ImageUrl = value.String
			}
		case concert.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case concert.FieldHeadline:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field headline", values[i])
			} else if value.Valid {
				c.Headline = value.String
			}
		case concert.FieldFeatured:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field featured", values[i])
			} else if value.Valid {
				c.Featured = value.Bool
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Concert.
// This includes values selected through modifiers, order, etc.
func (c *Concert) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryHasConcertSessions queries the "hasConcertSessions" edge of the Concert entity.
func (c *Concert) QueryHasConcertSessions() *ConcertSessionQuery {
	return NewConcertClient(c.config).QueryHasConcertSessions(c)
}

// QueryGenres queries the "genres" edge of the Concert entity.
func (c *Concert) QueryGenres() *GenreQuery {
	return NewConcertClient(c.config).QueryGenres(c)
}

// Update returns a builder for updating this Concert.
// Note that you need to call Concert.Unwrap() before calling this method if this Concert
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Concert) Update() *ConcertUpdateOne {
	return NewConcertClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Concert entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Concert) Unwrap() *Concert {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Concert is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Concert) String() string {
	var builder strings.Builder
	builder.WriteString("Concert(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("title=")
	builder.WriteString(c.Title)
	builder.WriteString(", ")
	builder.WriteString("artist=")
	builder.WriteString(c.Artist)
	builder.WriteString(", ")
	builder.WriteString("imageUrl=")
	builder.WriteString(c.ImageUrl)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("headline=")
	builder.WriteString(c.Headline)
	builder.WriteString(", ")
	builder.WriteString("featured=")
	builder.WriteString(fmt.Sprintf("%v", c.Featured))
	builder.WriteByte(')')
	return builder.String()
}

// Concerts is a parsable slice of Concert.
type Concerts []*Concert
