package todo

type Status int

const (
	Pending = iota
	Succeeded
	Cancelled
)

// String define the string representation of the Status
func (s Status) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Succeeded:
		return "Succeeded"
	case Cancelled:
		return "Cancelled"
	}

	return "Invalid status"
}
