package ProtoFlatBench

import (
	flatbuffers "github.com/google/flatbuffers/go"
	protocolBuffers "github.com/golang/protobuf/proto"
	protoBase "bahadir.rocks/ProtoFlatBench/schemas/proto"
	flatBase "bahadir.rocks/ProtoFlatBench/schemas/flat"
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/dustin/gojson"
	"github.com/dustin/go-humanize"
)

const (
	NAME = "Loremipsum Dolor Sitamet"
)
var BYTES_SAMPLE = []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")

var flatResult, protoResult, jsonResult []byte

/***************************************************************
** FlatBuffers
****************************************************************/
func FlatWrite() []byte{
	builder := flatbuffers.NewBuilder(0)

	name := builder.CreateString(NAME)
	profileImage := builder.CreateByteVector(BYTES_SAMPLE)

	flatBase.ProfileStart(builder)
	flatBase.ProfileAddId(builder, uint64(time.Now().Unix()))
	flatBase.ProfileAddName(builder, name)
	flatBase.ProfileAddGender(builder, flatBase.GenderMale)
	flatBase.ProfileAddProfileImage(builder, profileImage)

	p := flatBase.ProfileEnd(builder)
	builder.Finish(p)

	buf := builder.FinishedBytes()
	//println(buf)
	return buf
}

func TestFlatWrite(t *testing.T) {
	buf := FlatWrite()
	t.Log("Flatbuffers size: ", humanize.Bytes(uint64(len(buf))))
}

func BenchmarkFlatWrite(b *testing.B) {
	var buf []byte
	for n:= 0; n < b.N; n++ {
		buf = FlatWrite()
	}
	flatResult = buf
}


func BenchmarkFlatRead(b *testing.B) {
	buff := FlatWrite()
	b.ResetTimer()

	var profile *flatBase.Profile
	for n:= 0; n < b.N; n++ {
		profile = flatBase.GetRootAsProfile(buff, 0)
	}
	_ = profile
}


/***************************************************************
** ProtocolBuffers
****************************************************************/
func ProtoWrite() (buff []byte, err error) {
	stamp := time.Now().Unix()
	name := NAME
	gender := protoBase.Profile_Male
	p := protoBase.Profile{
		Id: &stamp,
		Name: &name,
		Gender: &gender,
		ProfileImage: BYTES_SAMPLE,
	}

	buff, err = protocolBuffers.Marshal(&p)
	return
}

func TestProtoWrite(t *testing.T) {
	buf, err := ProtoWrite()
	assert.Nil(t, err)
	t.Log("ProtoBuf size: ", humanize.Bytes(uint64(len(buf))))
}

func BenchmarkProtoWrite(b *testing.B) {
	b.ReportAllocs()
	var buf []byte
	for n:= 0; n < b.N; n++ {
		buf, _ = ProtoWrite()
	}
	protoResult = buf
}

func BenchmarkProtoRead(b *testing.B) {
	buff, err := ProtoWrite()
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	var d protoBase.Profile
	for n:= 0; n < b.N; n++ {
		protocolBuffers.Unmarshal(buff, &d)
	}
}

/***************************************************************
** JSON
****************************************************************/

func JSONWrite() (buff []byte, err error) {
	stamp := time.Now().Unix()
	name := NAME
	gender := protoBase.Profile_Male
	p := protoBase.Profile{
		Id: &stamp,
		Name: &name,
		Gender: &gender,
		ProfileImage: BYTES_SAMPLE,
	}
	buff, err = json.Marshal(p)
	return
}

func TestJSONWrite(t *testing.T) {
	buf, err := JSONWrite()
	assert.Nil(t, err)
	t.Log("JSON size: ", humanize.Bytes(uint64(len(buf))))
}

func BenchmarkJSONWrite(b *testing.B) {
	var buf []byte
	for n:= 0; n < b.N; n++ {
		buf, _ = JSONWrite()
	}
	jsonResult = buf

}

func BenchmarkJSONRead(b *testing.B) {
	buff, err := JSONWrite()
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	var d interface{}
	for n:= 0; n < b.N; n++ {
		json.Unmarshal(buff, &d)
	}
}