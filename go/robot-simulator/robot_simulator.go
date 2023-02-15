package robot

import (
	"sync"
	"time"
)

const (
	N Dir = iota
	E
	S
	W
)

var dirs = map[Dir]string{
	N: "N",
	E: "E",
	S: "S",
	W: "W",
}

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func (d Dir) String() string {
	return dirs[d]
}

// Step 2
// Define Action type here.
type Action Command

func StartRobot(command chan Command, action chan Action) {
	for c := range command {
		action <- Action(c)
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for a := range action {
		switch a {
		case 'A':
			switch robot.Dir {
			case N:
				if extent.Max.Northing > robot.Pos.Northing {
					robot.Pos.Northing++
				}
			case E:
				if extent.Max.Easting > robot.Pos.Easting {
					robot.Pos.Easting++
				}
			case S:
				if extent.Min.Northing < robot.Pos.Northing {
					robot.Pos.Northing--
				}
			case W:
				if extent.Min.Easting < robot.Pos.Easting {
					robot.Pos.Easting--
				}
			}
		case 'R':
			robot.Dir = (robot.Dir + 1) % 4
		case 'L':
			robot.Dir = (robot.Dir + 3) % 4
		}
	}
	report <- robot
	close(report)
}

// Step 3
// Define Action3 type here.
type Action3 struct {
	name   string
	action Action
}

var wg *sync.WaitGroup

func StartRobot3(name, script string, action chan Action3, log chan string) {
	if wg == nil {
		wg = &sync.WaitGroup{}
		go func() {
			wg.Wait()
			close(action)
		}()
	}
	wg.Add(1)
	defer wg.Done()

	action <- Action3{name: name, action: Action('S')}
	for _, a := range script {
		action <- Action3{name: name, action: Action(a)}
	}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	time.Sleep(time.Millisecond * 150)
	defer func() {
		wg = nil
	}()
	robomap := map[string]int{}
	posmap := map[Pos]bool{}
	for i, r := range robots {
		if _, ok := robomap[r.Name]; ok {
			log <- "duplicate name"
			rep <- robots
			return
		} else {
			if _, ok := posmap[r.Pos]; ok {
				log <- "duplicate position"
				rep <- robots
				return
			} else {
				posmap[r.Pos] = true
			}
			if r.Pos.Easting < extent.Min.Easting ||
				r.Pos.Easting > extent.Max.Easting ||
				r.Pos.Northing < extent.Min.Northing ||
				r.Pos.Northing > extent.Max.Northing {

				log <- "ouside room"
				rep <- robots
				return
			}
			robomap[r.Name] = i
		}
	}
	for a := range action {
		if a.name == "" {
			log <- "no name"
			rep <- robots
			return
		}
		switch a.action {
		case 'S':
			if _, ok := robomap[a.name]; !ok {
				log <- "bad robot name"
				rep <- robots
				return
			}
		case 'A':
			r := robomap[a.name]
			switch robots[r].Dir {
			case N:
				if extent.Max.Northing > robots[r].Pos.Northing {
					newPos := copyPos(robots[r].Pos)
					newPos.Northing++
					if bump(robots, newPos) {
						log <- "bump prevented"
					} else {
						robots[r].Pos.Northing++
					}
				} else {
					log <- "invalid move"
				}
			case E:
				if extent.Max.Easting > robots[r].Pos.Easting {
					newPos := copyPos(robots[r].Pos)
					newPos.Easting++
					if bump(robots, newPos) {
						log <- "bump prevented"
					} else {
						robots[r].Pos.Easting++
					}
				} else {
					log <- "invalid move"
				}
			case S:
				if extent.Min.Northing < robots[r].Pos.Northing {
					newPos := copyPos(robots[r].Pos)
					newPos.Northing--
					if bump(robots, newPos) {
						log <- "bump prevented"
					} else {
						robots[r].Pos.Northing--
					}
				} else {
					log <- "invalid move"
				}
			case W:
				if extent.Min.Easting < robots[r].Pos.Easting {
					newPos := copyPos(robots[r].Pos)
					newPos.Easting--
					if bump(robots, newPos) {
						log <- "bump prevented"
					} else {
						robots[r].Pos.Easting--
					}
				} else {
					log <- "invalid move"
				}
			}
		case 'R':
			r := robomap[a.name]
			robots[r].Dir = (robots[r].Dir + 1) % 4
		case 'L':
			r := robomap[a.name]
			robots[r].Dir = (robots[r].Dir + 3) % 4
		default:
			log <- "bad command"
			rep <- robots
			return
		}
	}
	rep <- robots
}

func bump(robots []Step3Robot, pos Pos) bool {
	for _, r := range robots {
		if r.Pos == pos {
			return true
		}
	}
	return false
}

func copyPos(pos Pos) Pos {
	return Pos{pos.Easting, pos.Northing}
}
