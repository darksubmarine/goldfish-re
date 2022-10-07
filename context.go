package goldfish_re

type _factContext map[string]iFact

func (f _factContext) get(token string) (iFact, bool) {
	v, ok := f[token]
	return v, ok
}

func (f _factContext) set(fact iFact) {
	f[fact.token()] = fact
}
