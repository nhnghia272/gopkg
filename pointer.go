package gopkg

func Pointer[E comparable](v E) *E {
	return &v
}

func PointerSlice[E comparable](vs []E) []*E {
	ps := make([]*E, len(vs))
	for i, v := range vs {
		vv := v
		ps[i] = &vv
	}
	return ps
}

func PointerMap[E comparable](vs map[string]E) map[string]*E {
	ps := make(map[string]*E, len(vs))
	for k, v := range vs {
		vv := v
		ps[k] = &vv
	}
	return ps
}

func Value[E comparable](p *E) E {
	if p == nil {
		return *new(E)
	}
	return *p
}

func ValueSlice[E comparable](ps []*E) []E {
	vs := make([]E, len(ps))
	for i, p := range ps {
		pp := p
		if pp == nil {
			vs[i] = *new(E)
		} else {
			vs[i] = *pp
		}
	}
	return vs
}

func ValueMap[E comparable](ps map[string]*E) map[string]E {
	vs := make(map[string]E, len(ps))
	for k, p := range ps {
		pp := p
		if pp == nil {
			vs[k] = *new(E)
		} else {
			vs[k] = *pp
		}
	}
	return vs
}
