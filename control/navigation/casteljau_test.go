package navigation

import (
	"testing"

	"github.com/furgbol/ai/model"
)

type CasteljauTester struct {
	pathPlanner PathPlanner
}

func TestCasteljau(t *testing.T) {
	initialPose := model.Pose{
		Position2D: model.Position2D{
			X: 2.0,
			Y: 3.0,
		},
		Orientation: 67.3,
	}

	targetPose := model.Pose{
		Position2D: model.Position2D{
			X: 11.5,
			Y: 16.7,
		},
		Orientation: 12.145,
	}

	testingPath := Path{
		model.Position2D{
			X: 1.6027491507831224,
			Y: 1.404642768352762,
		},
		model.Position2D{
			X: 1.5008985494830758,
			Y: 0.9116882028722657,
		},
		model.Position2D{
			X: 1.6800405943177994,
			Y: 1.343056199976381,
		},
		model.Position2D{
			X: 2.1257676835052317,
			Y: 2.520666656082983,
		},
		model.Position2D{
			X: 2.823672215263313,
			Y: 4.266439467609941,
		},
		model.Position2D{
			X: 3.759346587809981,
			Y: 6.402294530975126,
		},
		model.Position2D{
			X: 4.918383199363176,
			Y: 8.750151742596413,
		},
		model.Position2D{
			X: 6.286374448140837,
			Y: 11.131930998891674,
		},
		model.Position2D{
			X: 7.848912732360902,
			Y: 13.369552196278775,
		},
		model.Position2D{
			X: 9.591590450241311,
			Y: 15.284935231175595,
		},
		model.Position2D{
			X: 11.500000000000004,
			Y: 16.700000000000003,
		},
	}

	tester := CasteljauTester{pathPlanner: NewCasteljauPathPlanner(11, 11, 2)}
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
