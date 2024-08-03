package helpmeunpack

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"

	"github.com/Sc4ramouche/hackattic/pkg/tools"
)

type solution struct {
	Int               int32   `json:"int"`
	Uint              uint32  `json:"uint"`
	Short             int16   `json:"short"`
	Float             float64 `json:"float"`
	Double            float64 `json:"double"`
	Big_endian_double float64 `json:"big_endian_double"`
	Playground        int32   `json:"playground"`
}

func HelpMeUnpack() {
	respBytes := tools.GetTask("https://hackattic.com/challenges/help_me_unpack/problem?access_token=88b055aabe85d853")
	var bytes map[string]interface{}
	json.Unmarshal(respBytes, &bytes)

	base64Str := bytes["bytes"].(string)
	unpackedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", unpackedBytes)

	var integer int32
	integer = int32(binary.LittleEndian.Uint32(unpackedBytes[:4]))
	fmt.Println("int:", integer)

	unsignedInteger := binary.LittleEndian.Uint32(unpackedBytes[4:8])
	fmt.Println("uint:", unsignedInteger)

	short := int16(binary.LittleEndian.Uint16(unpackedBytes[8:12]))
	fmt.Println("short:", short)

	unsignedIntForFloat := binary.LittleEndian.Uint32(unpackedBytes[12:16])
	unsignedFloat32 := float64(math.Float32frombits(unsignedIntForFloat))
	fmt.Println("float:", unsignedFloat32)

	unsignedIntForDouble := binary.LittleEndian.Uint64(unpackedBytes[16:24])
	double := math.Float64frombits(unsignedIntForDouble)
	fmt.Println("double:", double)

	unsignedIntForBigIndianDouble := binary.BigEndian.Uint64(unpackedBytes[24:32])
	bigIndianDouble := math.Float64frombits(unsignedIntForBigIndianDouble)
	fmt.Println("big indian double:", bigIndianDouble)

	solution := solution{Int: integer, Uint: unsignedInteger, Short: short, Float: unsignedFloat32, Double: double, Big_endian_double: bigIndianDouble, Playground: 1}
	solutionJson, err := json.Marshal(&solution)
	if err != nil {
		panic(err)
	}

	tools.PostTask("https://hackattic.com/challenges/help_me_unpack/solve?access_token=88b055aabe85d853&playground=1", solutionJson)
}
