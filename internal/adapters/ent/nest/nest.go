// Code generated by ent, DO NOT EDIT.

package nest

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the nest type in the database.
	Label = "nest"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// EdgeParts holds the string denoting the parts edge name in mutations.
	EdgeParts = "parts"
	// EdgeRemnant holds the string denoting the remnant edge name in mutations.
	EdgeRemnant = "remnant"
	// EdgeRemnantUsed holds the string denoting the remnant_used edge name in mutations.
	EdgeRemnantUsed = "remnant_used"
	// Table holds the table name of the nest in the database.
	Table = "nests"
	// PartsTable is the table that holds the parts relation/edge. The primary key declared below.
	PartsTable = "nest_parts"
	// PartsInverseTable is the table name for the Part entity.
	// It exists in this package in order to avoid circular dependency with the "part" package.
	PartsInverseTable = "parts"
	// RemnantTable is the table that holds the remnant relation/edge.
	RemnantTable = "remnants"
	// RemnantInverseTable is the table name for the Remnant entity.
	// It exists in this package in order to avoid circular dependency with the "remnant" package.
	RemnantInverseTable = "remnants"
	// RemnantColumn is the table column denoting the remnant relation/edge.
	RemnantColumn = "nest_remnant"
	// RemnantUsedTable is the table that holds the remnant_used relation/edge.
	RemnantUsedTable = "nests"
	// RemnantUsedInverseTable is the table name for the Remnant entity.
	// It exists in this package in order to avoid circular dependency with the "remnant" package.
	RemnantUsedInverseTable = "remnants"
	// RemnantUsedColumn is the table column denoting the remnant_used relation/edge.
	RemnantUsedColumn = "remnant_remnant_used"
)

// Columns holds all SQL columns for nest fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldLength,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "nests"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"remnant_remnant_used",
}

var (
	// PartsPrimaryKey and PartsColumn2 are the table columns denoting the
	// primary key for the parts relation (M2M).
	PartsPrimaryKey = []string{"nest_id", "part_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LengthValidator is a validator for the "length" field. It is called by the builders before save.
	LengthValidator func(float64) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
