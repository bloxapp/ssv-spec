package messages

import (
	"encoding/hex"

	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// UnmarshalJustifications tests unmarshalling justifications
func UnmarshalJustifications() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	encodedRCMsg, _ := hex.DecodeString("b77c035c1d1d9c6c7cc22810c50a85c9e560ad791509055061dfe07d403edaa1304161466d49f64844b6e5cf4b09709f069debfd91438f97d414a4f64cdcb4f8cf1e9703f8da141be9f74509087e2f84492314b0341966c4b8fc16d931f6ba086c00000074000000c400000001000000000000000300000000000000000000000000000002000000000000004c000000be956fb7df4ef37531682d588320084fc914c3f0fed335263e5b44062e6c29b40000000000000000500000005000000001020304")
	encodedPrepareMsg, _ := hex.DecodeString("8129e6862a5120bd085e1936b4efb5a55fc7d19c0d0fda0e9ec576d18abd4a17ab3a033f5296b74c5fdaf85cb7b3da3201b63feca76b883613e3b1ca137e763a342e3b1dddbce016f8ca3cbce32c8b125dd8c25a7639819c20b539e9e7c6c5796c00000074000000c400000001000000000000000100000000000000000000000000000001000000000000004c000000be956fb7df4ef37531682d588320084fc914c3f0fed335263e5b44062e6c29b40000000000000000500000005000000001020304")

	msg := testingutils.TestingProposalMessageWithParams(
		ks.Shares[1], types.OperatorID(1), 2, qbft.FirstHeight, testingutils.TestingQBFTRootData,
		[][]byte{encodedRCMsg}, [][]byte{encodedPrepareMsg})

	r, err := msg.GetRoot()
	if err != nil {
		panic(err)
	}

	b, err := msg.Encode()
	if err != nil {
		panic(err)
	}

	return &tests.MsgSpecTest{
		Name: "unmarshal justifications",
		Messages: []*qbft.SignedMessage{
			msg,
		},
		EncodedMessages: [][]byte{
			b,
		},
		ExpectedRoots: [][32]byte{
			r,
		},
	}
}
