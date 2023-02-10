package school

import "sort"

type School struct {
	grades map[int][]string
}

type Grade struct {
	n        int
	students []string
}

func New() *School {
	return &School{make(map[int][]string)}
}

func (s *School) Add(student string, grade int) {
	if students, ok := s.grades[grade]; ok {
		s.grades[grade] = append(students, student)
		sort.Strings(s.grades[grade])
	} else {
		s.grades[grade] = []string{student}
	}
}

func (s *School) Grade(level int) []string {
	if students, ok := s.grades[level]; ok {
		return students
	}
	return nil
}

func (s *School) Enrollment() []Grade {
	grades := make([]Grade, 0, len(s.grades))
	for n, students := range s.grades {
		grades = append(grades, Grade{n, students})
	}
	sort.Slice(grades, func(i, j int) bool {
		return grades[i].n < grades[j].n
	})
	return grades
}
