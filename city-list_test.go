package sncf

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCityList(t *testing.T) {
	Convey("Testing GetCityList", t, func() {
		if os.Getenv("SKIP_NETWORK_TESTS") == "1" {
			t.Skip()
		}
		cities, err := GetCityList()
		So(err, ShouldBeNil)
		So(len(cities) >= 4233, ShouldBeTrue)
		So(cities[0], ShouldEqual, "Aéroport Paris - Roissy - Charles-de-Gaulle (CDG 2) - Gare TGV")
	})
}
