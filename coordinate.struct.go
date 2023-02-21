package coordinate_converter

/**
 * Coordinate defines a data structure that represents a coordinate pair.
 *
 * The coordinate pair consists of a `Latitude` value between -90.0 and 90.0
 * and a longitude value between -180.0 and 180.0
 *
 */
type Coordinate struct {
	Latitude  float64 `validate:"required,string,min=-90,max=90"`
	Longitude float64 `validate:"required,string,min=-180,max=180"`
}
