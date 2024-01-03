package validator

import (
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

func InvalidID() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	msgs := []*types.SSVMessage{
		{
			MsgType: types.SSVPartialSignatureMsgType,
			MsgID:   types.NewMsgID(testingutils.TestingSSVDomainType, testingutils.TestingValidatorPubKey[:], types.BeaconRole(100)),
			Data:    []byte{},
		},
	}

	return &ValidatorTest{
		Name: "invalid id",

		KeySet:                 ks,
		Messages:               msgs,
		OutputMessages:         []*types.SSVMessage{},
		BeaconBroadcastedRoots: []string{},
		ExpectedError:          "could not get duty runner for msg ID",
	}
}