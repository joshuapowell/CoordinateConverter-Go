package coordinate_converter

/**
 * Direction defines a constant that represents a  limited set of values for
 * identifying the direction in which to represent a coordinate.
 *
 * Direction is added to a Coordinate to represent whether it is in the North
 * East, South, or West direction on the coordinate map thus giving the value
 * either a positive (North, East) or negative (South, West) value.
 *
 * Example:
 *
 *     ```
 *
 *     ```
 */
type Direction int64

const (
	North Direction = iota
	East
	South
	West
)
