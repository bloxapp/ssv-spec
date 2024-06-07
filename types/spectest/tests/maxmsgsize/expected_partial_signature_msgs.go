package maxmsgsize

import (
	"github.com/ssvlabs/ssv-spec/types"
)

const (
	expectedSizePartialSignatureMessage  = 144
	expectedSizePartialSignatureMessages = 164
)

func expectedPartialSignatureMessage() *types.PartialSignatureMessage {

	signature := [96]byte{1}

	return &types.PartialSignatureMessage{
		PartialSignature: signature[:],
		SigningRoot:      [32]byte{1},
		Signer:           1,
		ValidatorIndex:   1,
	}
}

func expectedPartialSignatureMessages(numSignatures int) *types.PartialSignatureMessages {

	msgs := make([]*types.PartialSignatureMessage, 0)
	for i := 0; i < numSignatures; i++ {
		msgs = append(msgs, maxPartialSignatureMessage())
	}

	return &types.PartialSignatureMessages{
		Type:     types.RandaoPartialSig,
		Slot:     1,
		Messages: msgs,
	}
}

func ExpectedPartialSignatureMessage() *MaxMessageTest {
	return &MaxMessageTest{
		Name:                  "expected PartialSignatureMessage",
		Object:                expectedPartialSignatureMessage(),
		ExpectedEncodedLength: expectedSizePartialSignatureMessage,
		IsMaxSizeForType:      false,
	}
}

func ExpectedPartialSignatureMessages() *MaxMessageTest {
	return &MaxMessageTest{
		Name:                  "expected PartialSignatureMessages",
		Object:                expectedPartialSignatureMessages(1),
		ExpectedEncodedLength: expectedSizePartialSignatureMessages,
		IsMaxSizeForType:      false,
	}
}
