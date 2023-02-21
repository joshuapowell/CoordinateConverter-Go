package coordinate_converter

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
func ConvertDmsToDd(degrees float32, minutes float32, seconds float32, direction Direction) float64 {

	var coordinate_dms = CoordinatePartDMS{
		Degrees: degrees,
		Minutes: minutes,
		Seconds: seconds}

	var _calculated_dd float64 = (float64(coordinate_dms.Degrees) +
		(float64(coordinate_dms.Minutes) / 60) +
		(float64(coordinate_dms.Seconds) / 3600))

	var coordinate_dd = CoordinatePartDD{
		DecimalDegree: _calculated_dd}

	// Negate _calculated_dd value when direction is 2 (South) or 3 (West)
	if direction == 2 || direction == 3 {
		coordinate_dd.DecimalDegree = -coordinate_dd.DecimalDegree
	}

	return coordinate_dd.DecimalDegree
}

/**
 * Convert Decimal Degrees (DD) to Degrees-Minutes-Seconds (DMS).
 *
 * Math: __unknown__
 *
 */
func ConvertDdToDms() string {

	return ""
}
