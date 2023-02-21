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
 * Math:
 *    Step 1: 30.26388888888889 » 30 degrees
 *    Step 2: 0.26388888888889*60 = 15.8333333333 » 15 minutes
 *    Step 3: 0.8333333333*60 = 49.999999998 » 50 Seconds
 *
 * Example: Given the inputs ConvertDdToDms(30.26388888888889) we should
 *          expect return of 30.26388888888889.
 */
func ConvertDdToDms(DecimalDegree float64, LatLng CoordinateType) string {

	// @todo Degrees = Separate the whole number from the numbers after the
	//       decimal

	// @todo Minutes = Separate the whole number from the numbers after the
	//       decminal

	// @todo Seconds = Multiple the remaining number by 60'

	// @todo Concatenate Degrees, Minutes, and Seconds

	// @todo if LatLng is Latitude and less a negative number it's West

	// @todo if LatLng is Longitude and less a negative number it's South

	return ""
}
