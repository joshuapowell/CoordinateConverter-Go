package coordinate_converter

/**
 * CoordinatePartDD represents a part, either latitude or longitude, of a
 * coordinate pair in decimal degrees or DD notation.
 */
type CoordinatePartDD struct {
	DecimalDegree float64 `validate:"required,float64"`
}
