package internal

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadArg() (Colony, string) {
	arg := os.Args[1:]

	if len(arg) != 1 {
		log.Fatalln("argument size is not correct")
	}

	dat, err := os.ReadFile(arg[0])
	if err != nil {
		log.Fatalln("error to read file")
	}

	input := strings.Split(string(dat), "\n")
	i := 0
	var colony Colony
	colony.Ants = make(map[string][]string)

	for i = 0; i < len(input); i++ {
		s := input[i]
		if i == 0 {
			colony.Cnt, err = strconv.Atoi(s)
			if err != nil || colony.Cnt <= 0 {
				log.Fatalln("Amount of ants is not correct")
			}
		} else if s == "##start" {
			if colony.End != "" {
				log.Fatalln("##start dublicate")
			}
			if i+1 < len(input) {
				colony.Start = getRoomName(input[i+1])
			} else {
				log.Fatalln("Starting room is not given")
			}
		} else if s == "##end" {
			if colony.End != "" {
				log.Fatalln("##end dublicate")
			}

			if i+1 < len(input) {
				colony.End = getRoomName(input[i+1])
			} else {
				log.Fatalln("Ending room is not given")
			}
		} else if s[0] == '#' || s[0] == 'L' {
			continue
		} else {
			row := strings.Split(s, " ")
			if len(row) != 3 {
				break
			}
			room := getRoomName(s)
			roomCreated := newRoom(&colony, room)
			if !roomCreated { // in case if it is not room, but the path
				break
			}
		}
	}
	if colony.Start == "" || colony.End == "" {
		log.Fatalln("##start or ##end is missing")
	}
	for ; i < len(input); i++ {
		s := input[i]
		if len(s) == 0 {
			log.Fatalln("not correct input")
		} else if s[0] == '#' || s[0] == 'L' {
			continue
		}
		row := strings.Split(s, "-")
		if len(row) != 2 {
			log.Fatalln("input is not correct", input[i])
		}
		_, isRoomExists := colony.Ants[row[0]]
		if !isRoomExists {
			log.Fatalln("Room does not exist:", row[0])
		}
		_, isRoomExists = colony.Ants[row[1]]
		if !isRoomExists {
			log.Fatalln("Room does not exist:", row[1])
		}

		for _, neighbour := range colony.Ants[row[0]] {
			if neighbour == row[1] {
				log.Fatalln("dublicate:", row)
			}
		}
		colony.Ants[row[0]] = append(colony.Ants[row[0]], row[1])
		colony.Ants[row[1]] = append(colony.Ants[row[1]], row[0])

	}
	return colony, string(dat)
}

func getRoomName(s string) string {
	input := strings.Split(s, " ")
	if len(input) != 3 {
		return ""
	}
	_, err := strconv.Atoi(input[1])
	_, err1 := strconv.Atoi(input[2])
	if err != nil || err1 != nil {
		log.Fatalln("room's coordinates are not correct")
	}
	return input[0]
}

func newRoom(colony *Colony, room string) bool {
	if room == "" {
		return false
	}
	_, isRoomExists := colony.Ants[room]
	if isRoomExists {
		log.Fatalln("Dublicate room:" + room)
	}
	colony.Ants[room] = []string{}
	colony.RoomCnt++
	return true
}
