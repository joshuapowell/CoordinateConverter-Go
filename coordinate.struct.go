package coordinate_converter

/**
 * A Coordinate
 */
type Coordinate struct {
	Latitude  string `validate:"required,string,min=-90,max=90"`
	Longitude string `validate:"required,string,min=-180,max=180"`
}
