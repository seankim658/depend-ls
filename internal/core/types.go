package core

// Represents the specific location in code where an entity is used or defined.
type Reference struct {
	Line   uint32
	Column uint32
}

// Represents a code entity (function, class, etc.) and its dependencies.
type Dependency struct {
	// Name of the entity (e.g., function name)
	Name string
	// Type of the entity (e.g., "function", "class")
	Type string
	// Line where the entity is defined
	Line uint32
	// Function calls made by this entity
	Calls map[string][]Reference
	// Types used by this entity
	UsesTypes map[string][]Reference
	// Constants referenced by this entity
	Constants map[string][]Reference
}
