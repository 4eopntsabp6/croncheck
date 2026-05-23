// Package differ provides utilities for comparing two cron expressions
// side by side, highlighting field-level differences and divergence in
// their next scheduled run times.
//
// # Usage
//
// Use [Compare] to produce a [Diff] value that contains:
//   - Per-field comparison (Minute, Hour, Day, Month, Weekday)
//   - A list of upcoming [RunDiff] entries showing how far apart each
//     expression fires relative to a common starting time.
//
// Example:
//
//	d := differ.Compare("0 9 * * 1-5", "0 10 * * 1-5", time.Now(), 5)
//	for _, fd := range d.FieldDiffs {
//	    fmt.Printf("%-10s A=%-10s B=%-10s same=%v\n",
//	        fd.Field, fd.ValueA, fd.ValueB, fd.Same)
//	}
//	for _, rd := range d.NextRunDiffs {
//	    fmt.Printf("Run %d: A=%s  B=%s  delta=%v\n",
//	        rd.Index, rd.RunA.Format(time.RFC3339),
//	        rd.RunB.Format(time.RFC3339), rd.Delta)
//	}
package differ
