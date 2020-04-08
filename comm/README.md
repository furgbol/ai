# Communication Package

This package provides types and methods for exchanging data over the network. Supports UDP and TCP protocols.

# Contents

Data type and methods for UDP communication and the ZeroMQ library.

1. [UDP](#udp)
2. [ZMQ](#zmq)

UDP communication and the ZeroMQ library use the Sender and Receiver interfaces.

1. [Sender](#send)
2. [Receiver](#rec)

<a name="udp"></a>

## UDP Type

This type stores the connection data containing the following fields:

- *cnn()*

### UDPSender

Method for creating the UDP communication socket to send data:

- **NewUDPSender**(*addr string*) (**UDPSender, error*)

### UDPReceiver

Method for creating the UDP communication socket to receive data:

- **NewUDPReceiver**(*addr string*) (**UDPReceiver, error*)

<a name="zmq"></a>

## ZMQ Type

This type contains the following fields:

- *socket()*

### ZMQSender

Method to create the communication socket using ZMQ to send data:

- **NewZMQSender**(*addr string*) (**ZMQSender, error*)

### ZMQReceiver

Method to create the communication socket using ZMQ to receive data:

- **NewZMQReceiver**(*addr string*) (**ZMQReceiver, error*)

<a name="send"></a>
## Sender

The Sender interface contains methods for creating a socket to send data over the network.

### Sender Interface Methods

- **Send**(*bytes []byte*) (*int, error*)
- **Close**() (*error*)

<a name="rec"></a>
## Receiver

The Receiver interface contains methods for creating a socket to receive data over the network.

### Receiver Interface Methods

- **Receive**() (*[]byte, error*)
- **Close**() (*error*)