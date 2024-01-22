package config

import (
	"time"

	"gorm.io/gen"
	"gorm.io/gorm"
)

type CommonMethod struct {
	gorm.Model
	ID int64
}

// IsEmpty determines whether the structure is empty
func (m *CommonMethod) IsEmpty() bool {
	if m == nil {
		return true
	}
	return m.ID == 0
}

// GetID get the ID of the database table
func (m *CommonMethod) GetID() int64 {
	return int64(m.ID)
}

// Querier
type Querier interface {
	// returns struct and error
	//
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int64) (gen.T, error)

	// returns map and error
	//
	// SELECT * FROM @@table WHERE id=@id
	GetByIDToMap(id int64) (gen.M, error)

	// Query the corresponding value by column name
	//
	// SELECT * FROM @@table WHERE @@column=@value
	FilterWithColumn(column string, value string) (gen.T, error)

	// Query data based on created_at
	//
	// SELECT * FROM @@table
	//  {{where}}
	//    {{if !start.IsZero()}}
	//      created_at > @start
	//    {{end}}
	//    {{if !end.IsZero()}}
	//      AND created_at < @end
	//    {{end}}
	//  {{end}}
	FilterCreatedWithTime(start, end time.Time) ([]gen.T, error)
}
