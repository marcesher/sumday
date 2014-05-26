package sumday

import (
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
}

//var _ = Describe("SumDay", func() {
//	Context("initialization sets hour and minutes", func() {
//		t1 := NewTimeOfDay("8")
//		It("should set hour and minutes with single digit", func() {
//			Expect(t1.TimeString).To(Equal("8"))
//			Expect(t1.Err).To(BeNil())
//			Expect(t1.Hour).To(Equal(8))
//			Expect(t1.Minutes).To(Equal(00))
//		})
//		t2 := NewTimeOfDay("11")
//		It("should set hour and minutes with two digits", func() {
//			Expect(t2.TimeString).To(Equal("11"))
//			Expect(t2.Err).To(BeNil())
//			Expect(t2.Hour).To(Equal(11))
//			Expect(t2.Minutes).To(Equal(00))
//		})
//		t3 := NewTimeOfDay("130")
//		It("should set hour and minutes with three digits", func() {
//			Expect(t3.TimeString).To(Equal("130"))
//			Expect(t3.Err).To(BeNil())
//			Expect(t3.Hour).To(Equal(1))
//			Expect(t3.Minutes).To(Equal(30))
//		})
//		t4 := NewTimeOfDay("1245")
//		It("should set hour and minutes with four digits", func() {
//			Expect(t4.TimeString).To(Equal("1245"))
//			Expect(t4.Err).To(BeNil())
//			Expect(t4.Hour).To(Equal(12))
//			Expect(t4.Minutes).To(Equal(45))
//		})
//		terr := NewTimeOfDay("[a]130")
//		It("should set error with invalid input", func() {
//			Expect(terr.TimeString).To(Equal("[a]130"))
//			Expect(terr.Err).ToNot(BeNil())
//			Expect(terr.Hour).To(Equal(0))
//			Expect(terr.Minutes).To(Equal(0))
//		})
//	})
//
//	Context("MakeTimes ensures end after start", func() {
//		It("should create correct time span for start and end times", func() {
//			start, end := MakeTimes("830", "1015")
//			duration := end.Time.Sub(start.Time)
//			actual, _ := time.ParseDuration("1h45m")
//			Expect(duration).To(Equal(actual))
//			Expect(duration.Hours()).To(Equal(1.75))
//
//			start, end = MakeTimes("830", "115")
//			duration = end.Time.Sub(start.Time)
//			actual, _ = time.ParseDuration("4h45m")
//			Expect(duration.Hours()).To(Equal(actual.Hours()))
//			Expect(duration.Hours()).To(Equal(4.75))
//		})
//	})
//
//	Context("ParseLine makes TimeOfDays", func() {
//		It("should create TimeOfDays from strings such as ' - 2 - 3: RM: Weekly RM Meeting'", func() {
//			start, end, cat := ParseLine(" - 2 - 3: cfgov: Weekly RM Meeting")
//			fmt.Printf("start %v, end %v, cat %v\n", start, end, cat)
//			Expect(cat).To(Equal("cfgov"))
//			duration := end.Time.Sub(start.Time)
//			Expect(duration.Hours()).To(Equal(1.0))
//		})
//	})
//
//	Context("ParseLine makes TimeOfDays for lines that do not start with spaces", func() {
//		It("should create TimeOfDays from strings such as '- 2 - 3: RM: Weekly RM Meeting'", func() {
//			start, end, cat := ParseLine("- 2 - 3: cfgov: Weekly RM Meeting")
//			fmt.Printf("start %v, end %v, cat %v\n", start, end, cat)
//			Expect(cat).To(Equal("cfgov"))
//			duration := end.Time.Sub(start.Time)
//			Expect(duration.Hours()).To(Equal(1.0))
//		})
//	})
//
//	Context("SumDay sums durations for each category during a day", func() {
//
//		input :=
//			`
// - 9 - 915: Team: Email facilities about parking and whomever about monitor
// - 915 - 10: RM: Review weekly release notes
// - 10 - 11: cf.gov: Test counts for Bill
// - 11 - 1145: RM: RM planning: bippity boppity boo
// - 1230 - 1: Team: Regroup welcome
//  - 2 - 3: RM: Weekly RM Meeting
//           - Production patching... Qu got skipped so we're doing today
//           - GHE upgrade coming up
//           - Project Name: Whoosie whatsie
//
//  - 3 - 330: RM: Sally 1:1
//  - 330 - 4: RM: Johnny 1:1
//  - 4 - 415: CFGOV: Some thing
//  - 415 - 545: 20%: Fun stuff
//
//
//`
//		It("should group valid lines on categories", func() {
//			res := SumDay(input)
//			fmt.Printf("Result is %v\n\n", res)
//			exp := map[string]float64{
//				"team":  0.75,
//				"rm":    3.5,
//				"cfgov": 1.25,
//				"20":    1.5,
//			}
//			Expect(res).To(Equal(exp))
//		})
//
//	})
//})
