package entry

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/techplexengineer/frc-networktables-go/util"
)

// Raw entry
type Raw struct {
	Base
	trueValue    []byte
	isPersistent bool
}

// RawFromReader builds a raw entry using the provided parameters
func RawFromReader(name string, id [2]byte, sequence [2]byte, persist byte, reader io.Reader) (*Raw, error) {
	valLen, sizeData := util.PeekULeb128(reader)
	valData := make([]byte, valLen)
	_, err := io.ReadFull(reader, valData[:])
	if err != nil {
		return nil, err
	}
	persistent := (persist == flagPersist)
	value := append(sizeData, valData[:]...)
	return &Raw{
		trueValue:    valData[:],
		isPersistent: persistent,
		Base: Base{
			eName:  name,
			eType:  TypeRaw,
			eID:    id,
			eSeq:   sequence,
			eFlag:  persist,
			eValue: value,
		},
	}, nil
}

// RawFromItems builds a raw entry using the provided parameters
func RawFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value []byte) *Raw {
	valLen, sizeLen := util.ReadULeb128(bytes.NewReader(value))
	val := value[sizeLen : valLen-1]
	persistent := (persist == flagPersist)
	return &Raw{
		trueValue:    val,
		isPersistent: persistent,
		Base: Base{
			eName:  name,
			eType:  TypeRaw,
			eID:    id,
			eSeq:   sequence,
			eFlag:  persist,
			eValue: value,
		},
	}
}

// GetValue returns the raw value of this entry
func (o *Raw) GetValue() interface{} {
	return o.trueValue
}

// IsPersistent returns whether or not the entry should persist beyond restarts.
func (o *Raw) IsPersistent() bool {
	return o.isPersistent
}

// Clone returns an identical entry
func (o *Raw) Clone() *Raw {
	return &Raw{
		trueValue:    o.trueValue,
		isPersistent: o.isPersistent,
		Base:         o.Base.clone(),
	}
}

// CompressToBytes returns a byte slice representing the Raw entry
func (o *Raw) CompressToBytes() []byte {
	return o.Base.compressToBytes()
}

func (o Raw) GetName() string {
	return o.Base.eName
}
func (o Raw) GetID() uint16 {
	return binary.LittleEndian.Uint16(o.eID[:])
}
func (Raw) GetType() EntryType {
	return TypeRaw
}
func (o *Raw) SetValue(newValue interface{}) {

}
