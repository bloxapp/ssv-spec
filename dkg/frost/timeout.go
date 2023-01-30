package frost

import (
	"github.com/bloxapp/ssv-spec/dkg"
)

func (fr *Instance) UponRoundTimeout() error {
	if fr.state.currentRound != fr.state.roundTImer.Round() {
		return nil
	}
	prevRound := fr.state.currentRound
	fr.state.currentRound = Timeout

	msg := &ProtocolMsg{
		Round: Timeout,
		TimeoutMessage: &TimeoutMessage{
			Round: prevRound,
		},
	}
	bcastMsg, err := fr.saveSignedMsg(msg)
	if err != nil {
		return err
	}
	return fr.config.GetNetwork().BroadcastDKGMessage(bcastMsg)
}

func (fr *Instance) ProcessTimeoutMessage() (finished bool, protocolOutcome *dkg.ProtocolOutcome, err error) {
	return true, &dkg.ProtocolOutcome{}, nil
}
