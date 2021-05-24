package quick

import (
	"bytes"
	"encoding/gob"
	"net"
)

type Parcel struct {
	Type    ParcelType
	Id      uint32
	Index   uint32
	Chunks  uint32
	Payload []byte
}

type ParcelWithAddress struct {
	*Parcel
	RemoteAddr *net.UDPAddr
}

func (parcel *Parcel) marshalGob() (buf []byte, err error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err = enc.Encode(parcel)
	buf = buffer.Bytes()

	return
}

func (parcel *Parcel) unmarshalGob(byt []byte) (err error) {
	enc := gob.NewDecoder(bytes.NewReader(byt))
	err = enc.Decode(&parcel)

	return
}