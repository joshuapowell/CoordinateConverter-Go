package coordinate_converter

/**
 * CoordinatePartDMS represents a part, either latitude or longitude, of a
 * coordinate pair in degrees-minutes-seconds or DMS notation.
 */
type CoordinatePartDMS struct {
	Degree  float32 `validate:"required,float32"`
	Minutes float32 `validate:"required,float32"`
	Seconds float32 `validate:"required,float32"`
}
