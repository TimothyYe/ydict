package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsChinese(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Init a Chinese string", t, func() {
		str := "测试"

		Convey("Call isChinese func", func() {
			result := isChinese(str)

			Convey("result should be: true", func() {
				So(result, ShouldEqual, true)
			})
		})
	})
}

//func parseArgs(args []string) ([]string, bool) {
////match argument: -v
//if args[len(args)-1] == "-v" {
//return args[1 : len(args)-1], true
//}

//return args[1:], false
//}

func TestParseArgs(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Init args arrays: withVoice & withOutVoice", t, func() {
		withoutAll := []string{"aa", "bb", "cc"}
		withAll := []string{"aa", "bb", "-v", "-m"}

		withVoice := []string{"aa", "bb", "-v"}
		withMore := []string{"aa", "bb", "-m"}

		Convey("Call parse func", func() {
			_, ret1, ret2 := parseArgs(withoutAll)
			_, ret3, ret4 := parseArgs(withAll)

			_, ret5, ret6 := parseArgs(withVoice)
			_, ret7, ret8 := parseArgs(withMore)

			Convey("result should be: true & false", func() {
				So(ret1, ShouldEqual, false)
				So(ret2, ShouldEqual, false)

				So(ret3, ShouldEqual, true)
				So(ret4, ShouldEqual, true)

				So(ret5, ShouldEqual, true)
				So(ret6, ShouldEqual, false)

				So(ret7, ShouldEqual, false)
				So(ret8, ShouldEqual, true)
			})
		})
	})
}
