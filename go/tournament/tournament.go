package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	Name           string
	PL, W, D, L, P int
}

func NewTeam(name string) *Team {
	return &Team{Name: name}
}

func (t *Team) Win() {
	t.PL++
	t.W++
	t.P += 3
}

func (t *Team) Lose() {
	t.PL++
	t.L++
}

func (t *Team) Draw() {
	t.PL++
	t.D++
	t.P += 1
}

func (t *Team) String() string {
	return fmt.Sprintf("Name: %s, PL: %d, W: %d, D: %d, L: %d, P: %d", t.Name, t.PL, t.W, t.D, t.L, t.P)
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams := make(map[string]*Team)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()

		// ignore empty lines and comments
		if len(text) == 0 || text[0] == '#' {
			continue
		}

		// split tokens
		tokens := strings.Split(text, ";")
		if len(tokens) != 3 {
			return errors.New("malformed input line: " + text)
		}

		// add new teams to map
		for _, t := range tokens[:2] {
			if _, ok := teams[t]; !ok {
				team := NewTeam(t)
				teams[t] = team
			}
		}

		// update scores
		switch t1name, t2name, outcome := tokens[0], tokens[1], tokens[2]; outcome {
		case "win":
			teams[t1name].Win()
			teams[t2name].Lose()
		case "loss":
			teams[t1name].Lose()
			teams[t2name].Win()
		case "draw":
			teams[t1name].Draw()
			teams[t2name].Draw()
		default:
			return errors.New("unknown outcome: " + outcome)
		}
	}

	// get sorted slice of teams
	teamsSlice := sortTeams(teams)

	// print header
	format := "%-31s|%3s |%3s |%3s |%3s |%3s\n"
	fmt.Fprintf(writer, format, "Team", "MP", "W", "D", "L", "P")

	// print tally
	format = "%-31s|%3d |%3d |%3d |%3d |%3d\n"
	for _, t := range teamsSlice {
		if _, err := fmt.Fprintf(writer, format, t.Name, t.PL, t.W, t.D, t.L, t.P); err != nil {
			return err
		}
	}
	return nil
}

func sortTeams(teams map[string]*Team) []*Team {
	tSlice := []*Team{}
	for _, v := range teams {
		tSlice = append(tSlice, v)
	}

	// sort first by decending points, second lexicographically
	sort.Slice(tSlice, func(i, j int) bool {
		if tSlice[i].P != tSlice[j].P {
			return tSlice[i].P > tSlice[j].P
		}
		return tSlice[i].Name < tSlice[j].Name
	})
	return tSlice
}
