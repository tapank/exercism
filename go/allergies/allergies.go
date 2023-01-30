package allergies

var allergies = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
var scores = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Allergies returns names of allergies based on score.
func Allergies(score uint) []string {
	alist := make([]string, 0, 8)
	for i := 0; i < len(allergies); i++ {
		if score&(1<<i) != 0 {
			alist = append(alist, allergies[i])
		}
	}
	return alist
}

// AllergicTo returns true if the allergen is part of score, else false.
func AllergicTo(score uint, allergen string) bool {
	return scores[allergen]&score != 0
}
