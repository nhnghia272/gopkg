package gopkg

func Pointer[E any](v E) *E {
	return &v
}

func PointerSlice[E any](vs []E) []*E {
	ps := make([]*E, len(vs))
	for i, v := range vs {
		ps[i] = Pointer(v)
	}
	return ps
}

func PointerMap[E any](vs map[string]E) map[string]*E {
	ps := make(map[string]*E, len(vs))
	for k, v := range vs {
		ps[k] = Pointer(v)
	}
	return ps
}

func Value[E any](p *E) E {
	if p == nil {
		return *new(E)
	}
	return *p
}

func ValueSlice[E any](ps []*E) []E {
	vs := make([]E, len(ps))
	for i, p := range ps {
		vs[i] = Value(p)
	}
	return vs
}

func ValueMap[E any](ps map[string]*E) map[string]E {
	vs := make(map[string]E, len(ps))
	for k, p := range ps {
		vs[k] = Value(p)
	}
	return vs
}
