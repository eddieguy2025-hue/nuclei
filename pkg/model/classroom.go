package model

// Classroom represents a classroom entity with basic attributes.
type Classroom struct {
	// Name is the name of the classroom.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Capacity is the maximum number of students the classroom can hold.
	Capacity int `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	// Location is the physical location or room number of the classroom.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}

// NewClassroom creates a new Classroom with the given name, capacity, and location.
func NewClassroom(name string, capacity int, location string) *Classroom {
	return &Classroom{
		Name:     name,
		Capacity: capacity,
		Location: location,
	}
}
