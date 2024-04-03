package internal

type ant map[string][]string

type Colony struct {
	Cnt     int
	Ants    ant
	Start   string
	End     string
	RoomCnt int
}

type Ant struct {
	ID        int
	Room      string
	PrevRooms int
}
