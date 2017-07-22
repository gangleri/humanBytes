// Package humanbytes is used to convert bytes to and from human readable representations. It can convert a string such as
// '1 MB' into bytes. Convert bytes to another unit eg KB, MB, GB... and also generate a human readable representation
// of the bytes e.g. 1024 would be displayed as '1 KB'.
package humanbytes //import "gangleri.io/pkg/humanbytes"

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

// silly change

// constants used for conversions between different units
const (
	bytes      = 1.0
	kilobytes  = 1024 * bytes
	megabytes  = 1024 * kilobytes
	gigabytes  = 1024 * megabytes
	terabytes  = 1024 * gigabytes
	zettabytes = 1024 * terabytes
)

// map used to get the multiplier to convert each of the standard prefixes
var conversion = map[string]float64{
	"B":  bytes,
	"KB": kilobytes,
	"MB": megabytes,
	"GB": gigabytes,
	"TB": terabytes,
	"ZB": zettabytes}

// ParseBytes takes a string with a unit prefix (e.g. MB) and converts the value to bytes
func ParseBytes(str string) (int, error) {
	re := regexp.MustCompile(`([0-9.]+)\ ?(B|KB|MB|GB|TB|PB|ZB)$`)
	m := re.FindAllStringSubmatch(str, -1)

	if len(m) != 1 || len(m[0]) != 3 {
		return 0, errors.Errorf("Invalid string [%s] unable to parse", str)
	}

	// regex ensure this will be a valid number
	i, _ := strconv.ParseFloat(m[0][1], 64)

	multiplier := conversion[m[0][2]]
	return int(i * multiplier), nil
}

// Convert takes the number of bytes and converts it to the equivalent in the unit identified by the type
// string parameter e.g. 'KB' or 'MB'
func Convert(b int, t string) (float64, error) {
	if _, ok := conversion[t]; !ok {
		return 0, errors.Errorf("Invalid type %s unable to perform conversion", t)
	}
	return float64(b) / conversion[t], nil
}

// ConvertToBytes takes a value in the specified unit and converts it to bytes
func ConvertToBytes(x float64, t string) (int, error) {
	if _, ok := conversion[t]; !ok {
		return 0, errors.Errorf("Invalid type %s unable to perform conversion", t)
	}
	return int(x * conversion[t]), nil
}

// Sprint generates a string representation using the specified unit
func Sprint(b int, t string) (string, error) {
	if _, ok := conversion[t]; !ok {
		return "", errors.Errorf("Invalid type %s unable to convert to string", t)
	}
	s := strconv.FormatFloat(float64(b)/conversion[t], 'f', -1, 64)
	return fmt.Sprintf("%s %s", s, t), nil
}
