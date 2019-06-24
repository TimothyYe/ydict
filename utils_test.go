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

func TestParseArgs(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Init args arrays: withVoice & withOutVoice & withCache", t, func() {
		withoutAll := []string{"aa", "bb", "cc"}
		withAll := []string{"-v", "-m", "-q", "-c", "-clear", "aa", "bb"}

		Convey("Call parse func", func() {
			words01, ret01, ret02, ret03, ret04, ret05 := parseArgs(withoutAll)
			words11, ret11, ret12, ret13, ret14, ret15 := parseArgs(withAll)

			Convey("result should be: true & false", func() {
				So(words01, ShouldContain, "aa")
				So(words01, ShouldContain, "bb")
				So(words01, ShouldContain, "cc")
				So(ret01, ShouldEqual, 0)
				So(ret02, ShouldEqual, false)
				So(ret03, ShouldEqual, false)
				So(ret04, ShouldEqual, false)
				So(ret05, ShouldEqual, false)

				So(words11, ShouldContain, "aa")
				So(words11, ShouldContain, "bb")
				So(words11, ShouldNotContain, "-v")
				So(words11, ShouldNotContain, "-m")
				So(words11, ShouldNotContain, "-q")
				So(words11, ShouldNotContain, "-c")
				So(words11, ShouldNotContain, "-clear")
				So(ret11, ShouldEqual, 1)
				So(ret12, ShouldEqual, true)
				So(ret13, ShouldEqual, true)
				So(ret14, ShouldEqual, true)
				So(ret15, ShouldEqual, true)
			})
		})
	})
}
