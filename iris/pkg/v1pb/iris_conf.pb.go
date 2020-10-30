// Code generated by protoc-gen-gogo.
// source: pkg/v1pb/iris_conf.proto
// DO NOT EDIT!

/*
	Package v1pb is a generated protocol buffer package.

	It is generated from these files:
		pkg/v1pb/iris_conf.proto

	It has these top-level messages:
		Config
		Iris
		Cluster
*/
package v1pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import envoy_api_v21 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
import envoy_api_v22 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
import envoy_api_v23 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type LogLevel int32

const (
	LogLevel_LOG_LEVEL_INVALID LogLevel = 0
	LogLevel_DEBUG             LogLevel = 1
	LogLevel_INFO              LogLevel = 2
	LogLevel_WARN              LogLevel = 3
	LogLevel_ERROR             LogLevel = 4
)

var LogLevel_name = map[int32]string{
	0: "LOG_LEVEL_INVALID",
	1: "DEBUG",
	2: "INFO",
	3: "WARN",
	4: "ERROR",
}
var LogLevel_value = map[string]int32{
	"LOG_LEVEL_INVALID": 0,
	"DEBUG":             1,
	"INFO":              2,
	"WARN":              3,
	"ERROR":             4,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}
func (LogLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptorIrisConf, []int{0} }

type Config struct {
	Iris      *Iris                               `protobuf:"bytes,1,opt,name=iris" json:"iris,omitempty"`
	Listeners []*envoy_api_v22.Listener           `protobuf:"bytes,2,rep,name=listeners" json:"listeners,omitempty"`
	Routes    []*envoy_api_v23.RouteConfiguration `protobuf:"bytes,3,rep,name=routes" json:"routes,omitempty"`
	Clusters  []*Cluster                          `protobuf:"bytes,4,rep,name=clusters" json:"clusters,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorIrisConf, []int{0} }

func (m *Config) GetIris() *Iris {
	if m != nil {
		return m.Iris
	}
	return nil
}

func (m *Config) GetListeners() []*envoy_api_v22.Listener {
	if m != nil {
		return m.Listeners
	}
	return nil
}

func (m *Config) GetRoutes() []*envoy_api_v23.RouteConfiguration {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *Config) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type Iris struct {
	Namespace          string         `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	LogLevel           LogLevel       `protobuf:"varint,2,opt,name=log_level,json=logLevel,proto3,enum=qubit.iris.v1.LogLevel" json:"log_level,omitempty"`
	Parallelism        uint32         `protobuf:"varint,3,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	DefaultPodPort     uint32         `protobuf:"varint,4,opt,name=default_pod_port,json=defaultPodPort,proto3" json:"default_pod_port,omitempty"`
	StateRefreshPeriod *time.Duration `protobuf:"bytes,5,opt,name=state_refresh_period,json=stateRefreshPeriod,stdduration" json:"state_refresh_period,omitempty"`
	StateBufferSize    uint32         `protobuf:"varint,6,opt,name=state_buffer_size,json=stateBufferSize,proto3" json:"state_buffer_size,omitempty"`
}

func (m *Iris) Reset()                    { *m = Iris{} }
func (m *Iris) String() string            { return proto.CompactTextString(m) }
func (*Iris) ProtoMessage()               {}
func (*Iris) Descriptor() ([]byte, []int) { return fileDescriptorIrisConf, []int{1} }

func (m *Iris) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Iris) GetLogLevel() LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return LogLevel_LOG_LEVEL_INVALID
}

func (m *Iris) GetParallelism() uint32 {
	if m != nil {
		return m.Parallelism
	}
	return 0
}

func (m *Iris) GetDefaultPodPort() uint32 {
	if m != nil {
		return m.DefaultPodPort
	}
	return 0
}

func (m *Iris) GetStateRefreshPeriod() *time.Duration {
	if m != nil {
		return m.StateRefreshPeriod
	}
	return nil
}

func (m *Iris) GetStateBufferSize() uint32 {
	if m != nil {
		return m.StateBufferSize
	}
	return 0
}

type Cluster struct {
	ServiceName string                 `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	PortName    string                 `protobuf:"bytes,2,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	Config      *envoy_api_v21.Cluster `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptorIrisConf, []int{2} }

func (m *Cluster) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Cluster) GetPortName() string {
	if m != nil {
		return m.PortName
	}
	return ""
}

