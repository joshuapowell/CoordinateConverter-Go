package coordinate_converter

/**
 * CoordinateType defines a value that is either latitude or longitude.
 *
 * Example:
 *
 *     ```
 *
 *     ```
 */
type CoordinateType int64

const (
	Latitude Direction = iota
	Longitude
)
