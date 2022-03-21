package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name          string
	matchesPlayed int
	won           int
	drawn         int
	lost          int
	point         int
}

type TeamList []Team

func (p TeamList) Len() int      { return len(p) }
func (p TeamList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p TeamList) Less(i, j int) bool {

	if p[i].point == p[j].point {
		return p[i].name < p[j].name
	}

	return p[i].point > p[j].point
}

func Tally(reader io.Reader, writer io.Writer) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	s := buf.String()

	lines := strings.Split(s, "\n")
	games := map[string]Team{}
	for _, v := range lines {
		if v == "" || strings.HasPrefix(v, "#") {
			continue
		}

		line := strings.Split(v, ";")
		if len(line) != 3 {
			return errors.New("failed")
		}

		firstTeamName := line[0]
		secondTeamName := line[1]
		result := line[2]
		firstTeam, firstTeamExist := games[firstTeamName]
		if !firstTeamExist {
			firstTeam = Team{
				name:          firstTeamName,
				matchesPlayed: 0,
				won:           0,
				drawn:         0,
				lost:          0,
				point:         0,
			}
		}
		secondTeam, secondTeamExist := games[secondTeamName]
		if !secondTeamExist {
			secondTeam = Team{
				name:          secondTeamName,
				matchesPlayed: 0,
				won:           0,
				drawn:         0,
				lost:          0,
				point:         0,
			}
		}

		if result == "draw" {
			firstTeam.matchesPlayed += 1
			secondTeam.matchesPlayed += 1

			firstTeam.drawn += 1
			firstTeam.point += 1
			secondTeam.drawn += 1
			secondTeam.point += 1

		} else if result == "win" {
			firstTeam.matchesPlayed += 1
			secondTeam.matchesPlayed += 1

			firstTeam.won += 1
			firstTeam.point += 3
			secondTeam.lost += 1

		} else if result == "loss" {
			firstTeam.matchesPlayed += 1
			secondTeam.matchesPlayed += 1

			firstTeam.lost += 1
			secondTeam.won += 1
			secondTeam.point += 3
		} else {
			return errors.New("failed")
		}
		games[firstTeamName] = firstTeam
		games[secondTeamName] = secondTeam

	}

	p := make(TeamList, len(games))
	i := 0
	for _, v := range games {
		p[i] = v
		i++
	}

	sort.Sort(p)

	fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, v := range p {
		fmt.Fprintf(writer, "%-31s|%3d |%3d |%3d |%3d |%3d\n", v.name, v.matchesPlayed, v.won, v.drawn, v.lost, v.point)
	}

	return nil
}
