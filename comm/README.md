# Communication Package

This package provides types and methods for exchanging data over the network. Supports UDP and TCP protocols.

# Contents

1. [Driver UDP](#udp)
2. [ZMQ](#zmq)

<a name="udp"></a>
## UDP Sender

Sends data (byte) to an address on the network.

### Sender Methods

- **NewUDPSender**(*addr string*) (**UDPSender, error*)
- **WriteDatagram**(*bytes []byte*) (*int, error*)
- **Close**() (*error*)

## UDP Receiver

Receives data (byte) over the network.

### Receiver Methods

- **NewUDPReceiver**(*addr string*) (**UDPReceiver, error*)
- **ReceiveDatagram**() (*[]byte, error*)
- **Close**() (*error*)

<a name="zmq"></a>
## ZMQ

Using the ZeroMQ library to exchange data on the network. Used TCP protocol.

## Constants

- *PUB*
- *SUB*
- *PAIR*

## ZMQ Client

Connects to a server to exchange data.

### ZMQ Client Methods

- **NewZMQClient**(*addr string, tp zmq.Type*) (**ZMQClient, error*)
- **Send**(*data []byte*) (*int, error*)
- **Receive**() (*[]byte, error*)
- **Close**() (*error*)

## ZMQ Server

Creates a socket for data exchange.

### ZMQ Server Methods

- **NewZMQServer**(*addr string, tp zmq.Type*) (**ZMQClient, error*)
- **Send**(*data []byte*) (*int, error*)
- **Receive**() (*[]byte, error*)
- **Close**() (*error*)