func (m *Cluster) GetConfig() *envoy_api_v21.Cluster {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "qubit.iris.v1.Config")
	proto.RegisterType((*Iris)(nil), "qubit.iris.v1.Iris")
	proto.RegisterType((*Cluster)(nil), "qubit.iris.v1.Cluster")
	proto.RegisterEnum("qubit.iris.v1.LogLevel", LogLevel_name, LogLevel_value)
}
func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Iris != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(m.Iris.Size()))
		n1, err := m.Iris.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Listeners) > 0 {
		for _, msg := range m.Listeners {
			dAtA[i] = 0x12
			i++
			i = encodeVarintIrisConf(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Routes) > 0 {
		for _, msg := range m.Routes {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintIrisConf(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Clusters) > 0 {
		for _, msg := range m.Clusters {
			dAtA[i] = 0x22
			i++
			i = encodeVarintIrisConf(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Iris) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Iris) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if m.LogLevel != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(m.LogLevel))
	}
	if m.Parallelism != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(m.Parallelism))
	}
	if m.DefaultPodPort != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(m.DefaultPodPort))
	}
	if m.StateRefreshPeriod != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.StateRefreshPeriod)))
		n2, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.StateRefreshPeriod, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.StateBufferSize != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(m.StateBufferSize))
	}
	return i, nil
}

func (m *Cluster) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cluster) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ServiceName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(len(m.ServiceName)))
		i += copy(dAtA[i:], m.ServiceName)
	}
	if len(m.PortName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(len(m.PortName)))
		i += copy(dAtA[i:], m.PortName)
	}
	if m.Config != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintIrisConf(dAtA, i, uint64(m.Config.Size()))
		n3, err := m.Config.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeFixed64IrisConf(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32IrisConf(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintIrisConf(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Config) Size() (n int) {
	var l int
	_ = l
	if m.Iris != nil {
		l = m.Iris.Size()
		n += 1 + l + sovIrisConf(uint64(l))
	}
	if len(m.Listeners) > 0 {
		for _, e := range m.Listeners {
			l = e.Size()
			n += 1 + l + sovIrisConf(uint64(l))
		}
	}
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.Size()
			n += 1 + l + sovIrisConf(uint64(l))
		}
	}
	if len(m.Clusters) > 0 {
		for _, e := range m.Clusters {
			l = e.Size()
			n += 1 + l + sovIrisConf(uint64(l))
		}
	}
	return n
}

func (m *Iris) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovIrisConf(uint64(l))
	}
	if m.LogLevel != 0 {
		n += 1 + sovIrisConf(uint64(m.LogLevel))
	}
	if m.Parallelism != 0 {
		n += 1 + sovIrisConf(uint64(m.Parallelism))
	}
	if m.DefaultPodPort != 0 {
		n += 1 + sovIrisConf(uint64(m.DefaultPodPort))
	}
	if m.StateRefreshPeriod != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.StateRefreshPeriod)
		n += 1 + l + sovIrisConf(uint64(l))
	}
	if m.StateBufferSize != 0 {
		n += 1 + sovIrisConf(uint64(m.StateBufferSize))
	}
	return n
}

func (m *Cluster) Size() (n int) {
	var l int
	_ = l
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovIrisConf(uint64(l))
	}
	l = len(m.PortName)
	if l > 0 {
		n += 1 + l + sovIrisConf(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovIrisConf(uint64(l))
	}
	return n
}

