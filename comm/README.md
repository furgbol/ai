# Communication Package

This package provides types and methods for exchanging data over the network. Supports UDP protocols.

# Contents

1. [UDP Sender](#udpsender)
2. [UDP Receiver](#udpreceiver)

<a name="udpsender"></a>
## UDP Sender

Sends data (byte) to an address on the network.

### Sender Methods

- **NewUDPSender**(*addr string*) (**UDPSender, error*)
- **WriteDatagram**(*bytes []byte*) (*int, error*)
- **Close**() (*error*)

<a name="udpreceiver"></a>
## UDP Receiver

Receives data (byte) over the network.

### Receiver Methods

- **NewUDPReceiver**(*addr string*) (**UDPReceiver, error*)
- **ReceiveDatagram**() (*[]byte, error*)
- **Close**() (*error*)