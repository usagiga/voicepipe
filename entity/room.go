package entity

// Room represents ...
// To clarify difference between chan(Go) and channel(Discord),
// it is named as room.
type Room struct {
	ID string
}

// NewRoom is constructor of Room
func NewRoom(id string) *Room {
	return &Room{
		ID: id,
	}
}
