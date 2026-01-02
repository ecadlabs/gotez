package protocol_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type protoTestData struct {
	proto  *tz.ProtocolHash
	blocks []string
}

var testData = []protoTestData{
	{
		proto: &core.Proto024PtTALLiN,
		blocks: []string{
			"295520",
			"449131",
			"880107",
			"928754",
			"928755",
			"928792",
			"930637",
			"930989",
			"931138",
			"931152",
			"966982",
		},
	},
	{
		proto: &core.Proto023PtSeouLo,
		blocks: []string{
			"217650",
			"217651",
			"217652",
			"217653",
			"217654",
			"217655",
			"217656",
			"217657",
			"217658",
			"217659",
		},
	},
	{
		proto: &core.Proto022PsRiotum,
		blocks: []string{
			"457585",
			"584373",
			"586704",
		},
	},
	{
		proto: &core.Proto021PsQuebec,
		blocks: []string{
			"1196001",
			"1196033",
			"1196086",
			"1197064",
			"1197095",
			"1197101",
			"1197107",
			"1197158",
			"1197165",
			"1197167",
			"1197183",
			"1233892",
			"1234288",
			"1234330",
			"775785",
		},
	},
	{
		proto: &core.Proto019PtParisB,
		blocks: []string{
			"83409",
			"83410",
			"83411",
			"83412",
			"83413",
			"83414",
			"83415",
			"83416",
			"83417",
			"83418",
		},
	},
	{
		proto: &core.Proto018Proxford,
		blocks: []string{
			"494220",
			"494221",
			"494222",
			"494223",
		},
	},
	{
		proto: &core.Proto017PtNairob,
		blocks: []string{
			"10404",
			"10413",
			"577841",
			"621690",
			"649626",
			"650438",
			"650456",
			"7921",
		},
	},
	{
		proto: &core.Proto016PtMumbai,
		blocks: []string{
			"3279466",
			"3549429",
			"3596415",
			"3615666",
			"3514591",
			"181313",
			"298135",
			"298154",
			"327682",
			"332053",
			"332064",
			"332066",
			"332075",
			"332090",
			"332091",
			"332093",
			"332470",
			"332530",
			"332534",
			"39524",
			"41157",
			"41821",
		},
	},
	{
		proto: &core.Proto015PtLimaPt,
		blocks: []string{
			"2981889",
			"2981890",
			"2981891",
			"2981892",
		},
	},
	{
		proto: &core.Proto014PtKathma,
		blocks: []string{
			"2736129",
			"2736130",
			"2736131",
			"2736132",
		},
	},
	{
		proto: &core.Proto013PtJakart,
		blocks: []string{
			"2490369",
			"2490370",
			"2490371",
			"2490372",
			"2490373",
			"2490374",
			"2490375",
			"2490376",
		},
	},
	{
		proto: &core.Proto012Psithaca,
		blocks: []string{
			"2244609",
			"2244610",
			"2244611",
			"2244612",
			"2244613",
			"2244614",
			"2244615",
			"2244616",
			//"tlnt_2173954",
		},
	},
}

func TestBlock(t *testing.T) {
	for _, protoData := range testData {
		t.Run(protoData.proto.String(), func(t *testing.T) {
			for _, block := range protoData.blocks {
				t.Run(block, func(t *testing.T) {
					fileName := filepath.Join("test_data", core.ProtocolShortName(protoData.proto), block+".bin")
					buf, err := os.ReadFile(fileName)
					require.NoError(t, err)
					out, err := protocol.NewBlockInfo(protoData.proto)
					require.NoError(t, err)
					_, err = encoding.Decode(buf, out, encoding.Dynamic())
					if !assert.NoError(t, err) {
						//c := spew.NewDefaultConfig()
						//c.DisableMethods = true
						//c.Dump(out)
						//j, _ := json.MarshalIndent(out, "", "    ")
						//fmt.Println(string(j))
						if err, ok := err.(*encoding.Error); ok {
							fmt.Println(err.Path)
						}
						return
					}

					// check operation group signature lengths
					for _, list := range out.GetOperations() {
						for _, grp := range list {
							_, err := grp.GetContents().GetSignature()
							assert.NoError(t, err)
						}
					}
				})
			}
		})
	}
}

var headerTestData = []protoTestData{
	{
		proto: &core.Proto016PtMumbai,
		blocks: []string{
			"3279466",
		},
	},
}

func TestBlockHeader(t *testing.T) {
	for _, protoData := range headerTestData {
		t.Run(protoData.proto.String(), func(t *testing.T) {
			for _, block := range protoData.blocks {
				t.Run(block, func(t *testing.T) {
					fileName := filepath.Join("test_data", core.ProtocolShortName(protoData.proto), "header_"+block+".bin")
					buf, err := os.ReadFile(fileName)
					require.NoError(t, err)
					out, err := protocol.NewBlockHeaderInfo(protoData.proto)
					require.NoError(t, err)
					_, err = encoding.Decode(buf, out, encoding.Dynamic())
					if !assert.NoError(t, err) {
						if err, ok := err.(*encoding.Error); ok {
							fmt.Println(err.Path)
						}
					}
				})
			}
		})
	}
}
