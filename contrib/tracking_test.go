package contrib

import (
	"image"
	"testing"

	"gocv.io/x/gocv"
)

func BaseTestTracker(t *testing.T, tracker Tracker, name string) {
	if tracker == nil {
		t.Error("TestTracker " + name + " should not be nil")
	}

	img, err := gocv.IMRead("../images/face.jpg", 1)
	if err != nil {
		t.Error("TestTracker " + name + " input img failed to load")
	}
	defer img.Close()

	rect := image.Rect(250, 150, 250+200, 150+250)
	init := tracker.Init(img, rect)
	if !init {
		t.Error("TestTracker " + name + " failed in Init")
	}

	_, ok := tracker.Update(img)
	if !ok {
		t.Error("TestTracker " + name + " lost object in Update")
	}
}

func TestSingleTrackers(t *testing.T) {
	tab := []struct {
		name    string
		tracker Tracker
	}{
		{"MIL", NewTrackerMIL()},
		{"Boosting", NewTrackerBoosting()},
		{"MedianFlow", NewTrackerMedianFlow()},
		{"TLD", NewTrackerTLD()},
		{"KCF", NewTrackerKCF()},
		{"MOSSE", NewTrackerMOSSE()},
		{"CSRT", NewTrackerCSRT()},
	}

	for _, test := range tab {
		defer test.tracker.Close()

		BaseTestTracker(t, test.tracker, test.name)
	}
}