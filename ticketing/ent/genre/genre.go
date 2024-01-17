// Code generated by ent, DO NOT EDIT.

package genre

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the genre type in the database.
	Label = "genre"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeConcerts holds the string denoting the concerts edge name in mutations.
	EdgeConcerts = "concerts"
	// Table holds the table name of the genre in the database.
	Table = "genres"
	// ConcertsTable is the table that holds the concerts relation/edge. The primary key declared below.
	ConcertsTable = "concert_genres"
	// ConcertsInverseTable is the table name for the Concert entity.
	// It exists in this package in order to avoid circular dependency with the "concert" package.
	ConcertsInverseTable = "concerts"
)

// Columns holds all SQL columns for genre fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// ConcertsPrimaryKey and ConcertsColumn2 are the table columns denoting the
	// primary key for the concerts relation (M2M).
	ConcertsPrimaryKey = []string{"concert_id", "genre_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Genre queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByConcertsCount orders the results by concerts count.
func ByConcertsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newConcertsStep(), opts...)
	}
}

// ByConcerts orders the results by concerts terms.
func ByConcerts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newConcertsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newConcertsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ConcertsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ConcertsTable, ConcertsPrimaryKey...),
	)
}
