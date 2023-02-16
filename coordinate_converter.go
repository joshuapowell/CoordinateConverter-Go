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
 * Example: Given the inputs DmsToDd(30, 15, 50) we should expect
 *          a return of 30.26388888888889.
 *
 */
func DmsToDd(degree string, minute string, second string) float64 {

	// @todo check for valid inputs

	// @todo assign valid inputs to the CoordinateDMS struct

	// @todo convert the struct to a single decimal (float64)

	// @todo make sure that the converted decimal is between the max and min
	//       latitude and longitude

	// @todo return the newly converted decimal coordinate
	return 30.2638888888
}
