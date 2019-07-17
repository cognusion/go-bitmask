package bitmask

import (
	. "github.com/smartystreets/goconvey/convey"

	"fmt"
	"testing"
)

const (
	Alpha BitMask = 1 << iota
	Beta
	Gamma
)

const (
	Alpha64 BitMask64 = 1 << iota
	Beta64
	Gamma64
)

const (
	Alpha8 BitMask8 = 1 << iota
	Beta8
	Gamma8
)

func ExampleBitMask() {

	const (
		Alpha BitMask = 1 << iota
		Beta
		Gamma
	)

	var flags BitMask

	flags.Set(Alpha)
	flags.Set(Beta)

	if flags.IsSet(Gamma) {
		panic("Ohhhh noooooo!")
	}

	if flags.IsSet(Alpha) {
		fmt.Println("Alpha!")
	}
	// Output: Alpha!
}

func TestRawBitMask(t *testing.T) {
	flags := Alpha | Beta

	Convey("When a BitMask is set with Alpha and Beta, Alpha is set, Beta is set, Gamma is not set", t, func() {
		So(flags&Alpha, ShouldNotEqual, 0)
		So(flags&Beta, ShouldNotEqual, 0)
		So(flags&Gamma, ShouldEqual, 0)

		Convey("When Alpha is removed, Alpha is no longer set, Beta is still set, Gamma is still not set", func() {
			flags ^= Alpha

			So(flags&Alpha, ShouldEqual, 0)
			So(flags&Beta, ShouldNotEqual, 0)
			So(flags&Gamma, ShouldEqual, 0)

			Convey("When Gamma is added, Alpha is still not set, Beta is still set, Gamma is set", func() {
				flags |= Gamma

				So(flags&Alpha, ShouldEqual, 0)
				So(flags&Beta, ShouldNotEqual, 0)
				So(flags&Gamma, ShouldNotEqual, 0)
			})
		})

	})
}

func TestFunctionalBitMask(t *testing.T) {
	var flags BitMask

	flags.Set(Alpha)
	flags.Set(Beta)

	Convey("When a BitMask is set with Alpha and Beta, Alpha is set, Beta is set, Gamma is not set", t, func() {
		So(flags.IsSet(Alpha), ShouldBeTrue)
		So(flags.IsSet(Beta), ShouldBeTrue)
		So(flags.IsSet(Gamma), ShouldBeFalse)

		Convey("When Alpha is removed, Alpha is no longer set, Beta is still set, Gamma is still not set", func() {
			flags.Clear(Alpha)

			So(flags.IsSet(Alpha), ShouldBeFalse)
			So(flags.IsSet(Beta), ShouldBeTrue)
			So(flags.IsSet(Gamma), ShouldBeFalse)

			Convey("When Gamma is added, Alpha is still not set, Beta is still set, Gamma is set", func() {
				flags.Set(Gamma)

				So(flags.IsSet(Alpha), ShouldBeFalse)
				So(flags.IsSet(Beta), ShouldBeTrue)
				So(flags.IsSet(Gamma), ShouldBeTrue)
			})
		})

	})
}

func TestFunctionalBitMask64(t *testing.T) {
	var flags BitMask64

	flags.Set(Alpha64)
	flags.Set(Beta64)

	Convey("When a BitMask is set with Alpha and Beta, Alpha is set, Beta is set, Gamma is not set", t, func() {
		So(flags.IsSet(Alpha64), ShouldBeTrue)
		So(flags.IsSet(Beta64), ShouldBeTrue)
		So(flags.IsSet(Gamma64), ShouldBeFalse)

		Convey("When Alpha is removed, Alpha is no longer set, Beta is still set, Gamma is still not set", func() {
			flags.Clear(Alpha64)

			So(flags.IsSet(Alpha64), ShouldBeFalse)
			So(flags.IsSet(Beta64), ShouldBeTrue)
			So(flags.IsSet(Gamma64), ShouldBeFalse)

			Convey("When Gamma is added, Alpha is still not set, Beta is still set, Gamma is set", func() {
				flags.Set(Gamma64)

				So(flags.IsSet(Alpha64), ShouldBeFalse)
				So(flags.IsSet(Beta64), ShouldBeTrue)
				So(flags.IsSet(Gamma64), ShouldBeTrue)
			})
		})

	})
}

func TestFunctionalBitMask8(t *testing.T) {
	var flags BitMask8

	flags.Set(Alpha8)
	flags.Set(Beta8)

	Convey("When a BitMask is set with Alpha and Beta, Alpha is set, Beta is set, Gamma is not set", t, func() {
		So(flags.IsSet(Alpha8), ShouldBeTrue)
		So(flags.IsSet(Beta8), ShouldBeTrue)
		So(flags.IsSet(Gamma8), ShouldBeFalse)

		Convey("When Alpha is removed, Alpha is no longer set, Beta is still set, Gamma is still not set", func() {
			flags.Clear(Alpha8)

			So(flags.IsSet(Alpha8), ShouldBeFalse)
			So(flags.IsSet(Beta8), ShouldBeTrue)
			So(flags.IsSet(Gamma8), ShouldBeFalse)

			Convey("When Gamma is added, Alpha is still not set, Beta is still set, Gamma is set", func() {
				flags.Set(Gamma8)

				So(flags.IsSet(Alpha8), ShouldBeFalse)
				So(flags.IsSet(Beta8), ShouldBeTrue)
				So(flags.IsSet(Gamma8), ShouldBeTrue)
			})
		})

	})
}
