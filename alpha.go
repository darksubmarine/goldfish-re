package goldfish_re

// _alpha node for rete network
type _alpha struct {
	active []*_condition
	rel    map[cuid]_beta
}
