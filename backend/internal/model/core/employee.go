package core

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// IntSlice is a custom type to handle PostgreSQL integer[] arrays
type IntSlice []int

// Scan implements the sql.Scanner interface for IntSlice
func (is *IntSlice) Scan(value interface{}) error {
	if value == nil {
		*is = nil
		return nil
	}

	// Handle the case where the value is a []uint8 (byte slice) from PostgreSQL
	if b, ok := value.([]uint8); ok {
		str := string(b)
		if str == "{}" {
			*is = []int{}
			return nil
		}
		// Remove curly braces and split by comma
		str = str[1 : len(str)-1]
		parts := strings.Split(str, ",")
		temp := make([]int, len(parts))
		for i, part := range parts {
			var n int
			_, err := fmt.Sscanf(part, "%d", &n)
			if err != nil {
				return fmt.Errorf("failed to parse int from %s: %w", part, err)
			}
			temp[i] = n
		}
		*is = temp
		return nil
	}
	return fmt.Errorf("unsupported scan type for IntSlice: %T", value)
}

// Value implements the driver.Valuer interface for IntSlice
func (is IntSlice) Value() (driver.Value, error) {
	if is == nil {
		return nil, nil
	}
	// Convert to PostgreSQL array format (e.g., "{1,2,-1}")
	str := fmt.Sprintf("{%s}", strings.Join(strings.Split(fmt.Sprintf("%v", is), " ")[1:len(is)*2:2], ","))
	return str, nil
}

type Employee struct {
	ID                   int       `db:"id"`
	CompanyID            int       `db:"company_id"`
	Name                 string    `db:"name"`
	Position             string    `db:"position"`
	Mail                 string    `db:"mail"`
	Password             string    `db:"password"`
	Salt                 string    `db:"salt"`
	Image                string    `db:"image"`
	Rating               IntSlice  `db:"rating"` // Changed from []int to IntSlice
	ActiveProjectsCount  int       `db:"active_projects_count"`
	OverdueProjectsCount int       `db:"overdue_projects_count"`
	TotalProjectsCount   int       `db:"total_projects_count"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}
