package greetings

type CoordinateDMS struct {
	Degree  string `validate:"required,string"`
	Minutes string `validate:"required,string"`
	Seconds string `validate:"required,string"`
}

/**
 * Convert Degrees Minutes Seconds (DMS) to Decimal Degrees (DD)
 *
 *
 *
 */
func DmsToDd(degree string, minute string, second string) float64 {
}
