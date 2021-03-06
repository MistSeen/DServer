package packet

type IPacketReader interface {
	ReadByte() (ret byte, err error)
	ReadBool() (ret bool, err error)
	ReadBytes() (ret []byte, err error)
	ReadString() (ret string, err error)
	ReadUint8() (ret uint8, err error)
	ReadInt8() (ret int8, err error)
	ReadUint16() (ret uint16, err error)
	ReadInt16() (ret int16, err error)
	ReadUint24() (ret uint32, err error)
	ReadInt24() (ret int32, err error)
	ReadUint32() (ret uint32, err error)
	ReadInt32() (ret int32, err error)
	ReadUint64() (ret uint64, err error)
	ReadInt64() (ret int64, err error)
	ReadFloat32() (ret float32, err error)
	ReadFloat64() (ret float64, err error)
}

type IPacketWriter interface {
	WriteZeros(n int)
	WriteBool(v bool)
	WriteByte(v byte)
	WriteBytes(v []byte)
	WriteRawBytes(v []byte)
	WriteString(v string)
	WriteInt8(v int8)
	WriteUInt8(v uint8)
	WriteUint16(v uint16)
	WriteInt16(v int16)
	WriteUint24(v uint32)
	WriteUint32(v uint32)
	WriteInt32(v int32)
	WriteUint64(v uint64)
	WriteInt64(v int64)
	WriteFloat32(f float32)
	WriteFloat64(f float64)
}

type IUnpacker interface {

}