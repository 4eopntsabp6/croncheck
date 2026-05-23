# croncheck

A utility to validate, visualize, and simulate cron expressions with next-run previews.

---

## Installation

```bash
go install github.com/yourname/croncheck@latest
```

Or build from source:

```bash
git clone https://github.com/yourname/croncheck.git && cd croncheck && go build -o croncheck .
```

---

## Usage

```bash
# Validate a cron expression
croncheck validate "*/5 * * * *"

# Preview the next N scheduled run times
croncheck next "0 9 * * MON-FRI" --count 5

# Visualize a cron expression in human-readable form
croncheck explain "30 18 1 * *"
```

**Example output:**

```
Expression : 30 18 1 * *
Description: At 6:30 PM on day 1 of every month

Next runs:
  1. 2024-02-01 18:30:00
  2. 2024-03-01 18:30:00
  3. 2024-04-01 18:30:00
```

---

## Supported Fields

| Field        | Allowed Values  | Special Characters |
|--------------|-----------------|--------------------|
| Minute       | 0–59            | `* , - /`          |
| Hour         | 0–23            | `* , - /`          |
| Day of Month | 1–31            | `* , - /`          |
| Month        | 1–12 or JAN–DEC | `* , - /`          |
| Day of Week  | 0–6 or SUN–SAT  | `* , - /`          |

---

## License

This project is licensed under the [MIT License](LICENSE).