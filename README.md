# Lem-in

Lem-in is a program designed to simulate the movement of ants through an ant farm, finding the quickest path for them to traverse from a start room to an end room.

## How it Works

1. **Setup**: You define an ant farm with rooms and tunnels.
2. **Objective**: All ants start in the room labeled ##start. The goal is to bring them to the room labeled ##end with the fewest moves possible.
3. **Pathfinding**: Lem-in calculates the shortest path or paths for the ants to reach the end room while avoiding traffic jams.
4. **Output**: The program displays the number of ants, the layout of rooms, and the movements of ants through the tunnels.

## Input Format

- Rooms: Defined by "name coord_x coord_y", e.g., "Room1 2 3".
- Links: Defined by "name1-name2", e.g., "Room1-Room2".
- Start and End Rooms: Labeled as ##start and ##end respectively.
- Comments: Lines starting with # are comments and are ignored.
- Commands: Any unknown command is ignored.

## Output Format

The program displays the movements of ants in the following format:

```
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

Where:
- x, z, r represent the ant numbers (1 to number_of_ants).
- y, w, o represent the rooms names.

## Example

```
##start
1 1 1
2 2 2
##end
3 3 3
1-2
2-3
```

## Usage

Run the program with the input file as an argument:

```bash
./lem-in input_file.txt
``
