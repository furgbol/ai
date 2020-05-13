package debug

import (
	"log"

	"github.com/furgbol/ai/control/navigation"
	"github.com/furgbol/ai/model"
	"github.com/furgbol/ai/parse"

	"github.com/furgbol/ai/comm"
)

// PathPlannerDebugger - contains the necessary attributes for debugging
type PathPlannerDebugger struct {
	senderDebug   comm.Sender
	receiverState comm.Receiver
	encoder       parse.Encoder
	decoder       parse.Decoder
	pathPlanner   *navigation.CasteljauPathPlanner
}

// NewPathPlannerDebugger returns an instance of PathPlannerDebugger
func NewPathPlannerDebugger() PathPlannerDebugger {
	return PathPlannerDebugger{}
}

// Init initializes and pre-configures all debugging system variables
func (ppd *PathPlannerDebugger) Init(numberOfPathPoints, numberOfUsedPoints int, distanceFactor float64) {
	socketState, err := comm.NewZMQReceiver(comm.AddrState)
	socketDebug, err2 := comm.NewZMQSender(comm.AddrDebugTeamYellow)
	if err != nil {
		log.Fatal("[State Receiver Error]", err)
	} else if err2 != nil {
		log.Fatal("[Debug Sender Error]", err2)
	}

	ppd.senderDebug = socketDebug
	ppd.receiverState = socketState
	ppd.pathPlanner = navigation.NewCasteljauPathPlanner(numberOfPathPoints, numberOfUsedPoints, distanceFactor)
}

// Run start the debugging system
func (ppd *PathPlannerDebugger) Run() {
	defer ppd.senderDebug.Close()
	defer ppd.receiverState.Close()

	for {
		serial, err := ppd.receiverState.Receive()
		if err != nil {
			log.Fatal("[Receiver Error]", err)
		}

		ppd.decoder = parse.ProtobufToWorldModel{
			Team:         parse.TeamYellow,
			ReceivedData: serial,
		}

		state, err := ppd.decoder.Decode()
		if err != nil {
			log.Fatal("[Decoder Error]", err)
		}

		db := parse.DebugToProtobuf{}

		for i := 0; i < 3; i++ {
			db.InitialPoses = append(db.InitialPoses, state.Team[i].CurrentPosition)
			db.FinalPoses = append(db.FinalPoses,
				model.Pose{
					Position2D: model.Position2D{
						X: state.Ball.X,
						Y: state.Ball.Y,
					},
					Orientation: 0,
				})

			db.PathRobots = append(db.PathRobots, ppd.pathPlanner.PlanPath(db.InitialPoses[i], db.FinalPoses[i]))
		}

		ppd.encoder = db

		serial, err = ppd.encoder.Encode()
		if err != nil {
			log.Fatal("[Encoder Error]", err)
		}

		_, err = ppd.senderDebug.Send(serial)
		if err != nil {
			log.Fatal("[Sender Error]", err)
		}
	}
}
