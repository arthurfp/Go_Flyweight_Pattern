package flyweight

// Flyweight defines the interface for flyweight objects.
type Flyweight interface {
	Operation(extrinsicState string) string
}
