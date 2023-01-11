package response

func convertArray[T any, S any, F func(a T) S](v []T, f F) []S {
	s := []S{}
	for _, e := range v {
		s = append(s, f(e))
	}

	return s
}
