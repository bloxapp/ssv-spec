package decided

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/herumi/bls-eth-go-binary/bls"
)

// CurrentInstanceFutureRound tests a decided msg received for current running instance for a future round
func CurrentInstanceFutureRound() *tests.ControllerSpecTest {
	identifier := types.NewBaseMsgID(testingutils.TestingValidatorPubKey[:], types.BNRoleAttester)
	ks := testingutils.Testing4SharesSet()
	inputData := &qbft.Data{
		Root:   testingutils.TestAttesterConsensusDataRoot,
		Source: testingutils.TestAttesterConsensusDataByts,
	}
	proposeMsgEncoded, _ := testingutils.SignQBFTMsg(ks.Shares[1], types.OperatorID(1), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  qbft.FirstRound,
	}, inputData).Encode()
	signedMsgEncoded, _ := testingutils.SignQBFTMsg(ks.Shares[1], types.OperatorID(1), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  qbft.FirstRound,
	}, &qbft.Data{Root: inputData.Root}).Encode()
	signedMsgEncoded2, _ := testingutils.SignQBFTMsg(ks.Shares[2], types.OperatorID(2), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  qbft.FirstRound,
	}, &qbft.Data{Root: inputData.Root}).Encode()
	signedMsgEncoded3, _ := testingutils.SignQBFTMsg(ks.Shares[3], types.OperatorID(3), &qbft.Message{
		Height: qbft.FirstHeight,
		Round:  qbft.FirstRound,
	}, &qbft.Data{Root: inputData.Root}).Encode()
	multiSignMsg := testingutils.MultiSignQBFTMsg(
		[]*bls.SecretKey{ks.Shares[1], ks.Shares[2], ks.Shares[3]},
		[]types.OperatorID{1, 2, 3},
		&qbft.Message{
			Height: qbft.FirstHeight,
			Round:  50,
		}, inputData)
	multiSignMsgEncoded, _ := multiSignMsg.Encode()
	return &tests.ControllerSpecTest{
		Name: "decide current instance future round",
		RunInstanceData: []*tests.RunInstanceData{
			{
				InputValue: inputData,
				InputMessages: []*types.Message{
					{
						ID:   types.PopulateMsgType(identifier, types.ConsensusProposeMsgType),
						Data: proposeMsgEncoded,
					},
					{
						ID:   types.PopulateMsgType(identifier, types.ConsensusPrepareMsgType),
						Data: signedMsgEncoded,
					},
					{
						ID:   types.PopulateMsgType(identifier, types.ConsensusPrepareMsgType),
						Data: signedMsgEncoded2,
					},
					{
						ID:   types.PopulateMsgType(identifier, types.ConsensusPrepareMsgType),
						Data: signedMsgEncoded3,
					},
					{
						ID:   types.PopulateMsgType(identifier, types.ConsensusCommitMsgType),
						Data: signedMsgEncoded,
					},
					{
						ID:   types.PopulateMsgType(identifier, types.ConsensusCommitMsgType),
						Data: signedMsgEncoded2,
					},
					{
						ID:   types.PopulateMsgType(identifier, types.DecidedMsgType),
						Data: multiSignMsgEncoded,
					},
				},
				SavedDecided:       multiSignMsg,
				DecidedVal:         inputData.Source,
				DecidedCnt:         1,
				ControllerPostRoot: "f1e11de2125020fcdf4b2881272978525e6a6ff296e6827abcb2cc6354dbf034",
			},
		},
	}
}
