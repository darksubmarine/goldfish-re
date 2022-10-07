package goldfish_re

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const emptyStr = ""

func objectName(fact string) string {
	if fact == emptyStr {
		return emptyStr
	}

	parts := strings.Split(fact, ".")
	if len(parts) > 1 {
		return parts[0]
	}

	return fact
}

func conditionToken(left, right, operator string, negated bool) string {
	if negated {
		return fmt.Sprintf("!%s_%s_%s", left, operator, right)
	}
	return fmt.Sprintf("%s_%s_%s", left, operator, right)
}

func growthSlice[T interface{}](s []T, size int) []T {
	return append(s, make([]T, size)...)
}

func termType(v interface{}) tTerm {
	switch v.(type) {
	case string, []string:
		return termString
	case int, int64:
		return termNumber
	case float64:
		return termFloat
	case bool:
		return termBoolean
	case time.Time, []time.Time:
		return termDate
	default:
		return termInvalid
	}
}

func indexPath(object, attribute string, value interface{}) string {
	return fmt.Sprintf("/%s/%s/%v", object, attribute, value)
}

func indexPathFact(fact iFact) string {
	return indexPath(fact.object(), fact.attribute(), fact.value())
}

func parseIntOrDefault(s string, def int64) int64 {
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return n
	}
	return def
}

func parseBooleanOrDefault(s string, def bool) bool {
	if b, err := strconv.ParseBool(s); err == nil {
		return b
	}
	return def
}

func parseDateOrDefault(s string, def time.Time) time.Time {
	const DATELAYUOT = "2006-01-02T15:04:05"
	if t, err := time.ParseInLocation(DATELAYUOT, s, time.UTC); err == nil {
		return t
	}
	return def
}

func parseFloatOrDefault(s string, def float64) float64 {
	if n, err := strconv.ParseFloat(s, 64); err == nil {
		return n
	}
	return def
}

func YearUTC(y int) time.Time {
	return time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC)
}

func CalendarDateUTC(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func CalendarFullDateUTC(y, m, d, h, mm, s, ns int) time.Time {
	return time.Date(y, time.Month(m), d, h, m, s, ns, time.UTC)
}
