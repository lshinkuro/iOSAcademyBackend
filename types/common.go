package types

import (
	"database/sql/driver"
	"encoding/json"
)

type StringArray []string

// Value makes StringArray implement the driver.Valuer interface
func (a StringArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}

// Scan makes StringArray implement the sql.Scanner interface
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, &a)
	case string:
		return json.Unmarshal([]byte(v), &a)
	default:
		*a = nil
		return nil
	}
}

// LearningPoint is an alias for StringArray to maintain semantic meaning
type LearningPoint = StringArray
