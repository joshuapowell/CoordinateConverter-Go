package coordinate_converter

type CoordinateDMS struct {
	Degree  float32 `validate:"required,float32"`
	Minutes float32 `validate:"required,float32"`
	Seconds float32 `validate:"required,float32"`
}

/**
 * Convert Degrees Minutes Seconds (DMS) to Decimal Degrees (DD)
 *
 * @todo document this function's parameters and return value
 *
 * Math: (degrees) + (minutes/60) + (seconds/3600) = decimal degrees
 *
 * Example: Given the inputs DmsToDd(30, 15, 50) we should expect
 *          a return of 30.26388888888889.
 *
 */
func DmsToDd(degree float32, minutes float32, seconds float32) float64 {

	var coordinate_dms = CoordinateDMS{
		Degree:  degree,
		Minutes: minutes,
		Seconds: seconds}

	var coordinate_dd float64 = (float64(coordinate_dms.Degree) +
		(float64(coordinate_dms.Minutes) / 60) +
		(float64(coordinate_dms.Seconds) / 3600))

	// @todo make sure that the converted decimal is between the max and min
	//       latitude and longitude

	return coordinate_dd
}
