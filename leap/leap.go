// Package leap provides a function to determine if a year is a leap year.
package leap

// IsLeapYear determines if a year is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
