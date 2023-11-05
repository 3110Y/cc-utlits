package utlits

func Map[F, T any](s []F, f func(F) T) []T {
	r := make([]T, len(s), len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Pointer[T any](str T) *T {
	return &str
}

func NullSafeFunction[T, R any](a *T, f func(*T) *R) *R {
	if a == nil {
		return nil
	}
	return f(a)
}

func ValueOrDefault[T any](a *T) T {
	var empty T
	if a == nil {
		return empty
	}
	return *a
}
