// Package history provides utilities for tracking and persisting
// previously validated cron expressions across croncheck sessions.
//
// # Overview
//
// The history package maintains an ordered log of cron expressions
// that have been checked by the user. Each entry records the
// expression, its human-readable description, whether it was valid,
// and the timestamp of the check.
//
// # Usage
//
// Load or create a history, add entries, and save back to disk:
//
//	h, err := history.Load(history.DefaultPath())
//	if err != nil {
//		log.Fatal(err)
//	}
//	h.Add("* * * * *", "every minute", true)
//	if err := history.Save(history.DefaultPath(), h); err != nil {
//		log.Fatal(err)
//	}
//
// The default storage location is ~/.croncheck/history.json.
package history
