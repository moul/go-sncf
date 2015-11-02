package sncf

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTrainTimesDeparture(t *testing.T) {
	Convey("Testing GetTrainTimesDeparture", t, func() {
		result, err := GetTrainTimesDeparture("RRD")
		So(err, ShouldBeNil)
		trains := result.Trains
		So(len(trains), ShouldEqual, 20)
		firstTrain := trains[0]
		So(firstTrain.OrigDest, ShouldNotBeEmpty)
		So(firstTrain.Num, ShouldNotBeEmpty)
		So(firstTrain.Heure, ShouldNotBeEmpty)
	})
}
