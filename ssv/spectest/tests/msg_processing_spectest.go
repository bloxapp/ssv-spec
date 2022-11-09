package tests

import (
	"encoding/hex"
	"github.com/bloxapp/ssv-spec/ssv"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/stretchr/testify/require"
	"testing"
)

type MsgProcessingSpecTest struct {
	Name                    string
	Runner                  ssv.Runner
	Duty                    *types.Duty
	Messages                []*types.Message
	PostDutyRunnerStateRoot string
	// OutputMessages compares pre/ post signed partial sigs to output. We exclude consensus msgs as it's tested in consensus
	OutputMessages []*ssv.SignedPartialSignatures
	DontStartDuty  bool // if set to true will not start a duty for the runner
	ExpectedError  string
}

func (test *MsgProcessingSpecTest) TestName() string {
	return test.Name
}

func (test *MsgProcessingSpecTest) Run(t *testing.T) {
	v := testingutils.BaseValidator(testingutils.KeySetForShare(test.Runner.GetBaseRunner().Share))
	v.DutyRunners[test.Runner.GetBaseRunner().BeaconRoleType] = test.Runner
	v.Network = test.Runner.GetNetwork()

	var lastErr error
	if !test.DontStartDuty {
		lastErr = v.StartDuty(test.Duty)
	}
	for _, msg := range test.Messages {
		err := v.ProcessMessage(msg)
		if err != nil {
			lastErr = err
		}
	}

	if len(test.ExpectedError) != 0 {
		require.EqualError(t, lastErr, test.ExpectedError)
	} else {
		require.NoError(t, lastErr)
	}

	// test output message
	broadcastedMsgs := v.Network.(*testingutils.TestingNetwork).BroadcastedMsgs
	if len(broadcastedMsgs) > 0 {
		index := 0
		for _, msg := range broadcastedMsgs {
			msgType := msg.GetID().GetMsgType()
			if msgType != types.PartialRandaoSignatureMsgType &&
				msgType != types.PartialContributionProofSignatureMsgType &&
				msgType != types.PartialSelectionProofSignatureMsgType &&
				msgType != types.PartialPostConsensusSignatureMsgType {
				continue
			}

			msg1 := &ssv.SignedPartialSignatures{}
			require.NoError(t, msg1.Decode(msg.Data))
			msg2 := test.OutputMessages[index]
			require.Len(t, msg1.PartialSignatures, len(msg2.PartialSignatures))

			// messages are not guaranteed to be in order so we map them and then test all roots to be equal
			roots := make(map[string]string)
			for i, partialSigMsg2 := range msg2.PartialSignatures {
				r2, err := partialSigMsg2.GetRoot()
				require.NoError(t, err)
				if _, found := roots[hex.EncodeToString(r2)]; !found {
					roots[hex.EncodeToString(r2)] = ""
				} else {
					roots[hex.EncodeToString(r2)] = hex.EncodeToString(r2)
				}

				partialSigMsg1 := msg1.PartialSignatures[i]
				r1, err := partialSigMsg1.GetRoot()
				require.NoError(t, err)

				if _, found := roots[hex.EncodeToString(r1)]; !found {
					roots[hex.EncodeToString(r1)] = ""
				} else {
					roots[hex.EncodeToString(r1)] = hex.EncodeToString(r1)
				}
			}
			for k, v := range roots {
				require.EqualValues(t, k, v, "missing output msg")
			}

			index++
		}

		require.Len(t, test.OutputMessages, index)
	}

	// post root
	postRoot, err := test.Runner.GetRoot()
	require.NoError(t, err)
	require.EqualValues(t, test.PostDutyRunnerStateRoot, hex.EncodeToString(postRoot))
}
