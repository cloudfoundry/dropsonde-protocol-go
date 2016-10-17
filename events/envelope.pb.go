// Code generated by protoc-gen-gogo.
// source: envelope.proto
// DO NOT EDIT!

/*
	Package events is a generated protocol buffer package.

	It is generated from these files:
		envelope.proto
		error.proto
		http.proto
		log.proto
		metric.proto
		uuid.proto

	It has these top-level messages:
		Envelope
		Error
		HttpStartStop
		LogMessage
		ValueMetric
		CounterEvent
		ContainerMetric
		UUID
*/
package events

import proto "github.com/scalingdata/gogo-protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/scalingdata/gogo-protobuf/gogoproto"

import github_com_gogo_protobuf_proto "github.com/scalingdata/gogo-protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// / Type of the wrapped event.
type Envelope_EventType int32

const (
	// Removed Heartbeat at position 1
	// Removed HttpStart at position 2
	// Removed HttpStop at position 3
	Envelope_HttpStartStop   Envelope_EventType = 4
	Envelope_LogMessage      Envelope_EventType = 5
	Envelope_ValueMetric     Envelope_EventType = 6
	Envelope_CounterEvent    Envelope_EventType = 7
	Envelope_Error           Envelope_EventType = 8
	Envelope_ContainerMetric Envelope_EventType = 9
)

var Envelope_EventType_name = map[int32]string{
	4: "HttpStartStop",
	5: "LogMessage",
	6: "ValueMetric",
	7: "CounterEvent",
	8: "Error",
	9: "ContainerMetric",
}
var Envelope_EventType_value = map[string]int32{
	"HttpStartStop":   4,
	"LogMessage":      5,
	"ValueMetric":     6,
	"CounterEvent":    7,
	"Error":           8,
	"ContainerMetric": 9,
}

func (x Envelope_EventType) Enum() *Envelope_EventType {
	p := new(Envelope_EventType)
	*p = x
	return p
}
func (x Envelope_EventType) String() string {
	return proto.EnumName(Envelope_EventType_name, int32(x))
}
func (x *Envelope_EventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Envelope_EventType_value, data, "Envelope_EventType")
	if err != nil {
		return err
	}
	*x = Envelope_EventType(value)
	return nil
}
func (Envelope_EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEnvelope, []int{0, 0} }

