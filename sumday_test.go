package sumday

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSumday(t *testing.T) {
	Convey("MakeTimes ensures end after start", t, func() {
		Convey("should create correct time span for start and end times", func() {
			start, end := MakeTimes("830", "1015")
			duration := end.Time.Sub(start.Time)
			actual, _ := time.ParseDuration("1h45m")
			So(duration, ShouldEqual, actual)
			So(duration.Hours(), ShouldEqual, 1.75)

			start, end = MakeTimes("830", "115")
			duration = end.Time.Sub(start.Time)
			actual, _ = time.ParseDuration("4h45m")
			So(duration.Hours(), ShouldEqual, actual.Hours())
			So(duration.Hours(), ShouldEqual, 4.75)
		})
	})
	Convey("initialization sets hour and minutes", t, func() {
		t1 := NewTimeOfDay("8")
		Convey("should set hour and minutes with single digit", func() {
			So(t1.TimeString, ShouldEqual, "8")
			So(t1.Err, ShouldBeNil)
			So(t1.Hour, ShouldEqual, 8)
			So(t1.Minutes, ShouldEqual, 00)
		})
		t2 := NewTimeOfDay("11")
		Convey("should set hour and minutes with two digits", func() {
			So(t2.TimeString, ShouldEqual, "11")
			So(t2.Err, ShouldBeNil)
			So(t2.Hour, ShouldEqual, 11)
			So(t2.Minutes, ShouldEqual, 00)
		})
		t3 := NewTimeOfDay("130")
		Convey("should set hour and minutes with three digits", func() {
			So(t3.TimeString, ShouldEqual, "130")
			So(t3.Err, ShouldBeNil)
			So(t3.Hour, ShouldEqual, 1)
			So(t3.Minutes, ShouldEqual, 30)
		})
		t4 := NewTimeOfDay("1245")
		Convey("should set hour and minutes with four digits", func() {
			So(t4.TimeString, ShouldEqual, "1245")
			So(t4.Err, ShouldEqual, nil)
			So(t4.Hour, ShouldEqual, 12)
			So(t4.Minutes, ShouldEqual, 45)
		})
		terr := NewTimeOfDay("[a]130")
		Convey("should set error with invalid input", func() {
			So(terr.TimeString, ShouldEqual, "[a]130")
			So(terr.Err, ShouldNotBeNil)
			So(terr.Hour, ShouldEqual, 0)
			So(terr.Minutes, ShouldEqual, 0)
		})
	})

	Convey("MakeTimes ensures end after start", t, func() {
		Convey("should create correct time span for start and end times", func() {
			start, end := MakeTimes("830", "1015")
			duration := end.Time.Sub(start.Time)
			actual, _ := time.ParseDuration("1h45m")
			So(duration, ShouldEqual, actual)
			So(duration.Hours(), ShouldEqual, 1.75)

			start, end = MakeTimes("830", "115")
			duration = end.Time.Sub(start.Time)
			actual, _ = time.ParseDuration("4h45m")
			So(duration.Hours(), ShouldEqual, actual.Hours())
			So(duration.Hours(), ShouldEqual, 4.75)
		})
	})

	Convey("ParseLine makes TimeOfDays", t, func() {
		Convey("should create TimeOfDays from strings such as ' - 2 - 3: RM: Weekly RM Meeting'", func() {
			start, end, cat := ParseLine(" - 2 - 3: cfgov: Weekly RM Meeting")
			fmt.Printf("start %v, end %v, cat %v\n", start, end, cat)
			So(cat, ShouldEqual, "cfgov")
			duration := end.Time.Sub(start.Time)
			So(duration.Hours(), ShouldEqual, 1.0)
		})
	})

	Convey("ParseLine makes TimeOfDays for lines that do not start with spaces", t, func() {
		Convey("should create TimeOfDays from strings such as '- 2 - 3: RM: Weekly RM Meeting'", func() {
			start, end, cat := ParseLine("- 2 - 3: cfgov: Weekly RM Meeting")
			fmt.Printf("start %v, end %v, cat %v\n", start, end, cat)
			So(cat, ShouldEqual, "cfgov")
			duration := end.Time.Sub(start.Time)
			So(duration.Hours(), ShouldEqual, 1.0)
		})
	})

	Convey("SumDay sums durations for each category during a day", t, func() {

		input :=
			`
- 9 - 915: Team: Email facilities about parking and whomever about monitor
- 915 - 10: RM: Review weekly release notes
- 10 - 11: cf.gov: Test counts for Bill
- 11 - 1145: RM: RM planning: bippity boppity boo
- 1230 - 1: Team: Regroup welcome
  - 2 - 3: RM: Weekly RM Meeting
           - Production patching... Qu got skipped so we're doing today
           - GHE upgrade coming up
           - Project Name: Whoosie whatsie

  - 3 - 330: RM: Sally 1:1
  - 330 - 4: RM: Johnny 1:1
  - 4 - 415: CFGOV: Some thing
  - 415 - 545: 20%: Fun stuff


`
		Convey("should group valid lines on categories", func() {
			res := SumDay(input)
			fmt.Printf("Result is %v\n\n", res)
			exp := map[string]float64{
				"team":  0.75,
				"rm":    3.5,
				"cfgov": 1.25,
				"20":    1.5,
			}
			So(res, ShouldResemble, exp)
		})
	})
}
