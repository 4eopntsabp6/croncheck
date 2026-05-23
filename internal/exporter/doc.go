// Package exporter provides utilities to export cron schedule run data
// into structured formats for downstream consumption or archival.
//
// Supported formats:
//
//   - JSON: structured array of run entries with index, absolute time,
//     and human-readable relative time.
//
//   - CSV: comma-separated rows with a header line, suitable for
//     spreadsheet import or pipeline processing.
//
// Example usage:
//
//	runs := simulator.NextRuns(expr, time.Now(), 5)
//
//	// Export as JSON
//	if err := exporter.ExportJSON(os.Stdout, runs, time.Now()); err != nil {
//		log.Fatal(err)
//	}
//
//	// Export as CSV
//	if err := exporter.ExportCSV(os.Stdout, runs, time.Now()); err != nil {
//		log.Fatal(err)
//	}
package exporter
