package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if equal(l1, l2) {
		return RelationEqual
	}
	if sub(l1, l2) {
		return RelationSuperlist
	}
	if sub(l2, l1) {
		return RelationSublist
	}
	return RelationUnequal
}

func sub(list, slist []int) bool {
	if len(slist) == 0 {
		return true
	}
	for i, v := range list {
		if v == slist[0] && len(slist) <= len(list)-i {
			if equal(list[i:i+len(slist)], slist) {
				return true
			}
		}
	}
	return false
}

func equal(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}
