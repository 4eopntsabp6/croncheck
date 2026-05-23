// Package validator provides field-level validation for standard 5-part cron
// expressions (minute hour day-of-month month day-of-week).
//
// Usage:
//
//	 result := validator.Validate("*/5 9-17 * * 1-5")
//	 if !result.Valid {
//	     for _, e := range result.Errors {
//	         fmt.Println(e)
//	     }
//	 }
//
// Supported field syntax:
//   - Wildcard:  *
//   - Value:     42
//   - Range:     1-5
//   - List:      1,2,3
//   - Step:      */15  or  0-30/5
//
// Allowed ranges per field:
//
//	minute       0-59
//	hour         0-23
//	day-of-month 1-31
//	month        1-12
//	day-of-week  0-6  (0 = Sunday)
package validator
