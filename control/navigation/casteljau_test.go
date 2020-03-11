package navigation

import (
	"testing"

	"github.com/furgbol/ai/model"
)

type CasteljauTester struct {
	pathPlanner *PathPlanner
}

func TestCasteljau(t *testing.T) {
	initialPose := model.Pose{,
		model.Position2D{,
			X: 2.0,
			Y: 3.0,
		},
		Orientation: 67.3,
	}

	targetPose := model.Pose{,
		model.Position2D{,
			X: 11.5,
			Y: 16.7,
		},
		Orientation: 12.145,
	}

	testingPath := Path{,
		model.Position2D{,
			X: 2.8276595507336237,
			Y: 5.204934277280196,
		},
		model.Position2D{,
			X: 3.4409368941756404,
			Y: 7.209425483066546,
		},
		model.Position2D{,
			X: 3.930420395499267,
			Y: 9.019061912729224,
		},
		model.Position2D{,
			X: 4.386698419877721,
			Y: 10.63943186163841,
		},
		model.Position2D{,
			X: 4.9003593324842205,
			Y: 12.076123625164279,
		},
		model.Position2D{,
			X: 5.561991498491983,
			Y: 13.334725498677006,
		},
		model.Position2D{,
			X: 6.462183283074223,
			Y: 14.420825777546764,
		},
		model.Position2D{,
			X: 7.691523051404162,
			Y: 15.340012757143732,
		},
		model.Position2D{,
			X: 9.340599168655014,
			Y: 16.097874732838086,
		},
		model.Position2D{,
			X: 11.499999999999998,
			Y: 16.699999999999996,
		},
		model.Position2D{,
			X: 14.260313910612329,
			Y: 17.15197685399965,
		},
	}

	tester := CasteljauTester{pathPlanner: NewCasteljauPathPlanner(11, 11, 4)}
	path := tester.pathPlanner.PlanPath(initialPose, targetPose)

	if len(path) != 11 {
		t.Errorf("Error on testing casteljau's algorithm: paths size are differents. Expected: 11, got: %v", len(path))
	}

	for i := 0; i < len(testingPath); i++ {
		if path[i].X != testingPath[i].X {
			t.Errorf("Error on testing casteljau's algorithm: x of %v point are different. Expected: %v, got: %v", i, testingPath[i].X, path[i].X)
		}
		if path[i].Y != testingPath[i].Y {
			t.Errorf("Error on testing casteljau's algorithm: y of %v point are different. Expected: %v, got: %v", i, testingPath[i].Y, path[i].Y)
		}
	}
}
