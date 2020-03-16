package comm

import zmq "github.com/pebbe/zmq4"

const (
	//PUB - Used to specify the connection type (publisher)
	PUB = zmq.PUB
	//SUB - Used to specify the connection type (subscribe)
	SUB = zmq.SUB
	//PAIR - Used to specify the connection type (pair)
	PAIR = zmq.PAIR
)