func sovIrisConf(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIrisConf(x uint64) (n int) {
	return sovIrisConf(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIrisConf
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Iris", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIrisConf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Iris == nil {
				m.Iris = &Iris{}
			}
			if err := m.Iris.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Listeners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIrisConf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Listeners = append(m.Listeners, &envoy_api_v22.Listener{})
			if err := m.Listeners[len(m.Listeners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIrisConf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Routes = append(m.Routes, &envoy_api_v23.RouteConfiguration{})
			if err := m.Routes[len(m.Routes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clusters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIrisConf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clusters = append(m.Clusters, &Cluster{})
			if err := m.Clusters[len(m.Clusters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIrisConf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIrisConf
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Iris) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIrisConf
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Iris: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Iris: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIrisConf
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogLevel", wireType)
			}
			m.LogLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogLevel |= (LogLevel(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parallelism", wireType)
			}
			m.Parallelism = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Parallelism |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultPodPort", wireType)
			}
			m.DefaultPodPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultPodPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRefreshPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIrisConf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StateRefreshPeriod == nil {
				m.StateRefreshPeriod = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.StateRefreshPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateBufferSize", wireType)
			}
			m.StateBufferSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StateBufferSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIrisConf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIrisConf
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Cluster) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIrisConf
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cluster: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cluster: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIrisConf
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIrisConf
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIrisConf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &envoy_api_v21.Cluster{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIrisConf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIrisConf
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIrisConf(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIrisConf
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowIrisConf
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthIrisConf
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIrisConf
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipIrisConf(dAtA[start:])
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
	ErrInvalidLengthIrisConf = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIrisConf   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pkg/v1pb/iris_conf.proto", fileDescriptorIrisConf) }

var fileDescriptorIrisConf = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x9c, 0xb8, 0xf9, 0x92, 0x09, 0x2d, 0xee, 0xd0, 0x16, 0x53, 0x50, 0x30, 0xdd, 0x60,
	0x55, 0x62, 0xac, 0x1a, 0x16, 0x6c, 0x9b, 0x36, 0x54, 0x11, 0x56, 0x1a, 0x06, 0x51, 0x24, 0x36,
	0x96, 0x7f, 0xc6, 0x66, 0x84, 0xe3, 0x31, 0x33, 0xb6, 0x05, 0x7d, 0x12, 0x5e, 0x80, 0x77, 0x61,
	0xc9, 0x86, 0x35, 0x28, 0x4f, 0x82, 0x66, 0xec, 0xb4, 0xa4, 0x82, 0xdd, 0xd5, 0x39, 0xe7, 0xfa,
	0xde, 0x7b, 0xce, 0x18, 0x98, 0xc5, 0x87, 0xd4, 0xa9, 0x8f, 0x8a, 0xd0, 0xa1, 0x9c, 0x0a, 0x3f,
	0x62, 0x79, 0x82, 0x0a, 0xce, 0x4a, 0x06, 0x37, 0x3f, 0x56, 0x21, 0x2d, 0x91, 0x84, 0x51, 0x7d,
	0xb4, 0x3f, 0x4a, 0x19, 0x4b, 0x33, 0xe2, 0x28, 0x32, 0xac, 0x12, 0x27, 0xae, 0x78, 0x50, 0x52,
	0x96, 0x37, 0xf2, 0xfd, 0x3d, 0x92, 0xd7, 0xec, 0xb3, 0x13, 0x14, 0xd4, 0xa9, 0x5d, 0x27, 0x8a,
	0xc5, 0x5f, 0xf1, 0xec, 0x1f, 0x38, 0xbf, 0xc2, 0x77, 0x52, 0x96, 0x32, 0x55, 0x3a, 0xb2, 0x6a,
	0xd0, 0x83, 0x1f, 0x1a, 0xe8, 0x9d, 0xb0, 0x3c, 0xa1, 0x29, 0x7c, 0x0c, 0x74, 0xb9, 0x93, 0xa9,
	0x59, 0x9a, 0x3d, 0x74, 0xef, 0xa0, 0xb5, 0x35, 0xd1, 0x94, 0x53, 0x81, 0x95, 0x00, 0x3e, 0x03,
	0x83, 0x8c, 0x8a, 0x92, 0xe4, 0x84, 0x0b, 0xb3, 0x63, 0x75, 0xed, 0xa1, 0xbb, 0x87, 0xd4, 0x54,
	0x14, 0x14, 0x14, 0xd5, 0x2e, 0xf2, 0x5a, 0x1a, 0x5f, 0x0b, 0xe1, 0x73, 0xd0, 0xe3, 0xac, 0x2a,
	0x89, 0x30, 0xbb, 0xaa, 0xc5, 0x5a, 0x6f, 0xc1, 0x92, 0x6b, 0x36, 0x69, 0xef, 0xc7, 0xad, 0x1e,
	0xba, 0xa0, 0x1f, 0x65, 0x95, 0x28, 0xe5, 0x38, 0xbd, 0x1d, 0xb7, 0xbe, 0xdc, 0x49, 0x43, 0xe3,
	0x2b, 0xdd, 0xc1, 0xd7, 0x0e, 0xd0, 0xe5, 0xca, 0xf0, 0x01, 0x18, 0xe4, 0xc1, 0x82, 0x88, 0x22,
	0x88, 0x88, 0x3a, 0x6d, 0x80, 0xaf, 0x01, 0x75, 0x0a, 0x4b, 0xfd, 0x8c, 0xd4, 0x24, 0x33, 0x3b,
	0x96, 0x66, 0x6f, 0xb9, 0x77, 0x6f, 0x7c, 0xdb, 0x63, 0xa9, 0x27, 0x69, 0xdc, 0xcf, 0xda, 0x0a,
	0x5a, 0x60, 0x58, 0x04, 0x3c, 0xc8, 0x32, 0x92, 0x51, 0xb1, 0x30, 0xbb, 0x96, 0x66, 0x6f, 0xe2,
	0x3f, 0x21, 0x68, 0x03, 0x23, 0x26, 0x49, 0x50, 0x65, 0xa5, 0x5f, 0xb0, 0xd8, 0x2f, 0x18, 0x2f,
	0x4d, 0x5d, 0xc9, 0xb6, 0x5a, 0x7c, 0xce, 0xe2, 0x39, 0xe3, 0x25, 0x7c, 0x05, 0x76, 0x44, 0x19,
	0x94, 0xc4, 0xe7, 0x24, 0xe1, 0x44, 0xbc, 0xf7, 0x0b, 0xc2, 0x29, 0x8b, 0xcd, 0x0d, 0x95, 0xc2,
	0x3d, 0xd4, 0xbc, 0x0e, 0xb4, 0x7a, 0x1d, 0xe8, 0xb4, 0x75, 0x67, 0xac, 0x7f, 0xf9, 0xf9, 0x50,
	0xc3, 0x50, 0x35, 0xe3, 0xa6, 0x77, 0xae, 0x5a, 0xe1, 0x21, 0xd8, 0x6e, 0x3e, 0x19, 0x56, 0x49,
	0x42, 0xb8, 0x2f, 0xe8, 0x25, 0x31, 0x7b, 0x6a, 0xfa, 0x6d, 0x45, 0x8c, 0x15, 0xfe, 0x9a, 0x5e,
	0x92, 0x83, 0x4f, 0xe0, 0xff, 0xd6, 0x3c, 0xf8, 0x08, 0xdc, 0x12, 0x84, 0xd7, 0x34, 0x22, 0xbe,
	0x34, 0xa8, 0x35, 0x6b, 0xd8, 0x62, 0xb3, 0x60, 0x41, 0xe0, 0x7d, 0x30, 0x90, 0xa7, 0x34, 0x7c,
	0x47, 0xf1, 0x7d, 0x09, 0x28, 0xf2, 0x09, 0xe8, 0x45, 0x2a, 0x3f, 0x65, 0xc8, 0xd0, 0xdd, 0x5d,
	0x0f, 0x78, 0x95, 0x51, 0x2b, 0x3a, 0x7c, 0x09, 0xfa, 0x2b, 0x6b, 0xe1, 0x2e, 0xd8, 0xf6, 0xce,
	0xcf, 0x7c, 0x6f, 0x72, 0x31, 0xf1, 0xfc, 0xe9, 0xec, 0xe2, 0xd8, 0x9b, 0x9e, 0x1a, 0xff, 0xc1,
	0x01, 0xd8, 0x38, 0x9d, 0x8c, 0xdf, 0x9c, 0x19, 0x1a, 0xec, 0x03, 0x7d, 0x3a, 0x7b, 0x71, 0x6e,
	0x74, 0x64, 0xf5, 0xf6, 0x18, 0xcf, 0x8c, 0xae, 0xa4, 0x27, 0x18, 0x9f, 0x63, 0x43, 0x1f, 0xdb,
	0xdf, 0x96, 0x23, 0xed, 0xfb, 0x72, 0xa4, 0xfd, 0x5a, 0x8e, 0x34, 0xb0, 0x13, 0xb1, 0xc5, 0x8d,
	0x14, 0x8b, 0xf0, 0x9d, 0x2e, 0xff, 0xc6, 0xb0, 0xa7, 0x9c, 0x7c, 0xfa, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0xe8, 0xe2, 0x58, 0xa0, 0x03, 0x00, 0x00,
}
