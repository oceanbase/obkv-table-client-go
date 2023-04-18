package route

import "strconv"

func CreateInStatement(values []int) string {
	// Create inStatement "(0,1,2...partNum);".
	var inStatement string
	inStatement += "("
	for i, v := range values {
		if i > 0 {
			inStatement += ", "
		}
		inStatement += strconv.Itoa(v)
	}
	inStatement += ");"
	return inStatement
}
