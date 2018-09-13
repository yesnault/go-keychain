// +build darwin ios

package keychain

/*
#cgo LDFLAGS: -framework CoreFoundation

#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import (
	"math"
	"time"
)

const nsPerSec = 1000000000

func unixToAbsoluteTime(s int64, ns int64) C.CFAbsoluteTime {
	return C.CFTimeInterval(ns)/nsPerSec + C.CFAbsoluteTime(s) + C.kCFAbsoluteTimeIntervalSince1970
}

func absoluteTimeToUnix(abs C.CFAbsoluteTime) (int64, int64) {
	int, frac := math.Modf(float64(abs - C.kCFAbsoluteTimeIntervalSince1970))
	return int64(int), int64(frac * nsPerSec)
}

func absoluteTimeToString(abs C.CFAbsoluteTime) string {
	dateFormatter := C.CFDateFormatterCreate(C.kCFAllocatorDefault, C.CFLocaleCopyCurrent(), C.kCFDateFormatterFullStyle, C.kCFDateFormatterFullStyle)
	defer Release(C.CFTypeRef(dateFormatter))
	cfStr := C.CFDateFormatterCreateStringWithAbsoluteTime(C.kCFAllocatorDefault, dateFormatter, abs)
	defer Release(C.CFTypeRef(cfStr))
	return CFStringToString(cfStr)
}

func cfDateToString(d C.CFDateRef) string {
	dateFormatter := C.CFDateFormatterCreate(C.kCFAllocatorDefault, C.CFLocaleCopyCurrent(), C.kCFDateFormatterFullStyle, C.kCFDateFormatterFullStyle)
	defer Release(C.CFTypeRef(dateFormatter))
	cfStr := C.CFDateFormatterCreateStringWithDate(C.kCFAllocatorDefault, dateFormatter, d)
	defer Release(C.CFTypeRef(cfStr))
	return CFStringToString(cfStr)
}

func releaseCFDateForTest(d C.CFDateRef) {
	Release(C.CFTypeRef(d))
}

func TimeToCFDate(t time.Time) C.CFDateRef {
	abs := unixToAbsoluteTime(t.Unix(), int64(t.Nanosecond()))
	return C.CFDateCreate(C.kCFAllocatorDefault, abs)
}

func CFDateToTime(d C.CFDateRef) time.Time {
	abs := C.CFDateGetAbsoluteTime(d)
	s, ns := absoluteTimeToUnix(abs)
	return time.Unix(s, ns)
}