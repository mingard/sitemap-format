package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

// BoolStr is a string representation of a boolean.
type BoolStr string

const (
	// Yes is a truthy string representation.
	Yes BoolStr = "yes"
	// No is a falsy string representation.
	No BoolStr = "no"
)
