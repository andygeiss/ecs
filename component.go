package ecs

// Component contains only the data (no behaviour at all).
//
// Each component type owns one bit of the mask — 1<<0, 1<<1, and so on —
// because FilterByMask matches with a bitwise AND. Number the types 1, 2, 3
// instead and FilterByMask(3) matches an entity carrying only types 1 and 2,
// which does not carry type 3 at all. Past 64 types, switch to
// ComponentWithName.
type Component interface {
	Mask() uint64
}

// ComponentWithName is a Component that also carries a name. FilterByNames
// matches on that name, so a game with more than 64 component types is not
// held to the 64 bits of the mask.
type ComponentWithName interface {
	Component
	Name() string
}