// / Envelope wraps an Event and adds metadata.
type Envelope struct {
	Origin     *string             `protobuf:"bytes,1,req,name=origin" json:"origin,omitempty"`
	EventType  *Envelope_EventType `protobuf:"varint,2,req,name=eventType,enum=events.Envelope_EventType" json:"eventType,omitempty"`
	Timestamp  *int64              `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Deployment *string             `protobuf:"bytes,13,opt,name=deployment" json:"deployment,omitempty"`
	Job        *string             `protobuf:"bytes,14,opt,name=job" json:"job,omitempty"`
	Index      *string             `protobuf:"bytes,15,opt,name=index" json:"index,omitempty"`
	Ip         *string             `protobuf:"bytes,16,opt,name=ip" json:"ip,omitempty"`
	Tags       map[string]string   `protobuf:"bytes,17,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Removed Heartbeat at position 3
	// Removed HttpStart at position 4
	// Removed HttpStop at position 5
	HttpStartStop    *HttpStartStop   `protobuf:"bytes,7,opt,name=httpStartStop" json:"httpStartStop,omitempty"`
	LogMessage       *LogMessage      `protobuf:"bytes,8,opt,name=logMessage" json:"logMessage,omitempty"`
	ValueMetric      *ValueMetric     `protobuf:"bytes,9,opt,name=valueMetric" json:"valueMetric,omitempty"`
	CounterEvent     *CounterEvent    `protobuf:"bytes,10,opt,name=counterEvent" json:"counterEvent,omitempty"`
	Error            *Error           `protobuf:"bytes,11,opt,name=error" json:"error,omitempty"`
	ContainerMetric  *ContainerMetric `protobuf:"bytes,12,opt,name=containerMetric" json:"containerMetric,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptorEnvelope, []int{0} }

func (m *Envelope) GetOrigin() string {
	if m != nil && m.Origin != nil {
		return *m.Origin
	}
	return ""
}

func (m *Envelope) GetEventType() Envelope_EventType {
	if m != nil && m.EventType != nil {
		return *m.EventType
	}
	return Envelope_HttpStartStop
}

func (m *Envelope) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Envelope) GetDeployment() string {
	if m != nil && m.Deployment != nil {
		return *m.Deployment
	}
	return ""
}

func (m *Envelope) GetJob() string {
	if m != nil && m.Job != nil {
		return *m.Job
	}
	return ""
}

func (m *Envelope) GetIndex() string {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return ""
}

func (m *Envelope) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Envelope) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Envelope) GetHttpStartStop() *HttpStartStop {
	if m != nil {
		return m.HttpStartStop
	}
	return nil
}

func (m *Envelope) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

func (m *Envelope) GetValueMetric() *ValueMetric {
	if m != nil {
		return m.ValueMetric
	}
	return nil
}

func (m *Envelope) GetCounterEvent() *CounterEvent {
	if m != nil {
		return m.CounterEvent
	}
	return nil
}

func (m *Envelope) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Envelope) GetContainerMetric() *ContainerMetric {
	if m != nil {
		return m.ContainerMetric
	}
	return nil
}

func init() {
	proto.RegisterType((*Envelope)(nil), "events.Envelope")
	proto.RegisterEnum("events.Envelope_EventType", Envelope_EventType_name, Envelope_EventType_value)
}
func (m *Envelope) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Envelope) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Origin == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("origin")
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Origin)))
		i += copy(data[i:], *m.Origin)
	}
	if m.EventType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("eventType")
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintEnvelope(data, i, uint64(*m.EventType))
	}
	if m.Timestamp != nil {
		data[i] = 0x30
		i++
		i = encodeVarintEnvelope(data, i, uint64(*m.Timestamp))
	}
	if m.HttpStartStop != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.HttpStartStop.Size()))
		n1, err := m.HttpStartStop.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.LogMessage != nil {
		data[i] = 0x42
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.LogMessage.Size()))
		n2, err := m.LogMessage.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.ValueMetric != nil {
		data[i] = 0x4a
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.ValueMetric.Size()))
		n3, err := m.ValueMetric.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.CounterEvent != nil {
		data[i] = 0x52
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.CounterEvent.Size()))
		n4, err := m.CounterEvent.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Error != nil {
		data[i] = 0x5a
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.Error.Size()))
		n5, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.ContainerMetric != nil {
		data[i] = 0x62
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.ContainerMetric.Size()))
		n6, err := m.ContainerMetric.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Deployment != nil {
		data[i] = 0x6a
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Deployment)))
		i += copy(data[i:], *m.Deployment)
	}
	if m.Job != nil {
		data[i] = 0x72
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Job)))
		i += copy(data[i:], *m.Job)
	}
	if m.Index != nil {
		data[i] = 0x7a
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Index)))
		i += copy(data[i:], *m.Index)
	}
	if m.Ip != nil {
		data[i] = 0x82
		i++
		data[i] = 0x1
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Ip)))
		i += copy(data[i:], *m.Ip)
	}
	if len(m.Tags) > 0 {
		for k, _ := range m.Tags {
			data[i] = 0x8a
			i++
			data[i] = 0x1
			i++
			v := m.Tags[k]
			mapSize := 1 + len(k) + sovEnvelope(uint64(len(k))) + 1 + len(v) + sovEnvelope(uint64(len(v)))
			i = encodeVarintEnvelope(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintEnvelope(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintEnvelope(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Envelope(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Envelope(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEnvelope(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Envelope) Size() (n int) {
	var l int
	_ = l
	if m.Origin != nil {
		l = len(*m.Origin)
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.EventType != nil {
		n += 1 + sovEnvelope(uint64(*m.EventType))
	}
	if m.Timestamp != nil {
		n += 1 + sovEnvelope(uint64(*m.Timestamp))
	}
	if m.HttpStartStop != nil {
		l = m.HttpStartStop.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.LogMessage != nil {
		l = m.LogMessage.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.ValueMetric != nil {
		l = m.ValueMetric.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.CounterEvent != nil {
		l = m.CounterEvent.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.ContainerMetric != nil {
		l = m.ContainerMetric.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Deployment != nil {
		l = len(*m.Deployment)
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Job != nil {
		l = len(*m.Job)
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Index != nil {
		l = len(*m.Index)
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Ip != nil {
		l = len(*m.Ip)
		n += 2 + l + sovEnvelope(uint64(l))
	}
	if len(m.Tags) > 0 {
		for k, v := range m.Tags {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovEnvelope(uint64(len(k))) + 1 + len(v) + sovEnvelope(uint64(len(v)))
			n += mapEntrySize + 2 + sovEnvelope(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEnvelope(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEnvelope(x uint64) (n int) {
	return sovEnvelope(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Envelope) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnvelope
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Origin = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			var v Envelope_EventType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (Envelope_EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EventType = &v
			hasFields[0] |= uint64(0x00000002)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Timestamp = &v
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpStartStop", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpStartStop == nil {
				m.HttpStartStop = &HttpStartStop{}
			}
			if err := m.HttpStartStop.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LogMessage == nil {
				m.LogMessage = &LogMessage{}
			}
			if err := m.LogMessage.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueMetric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValueMetric == nil {
				m.ValueMetric = &ValueMetric{}
			}
			if err := m.ValueMetric.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CounterEvent == nil {
				m.CounterEvent = &CounterEvent{}
			}
			if err := m.CounterEvent.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerMetric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContainerMetric == nil {
				m.ContainerMetric = &ContainerMetric{}
			}
			if err := m.ContainerMetric.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Deployment = &s
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Job", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Job = &s
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Index = &s
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Ip = &s
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthEnvelope
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapvalue uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapvalue |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapvalue := int(stringLenmapvalue)
			if intStringLenmapvalue < 0 {
				return ErrInvalidLengthEnvelope
			}
			postStringIndexmapvalue := iNdEx + intStringLenmapvalue
			if postStringIndexmapvalue > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := string(data[iNdEx:postStringIndexmapvalue])
			iNdEx = postStringIndexmapvalue
			if m.Tags == nil {
				m.Tags = make(map[string]string)
			}
			m.Tags[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnvelope(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnvelope
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("origin")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("eventType")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEnvelope(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnvelope
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEnvelope
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEnvelope
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEnvelope(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEnvelope = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnvelope   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envelope.proto", fileDescriptorEnvelope) }

var fileDescriptorEnvelope = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4f, 0x6f, 0x9b, 0x4e,
	0x10, 0x15, 0x38, 0x76, 0xcc, 0xe0, 0x3f, 0x64, 0x92, 0xdf, 0xaf, 0x2b, 0xab, 0xb2, 0x68, 0x7a,
	0xe1, 0x52, 0x22, 0x59, 0xaa, 0x6a, 0xb5, 0xea, 0xa1, 0xa9, 0x5c, 0xf5, 0xd0, 0x5c, 0x48, 0xd4,
	0x3b, 0x86, 0x0d, 0xa1, 0xc1, 0x2c, 0x5d, 0x16, 0xab, 0x7c, 0xc3, 0x1e, 0xfb, 0x11, 0x2a, 0x7f,
	0x8a, 0x1e, 0x2b, 0x06, 0xb0, 0x71, 0xd4, 0xdb, 0xbc, 0x79, 0xef, 0x31, 0xb3, 0x6f, 0x80, 0x09,
	0x4f, 0xb7, 0x3c, 0x11, 0x19, 0x77, 0x33, 0x29, 0x94, 0xc0, 0x01, 0xdf, 0xf2, 0x54, 0xe5, 0xb3,
	0x57, 0x51, 0xac, 0x1e, 0x8a, 0xb5, 0x1b, 0x88, 0xcd, 0x55, 0x24, 0x22, 0x71, 0x45, 0xf4, 0xba,
	0xb8, 0x27, 0x44, 0x80, 0xaa, 0xda, 0x36, 0x83, 0x07, 0xa5, 0xb2, 0xa6, 0x36, 0x12, 0x11, 0x35,
	0xe5, 0x68, 0xc3, 0x95, 0x8c, 0x83, 0x06, 0x99, 0x5c, 0x4a, 0x21, 0x6b, 0x70, 0xf9, 0xa7, 0x0f,
	0xc3, 0x55, 0x33, 0x1b, 0xff, 0x87, 0x81, 0x90, 0x71, 0x14, 0xa7, 0x4c, 0xb3, 0x75, 0xc7, 0xf0,
	0x1a, 0x84, 0x4b, 0x30, 0x68, 0x9f, 0xbb, 0x32, 0xe3, 0x4c, 0xb7, 0x75, 0x67, 0xb2, 0x98, 0xb9,
	0xf5, 0x86, 0x6e, 0x6b, 0x76, 0x57, 0xad, 0xc2, 0x3b, 0x88, 0xf1, 0x39, 0x18, 0x2a, 0xde, 0xf0,
	0x5c, 0xf9, 0x9b, 0x8c, 0x0d, 0x6c, 0xcd, 0xe9, 0x79, 0x87, 0x06, 0xbe, 0x83, 0x71, 0xb5, 0xf0,
	0xad, 0xf2, 0xa5, 0xba, 0x55, 0x22, 0x63, 0xa7, 0xb6, 0xe6, 0x98, 0x8b, 0xff, 0xda, 0x6f, 0x7f,
	0xee, 0x92, 0xde, 0xb1, 0x16, 0x17, 0x00, 0x89, 0x88, 0x6e, 0x78, 0x9e, 0xfb, 0x11, 0x67, 0x43,
	0x72, 0x62, 0xeb, 0xfc, 0xb2, 0x67, 0xbc, 0x8e, 0x0a, 0x5f, 0x83, 0xb9, 0xf5, 0x93, 0x82, 0xdf,
	0x50, 0x1e, 0xcc, 0x20, 0xd3, 0x79, 0x6b, 0xfa, 0x7a, 0xa0, 0xbc, 0xae, 0x0e, 0x97, 0x30, 0x0a,
	0x44, 0x91, 0x2a, 0x2e, 0xe9, 0x91, 0x0c, 0xc8, 0x77, 0xd1, 0xfa, 0x3e, 0x76, 0x38, 0xef, 0x48,
	0x89, 0x2f, 0xa1, 0x4f, 0x69, 0x33, 0x93, 0x2c, 0xe3, 0x7d, 0x6a, 0x55, 0xd3, 0xab, 0x39, 0xfc,
	0x00, 0xd3, 0x40, 0xa4, 0xca, 0x8f, 0x53, 0x2e, 0x9b, 0xcd, 0x46, 0x24, 0x7f, 0x76, 0x98, 0x70,
	0x44, 0x7b, 0x4f, 0xf5, 0x38, 0x07, 0x08, 0x79, 0x96, 0x88, 0x72, 0x53, 0xed, 0x37, 0xb6, 0x35,
	0xc7, 0xf0, 0x3a, 0x1d, 0xb4, 0xa0, 0xf7, 0x4d, 0xac, 0xd9, 0x84, 0x88, 0xaa, 0xc4, 0x0b, 0xe8,
	0xc7, 0x69, 0xc8, 0x7f, 0xb0, 0x29, 0xf5, 0x6a, 0x80, 0x13, 0xd0, 0xe3, 0x8c, 0x59, 0xd4, 0xd2,
	0xe3, 0x0c, 0x5d, 0x38, 0x51, 0x7e, 0x94, 0xb3, 0x33, 0xbb, 0xe7, 0x98, 0xff, 0x38, 0xfa, 0x9d,
	0x1f, 0xe5, 0xab, 0x54, 0xc9, 0xd2, 0x23, 0xdd, 0xec, 0x0d, 0x18, 0xfb, 0x56, 0x35, 0xf4, 0x91,
	0x97, 0x4c, 0xab, 0x87, 0x3e, 0xf2, 0xb2, 0x1a, 0x4a, 0xb9, 0x32, 0xbd, 0x1e, 0x4a, 0xe0, 0xad,
	0xbe, 0xd4, 0x2e, 0xbf, 0x83, 0xb1, 0xff, 0x81, 0xf0, 0x0c, 0xc6, 0x47, 0xa7, 0xb7, 0x4e, 0x70,
	0x02, 0x70, 0xb8, 0xa9, 0xd5, 0xc7, 0x29, 0x98, 0x9d, 0x73, 0x59, 0x03, 0xb4, 0x60, 0xd4, 0xbd,
	0x83, 0x75, 0x8a, 0x06, 0xf4, 0x29, 0x66, 0x6b, 0x88, 0xe7, 0x30, 0x7d, 0x12, 0xa1, 0x65, 0x5c,
	0xbf, 0xff, 0xb9, 0x9b, 0x6b, 0xbf, 0x76, 0x73, 0xed, 0xf7, 0x6e, 0xae, 0xc1, 0x0b, 0x21, 0x23,
	0x37, 0x48, 0x44, 0x11, 0xde, 0x8b, 0x22, 0x0d, 0x65, 0xe9, 0x86, 0x52, 0x64, 0xb9, 0x48, 0x43,
	0xde, 0xbc, 0xfa, 0x7a, 0x44, 0x5f, 0xfe, 0xe4, 0x07, 0x4a, 0xc8, 0xf2, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x44, 0xf5, 0x54, 0xc4, 0xb3, 0x03, 0x00, 0x00,
}
