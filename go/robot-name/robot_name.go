package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var names = make(map[string]bool)
var maxTries = 100

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if name, err := generateName(maxTries); err != nil {
			return "", err
		} else {
			r.name = name
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

type PreparedNames struct {
	preparedNames []string
	pos           int
	ready         bool
}

func (n *PreparedNames) prepare() {
	n.preparedNames = make([]string, 767000-len(names))
	curr := 0
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 1000; k++ {
				s := fmt.Sprintf("%c%c%03d", i+'A', j+'A', k)
				if !names[s] {
					n.preparedNames[curr] = s
					curr++
				}
			}
		}
	}
	n.ready = true
}

func (n *PreparedNames) next() string {
	if n.pos < len(n.preparedNames) {
		name := n.preparedNames[n.pos]
		n.pos++
		return name
	}
	return ""
}

var pNames = &PreparedNames{}

func generateName(max int) (string, error) {
	if !pNames.ready {
		chars := make([]rune, 5)
		rand.Seed(int64(time.Now().Nanosecond()))
		for i := 0; i < max; i++ {
			chars[0] = rune('A' + rand.Intn(26))
			chars[1] = rune('A' + rand.Intn(26))
			chars[2] = rune('0' + rand.Intn(10))
			chars[3] = rune('0' + rand.Intn(10))
			chars[4] = rune('0' + rand.Intn(10))
			name := string(chars)
			if !names[name] {
				names[name] = true
				return name, nil
			}
		}
		pNames.prepare()
	}
	if name := pNames.next(); name != "" {
		return name, nil
	}
	return "", fmt.Errorf("names exhausted")
}
