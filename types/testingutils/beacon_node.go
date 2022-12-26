package testingutils

import (
	"encoding/hex"
	v1 "github.com/attestantio/go-eth2-client/api/v1"
	altair "github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/types"
	ssz "github.com/ferranbt/fastssz"
	"github.com/goccy/go-yaml"
	"github.com/prysmaticlabs/go-bitfield"
)

var signBeaconObject = func(
	obj ssz.HashRoot,
	domainType spec.DomainType,
	ks *TestKeySet,
) spec.BLSSignature {
	domain, _ := NewTestingBeaconNode().DomainData(1, domainType)
	ret, _, _ := NewTestingKeyManager().SignBeaconObject(obj, domain, ks.ValidatorPK.Serialize())

	blsSig := spec.BLSSignature{}
	copy(blsSig[:], ret)

	return blsSig
}

var TestingAttestationData = &spec.AttestationData{
	Slot:            12,
	Index:           3,
	BeaconBlockRoot: spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	Source: &spec.Checkpoint{
		Epoch: 0,
		Root:  spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	},
	Target: &spec.Checkpoint{
		Epoch: 1,
		Root:  spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	},
}
var TestingWrongAttestationData = func() *spec.AttestationData {
	byts, _ := TestingAttestationData.MarshalSSZ()
	ret := &spec.AttestationData{}
	if err := ret.UnmarshalSSZ(byts); err != nil {
		panic(err.Error())
	}
	ret.Slot = 100
	return ret
}()

var TestingSignedAttestation = func(ks *TestKeySet) *spec.Attestation {
	aggregationBitfield := bitfield.NewBitlist(TestingAttesterDuty.CommitteeLength)
	aggregationBitfield.SetBitAt(TestingAttesterDuty.ValidatorCommitteeIndex, true)
	return &spec.Attestation{
		Data:            TestingAttestationData,
		Signature:       signBeaconObject(TestingAttestationData, types.DomainAttester, ks),
		AggregationBits: aggregationBitfield,
	}
}

var Transactions = func() []bellatrix.Transaction {
	// taken from https://github.com/attestantio/go-eth2-client/blob/1a94a0a534b9d75087403954c33f2a5819a896a9/api/v1/bellatrix/blindedbeaconblock_test.go#L149
	data := []byte(`{parent_hash: '0x17f4eeae822cc81533016678413443b95e34517e67f12b4a3a92ff6b66f972ef', fee_recipient: '0x58e809c71e4885cb7b3f1d5c793ab04ed239d779', state_root: '0x3d6e230e6eceb8f3db582777b1500b8b31b9d268339e7b32bba8d6f1311b211d', receipts_root: '0xea760203509bdde017a506b12c825976d12b04db7bce9eca9e1ed007056a3f36', logs_bloom: '0x0c803a8d3c6642adee3185bd914c599317d96487831dabda82461f65700b2528781bdadf785664f9d8b11c4ee1139dfeb056125d2abd67e379cabc6d58f1c3ea304b97cf17fcd8a4c53f4dedeaa041acce062fc8fbc88ffc111577db4a936378749f2fd82b4bfcb880821dd5cbefee984bc1ad116096a64a44a2aac8a1791a7ad3a53d91c584ac69a8973daed6daee4432a198c9935fa0e5c2a4a6ca78b821a5b046e571a5c0961f469d40e429066755fec611afe25b560db07f989933556ce0cea4070ca47677b007b4b9857fc092625f82c84526737dc98e173e34fe6e4d0f1a400fd994298b7c2fa8187331c333c415f0499836ff0eed5c762bf570e67b44', prev_randao: '0x76ff751467270668df463600d26dba58297a986e649bac84ea856712d4779c00', block_number: 2983837628677007840, gas_limit: 6738255228996962210, gas_used: 5573520557770513197, timestamp: 1744720080366521389, extra_data: '0xc648', base_fee_per_gas: '88770397543877639215846057887940126737648744594802753726778414602657613619599', block_hash: '0x42c294e902bfc9884c1ce5fef156d4661bb8f0ff488bface37f18c3e7be64b0f', transactions: ['0x101b883470f1cb7e0a74561be59e08a6eff6aedd3408c190fd359f0bfb628d2461354b7fe4fdad4b8e72b8775cd44e339ad6b5f22a8f53c418bda2b01200a07f1fe3d010bbb96f5ca0d4919192370c15bc46ed455c797b1b11154be359638f9e487121182fae03a7d26012ed7c85b64a63aa5d56a98ac589f9950a9f5bf1a42c1eea245a98f2f4e743c5f8eac1584893104853dc1b5576826156b371a50c59bcb238d0794a185dc0816dbd8c10a0b0e1b8fbe01c4e8dd719f1e4e9b2ded8613f87f1d01e3ea28e9311d135301b2d1260e4811789e1ecbc33e573346f4da94f4c272e21e23b1d414706429c7f0b40c3de243894349c4a59ece791e5fd086897aef81fc23e4d55bb52b28174c3fa9f2c44d370fbbac1dee561f120560c3dda34a731d4618fd22d595b725efa87bb62f8f93bd7d906c4782a647e2cb14ed293e58bc793852058ab5e6f3df76b30e99102b82c8e7d005e1b675fc74b95032616c590ff08da710dd085be570cc0b2c13891625b44b5b3e1846606c39ad39bb72b', '0x4b5edff58cf95969ead0470ab897da6c9d69b517e07fd4ed8fa48c14284a5a060ecb745f66cf1be55fb4905d62d8d376865f7d2d1816844fc4719f5d79ab905474d00f62aed6692e5a93be1b32740a8083fc2a61b0e1fc13ad409410d37f3cb0a5275c966abab015047c15251cd301cd31a7b2a0502f7f953e672d61606616b16d5117163064fb33d97eb566f7fce5f01d3833343c1c97e6221f9f0798415f3a4b87fd472e53a24c1a101e1dd55c8f65c2c0f4ccc4b46a133fda49db5dba96631b4cfd1a05662e42a8e15a26d3148a70be305c85f87dae4217fb91498c4098b946a9042355968b765e2e62bb0cf26d59e534c3af8795fa0f4a44ed0d39d258acd934c3416e4d4a738eaa473526d99bee037765d5f6034c830eb766ef067a1468630fbb65b7c5a862017fe84d4d1961f90c37f18a4bd2509fe2e96cb1e26971900e20295c8a9e9ed77b348d4509a8425090318be5c9d2bcda36bdfabb71bfb36755794f78c877df2825bd736a358933af77eae6edf701ed7ef168f57f677df3445d89c5eefc783184eadd3886fcfd75f5f142bc10904a019acdf7861caa7e0fba3e7831b0a549a56c0f174e80cffb8992346ddf7ce4eeff9e335531df3dff57d5f539bc0d3eac57f70a8f973e0864c87b25c3e0bea72e05eda8a120178186dfbfaa9a00f904f23a', '0x452af990975c3bcafee7bde4738fdf32b8479be0e9e30b11b0ccbf31a6f884f4098e47759b7a3eb7d7bd0cfa2510dd654b28d664696aca987f55c7bfe73be1cc70a768cdb2594a13a763dbb991186b8b8b2e913857aadc08f940239d03a0b181ae849d557da5b54bfc966231690bb4660e083cdae28caf8ed33e3f66672772ea827253421bade013af57a290a915dbc777f2afccd9cb29e260ecc5ea54cd3a1e25cf66f2937f8061b3ba6b1ebf3129568f16e3dd04d5c50992cc348f3e615af346dbf2c144aeb19932dcfbff0221fae0ed9706b53245176630d011b9cd2d8630848ab1196cf9a3cc0d94392df4295be246e0ac24545c2715a40dcbc57aabffd0a86acec362affaf1bafc5c75b7ca28698a1ac14ea2c8def8ac1a32d3bf65b98aac7d0cb6fd93e5ff16274ad6d0eedf773694f29fc7a234deebc893e4cea4a5483d876e4f35019d6a62f1c407739b68b7a9f5408b4fb854534344fbffe3239feb17c0e7ac269b447bc6246579e1208b6904751eeacb985cbd43bb7792de0b428f1476301c479a3922f61f650c8298fb0b7584b52c7bfcf4dfe51335ab68d571c5815fc78d772346b0b13dbbb8906f076a0452e7e7a414e005dc37cfee85810eccab5999e2d43e13709bf66e8a8936a4283885b158115050789d3b4d9c8dc8026ac720069d78d47dd5e183032f9c53d4c57640fcd6207118d9738f00cd5ddf587f3a7c401d923aa2fb08dba1768728001abc3436cc2b5cf978b558ae58a0578344e7464cd00e719135c70244e2faf1264571a8999789c26f401753e429a1135f18906ebef19e122489622738724a6424cad363bed43304c1285c8da4824fec75d7c51b0a34070d8b976e8e8c8fd50908a7e440092dddd970fd55793e2a4446342bd3daf5b96220977d8f0c', '0x0877448694993717a89c57b63640612ac4b0258cb5fdda4a311650c631c35c7313b6d2a094fdb207857d94500c37ab20ea0aa54af951fb04584a37b857981c6d13922e95cfecb70b69ab9a57ee6c13ccf8aa38c52de008ec16d9090aa4bf15db2f4afcdbb1bf4920efe5a1aeeff2c949d43460d67837af87bcdffd9e972340cb40de6d87fa11d83bbfb29e97ef2509097e8dec69a1318132a5dd7d95c1c1e13cc85d37a33c9f7d52379b4a47bf889903c8f3ebd2800526d0916e1aad00e02b682e55bc2865c3ff4ce0cc6aad1bd7d8e2901ea53f3a5e2c025dbb9a00f0ce88583c1dbd3d491ed04ba8260dc06fdb8a9162e022c75e9f057da0abed537b34214df234e1f8b26e7374cfa5470272ef03f7b41f7ed067c4c6011c8a17b2e65340e36af81ecb86420755fe5a0413495e16fabee3f9e5524ab7b12a3cffe20b1df7be32434d7da3fb1f3e9b16f42a4f550501120036b193701ac9eb6f760c2da70e3175a66b10463e43c0442a56217ca7cd25fdb46f3eff28cf1bdfe1b7eb3bf9e85ad8ffde207529c9bb1094dbae4db04a4bff6571c985bd629cdc2b78f739eca6694f32fded6944202859277740267d5dc4c1f74ddd1401c6d514ce23b885723c4618789b5c2ddeffec2179be8dec1347ca4ebed5e8bb10d7d17d41c9709a978fc6189f0c3d49d1b7f41bf8f1dc112f2ece84b6c6a687f44b95e62d274f89fa07bd0d3fdd3ee2a97233e363329ef7096ae5a45b2982859e983f5d989a928e9b88579308ece2391dbde378b81a54d38b3f81225d59f8bb511ba7eb590154ceb8258b804b6da98b3af6f395f5b3f1d12f5fd3c29ef54f31ccea0a36ec2daec0a87030daba8d079093ddde17871c4aa1a7dc3dbd4d760be2152dc250ca2bba34a55daf257b9e3704c3ee244081524cb1ae7c22a0d22f1c65b88b1e534ea1cb8f75cea4a7c03d6786f85327876da72dff1d4d049b51ecc10124279a0cbc151e76ddd475cf3dfeede59a902f4c7145786b5993c8bf8016265ec36298b27d0c6a21c7484fc01a8f14c8c287d14ab86789e34699fdb57f6c43486f0fd9013f2f2c62c60b75b1dc3e4d38a6f7c06e7029f874a204b059d834ffb44c99f843ae33ed0950', '0x259310d5134d22e0ef42c3686b2adadc6cd1ae7d7836ac71a69d8ba2d02d0152c320610c12c57cba182c5d1e21198e787b21d0c522106aa8243ec994c4ea0b7959a3269d13566f3d0a3eb5ed276d9e22b33fc12e26cafde04b24ec0fd90455dc26d30a9fc25588b762681ca69aecc19e7971dd4cd063d4e31ee99f3c82015ed9a70e58b9d7cd9a8de38eba90ceffc629e7d6c06e4f2fe9cd45bae557652fe58cc9be54de9a994bf14c3bce787a106778416fbe966e95e51a35ba78d3dc4e4f7551e2791af00d50362493697d55ea6718ba2f089eda330e26100fb5adbb939afaf74982795414422712e8560cec372eaf1a56b50ba4e00e42b145537e94e88eddd7d200d0153edb6dcc12eb298666b0aef9ea495fccfda06b7affe7227bceb41d9ac9dd150df8642cdb11df5ab92b69629b6a5ccdc7ba92b6cf12172217057d291b2fdc6f104e86617be1b7fffeb36af59a018f50c055e24ef9b6daff839083bc9dfeae9101f6aff49f808f603834802d160283cd71b275bf97eb49f5612215cc8fe93875', '0xfa082adb51ff0dca75ff57ba7852d794284db5b8d498002a821e5fc2c57a27cc1fe591f12860895b0057bc75ecbff9824bb46e2af06a785b6adfa9e32f49d6235776c3bcace18b330cec1126ac2bb5f3679339037817eab536fed29fe5c62ef790cb74893f1524280b0111e24fc0172130d00a88361e63511eb56a96e552d02c4944544c193189a152844ca49cadae38b7424426a74d61763716a068ba5ca9f3bdc2e0e9b644f1f2e02596bd3f446bb9f13dfb18ad2c9a2ba97771bb994f801affe2e8dda8d638e366c5b8263cc891a8a35e52c79f0856bfeaf0719bd1bfb54acc467783b6c08e55491d39b20f218342991381c38d357a4659baf8d44ddf5045d7f116bcc55f78c28a0dc100d06dff44e639b8e7adf967816c03de54eccdf9f6402a2806889d2cd980f34e5197771f13f1fa6b8c7c732031664bb675ff12d642ae62459e92b1cc28c62349e1636850b8a3ff411e1f14b097f7fa3b23eaa173d17a13f0703eecd7358aa623057325eed381a2ba13c50a13a3d8adaad0295960d925709d6e7cb101e894d9ae8db3dcb82a45570b6e02ab68aa4ea94f8779a9c45a5599', '0x71ea3e9ec7ac8a144753ec38e78401b14b489a79266dd0527a4505adca584fe7406d9f7f05ef46d262384fbf1c0607f745a40681d4855aa34c37375b2cf7c99d46b6ae4e3aa9ee209782dce9d167bbff79686e59426a0513b951818ee02a9ea5ed8bcd0f991f46eee44c6e82fc6d07823a9f44f2fd7bfa0e250d6699cf7ad2577eb9eaf0dd9b1595cf018383c3d45956e508fa982bbb7744522a03bac5cba22f58a41801f579434126f5b866dcc6ca0454e9be1c7f2f6649254f3cc09142068f412d5d454b4e4d5b54e459719550f2df1901a18467d9d5297ed4a9eedb9f402f2bfd4bd2e50ab697ea5bef5e7e8082650b823635a4f0f55c18712d0f4d365824c79827d454425aec0b4ea6561ce3f1c5411ab4dff26e2412791cdaa28bf6c8fa53af412828c599d7876508f78f2c82ee67e8357947c6848af143fc5d20409049925cd1b194244466711ce7c34a72a165392564b96e280406a95da927d14ebbb6b999ef446d5ad49881c219dac8ad02d994e059569e91b84f211a4c39a3fbd594a8c835aa5976e0c8899d81a5abaa14301662b9e14fb96ea89862ad898cd7d77bcd2547f0f40471f86d4a4f26e4274f2e95fd9b2b604de867be95630b1ea6ce45ba79acc81b986605e46acda208bd3302ffcc6e83f91bff8362b3f9641ca0227d8b31341e320003', '0x2b2acf1e7e043b2c38ea1750a6c31f174b0a8b00137ec3eea24ed9fdf979bf04ed923a2cf5cc05a9acf697d27ccdb2a903a4846729cd9b2eead414da983d4f5bfd0ff4f2e7eb75988e58c3a635c271', '0xce0c3b31143f1d68987deaac86a1e079deb07b3d2d497de5ebe8d94486e9a7b300c691621a68b3af4fee781c7c05a931123f910054d096c2e154950d32a26c38edb70dda50a242be4d15ce60c265767b141011cd84aa585c3af798fc1eb5ea63e93e0a426c3e1468f402f7e64f20281a4bd73cc1234174f1762a9a41989f570997036c885b1fdbf9c8153a1d19afb6526a123ca6a2fe6e98c6009f8439f6b6eea881453cc58ef344338ffd04783ce9c28c2e373812a65157643f679ed99ab35f4024e6ee31877e536b03c38616fccc993365143ddaffce39c805b391674070e993c6464156043a84266860c769de218bb5f5698f4c7a9b74142535cbcc08d5b3f747cbf6a7dfbb6c7b0ccc58af4886bc441558496e9c84d80660117777f01cdb84c0a2d3b0f2d8a6eacff5a0e55bd1be6387c3793ae8f5231c59697ae914894a49b3ae13a3a124cbf6cea33b7eb575cf6cd13e6073ddf4d033a3f87366c27089afc03707ef9a44c828388733d70f09bc08de574d07cf193f0a26a08ce8253bfdb26d5306632012c6dae194783deb31b72ac35bbd57f07f1667746d85f5d5b2132b885f8cb0206949780f9406307396dc7ece7938225fc5a55c65ed3608c27f668ba9b08875979e249655d15a4c3d77b1433ab26f56d80e2ca9178d88501ee682a14d07b17635e9968da2d98aa9eb4a5ae639e42eed0c7735313d3308fd079589d5cb2dd381a9298f67325976a6dcd3e38f836e4bc3106885f6348c57c90beadd7cbca17570b90130aaadc7eb0ed2bfb97c5c8a791d0b6b2736d43b0f444909449e6ea32ba706ddfe28c407a85e82e3f22064696d2fd7b37b33e01e0016ad91047f95f0cf5c5d6abc8fd7698c4155c290a1a1db842e77e656c69ad8bc06cea9e00dadbec886ca64af5fe3045ebc8c549bf23e71c634f02da8b250618c169a54e4af45d799f2de6747bde01395a13f1a605e9be5', '0x9ced04a51e77ed0730f3420571164ea5247a670b962ebf6b453660748ca5']}`)
	var res bellatrix.ExecutionPayload
	if err := yaml.Unmarshal(data, &res); err != nil {
		panic(err.Error())
	}
	return res.Transactions
}()

var TestingBeaconBlock = &bellatrix.BeaconBlock{
	Slot:          12,
	ProposerIndex: 10,
	ParentRoot:    spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	StateRoot:     spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
	Body: &bellatrix.BeaconBlockBody{
		RANDAOReveal: spec.BLSSignature{},
		ETH1Data: &spec.ETH1Data{
			DepositRoot:  spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
			DepositCount: 100,
			BlockHash:    []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
		},
		Graffiti:          []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
		ProposerSlashings: []*spec.ProposerSlashing{},
		AttesterSlashings: []*spec.AttesterSlashing{},
		Attestations: []*spec.Attestation{
			{
				AggregationBits: bitfield.NewBitlist(122),
				Data:            TestingAttestationData,
				Signature:       spec.BLSSignature{},
			},
		},
		Deposits:       []*spec.Deposit{},
		VoluntaryExits: []*spec.SignedVoluntaryExit{},
		SyncAggregate: &altair.SyncAggregate{
			SyncCommitteeBits:      bitfield.NewBitvector512(),
			SyncCommitteeSignature: spec.BLSSignature{},
		},
		ExecutionPayload: &bellatrix.ExecutionPayload{
			ParentHash:    spec.Hash32{},
			FeeRecipient:  bellatrix.ExecutionAddress{},
			StateRoot:     spec.Hash32{},
			ReceiptsRoot:  spec.Hash32{},
			LogsBloom:     [256]byte{},
			PrevRandao:    [32]byte{},
			BlockNumber:   100,
			GasLimit:      1000000,
			GasUsed:       800000,
			Timestamp:     123456789,
			BaseFeePerGas: [32]byte{},
			BlockHash:     spec.Hash32{},
			Transactions:  Transactions,
		},
	},
}
var TestingBlindedBeaconBlock = func() *v1.BlindedBeaconBlock {
	fullBlk := TestingBeaconBlock
	txRoot, _ := types.SSZTransactions(fullBlk.Body.ExecutionPayload.Transactions).HashTreeRoot()
	ret := &v1.BlindedBeaconBlock{
		Slot:          fullBlk.Slot,
		ProposerIndex: fullBlk.ProposerIndex,
		ParentRoot:    fullBlk.ParentRoot,
		StateRoot:     fullBlk.StateRoot,
		Body: &v1.BlindedBeaconBlockBody{
			RANDAOReveal:      fullBlk.Body.RANDAOReveal,
			ETH1Data:          fullBlk.Body.ETH1Data,
			Graffiti:          fullBlk.Body.Graffiti,
			ProposerSlashings: fullBlk.Body.ProposerSlashings,
			AttesterSlashings: fullBlk.Body.AttesterSlashings,
			Attestations:      fullBlk.Body.Attestations,
			Deposits:          fullBlk.Body.Deposits,
			VoluntaryExits:    fullBlk.Body.VoluntaryExits,
			SyncAggregate:     fullBlk.Body.SyncAggregate,
			ExecutionPayloadHeader: &bellatrix.ExecutionPayloadHeader{
				ParentHash:       fullBlk.Body.ExecutionPayload.ParentHash,
				FeeRecipient:     fullBlk.Body.ExecutionPayload.FeeRecipient,
				StateRoot:        fullBlk.Body.ExecutionPayload.StateRoot,
				ReceiptsRoot:     fullBlk.Body.ExecutionPayload.ReceiptsRoot,
				LogsBloom:        fullBlk.Body.ExecutionPayload.LogsBloom,
				PrevRandao:       fullBlk.Body.ExecutionPayload.PrevRandao,
				BlockNumber:      fullBlk.Body.ExecutionPayload.BlockNumber,
				GasLimit:         fullBlk.Body.ExecutionPayload.GasLimit,
				GasUsed:          fullBlk.Body.ExecutionPayload.GasUsed,
				Timestamp:        fullBlk.Body.ExecutionPayload.Timestamp,
				ExtraData:        fullBlk.Body.ExecutionPayload.ExtraData,
				BaseFeePerGas:    fullBlk.Body.ExecutionPayload.BaseFeePerGas,
				BlockHash:        fullBlk.Body.ExecutionPayload.BlockHash,
				TransactionsRoot: txRoot,
			},
		},
	}

	return ret
}()
var TestingWrongBeaconBlock = func() *bellatrix.BeaconBlock {
	byts, err := TestingBeaconBlock.MarshalSSZ()
	if err != nil {
		panic(err.Error())
	}
	ret := &bellatrix.BeaconBlock{}
	if err := ret.UnmarshalSSZ(byts); err != nil {
		panic(err.Error())
	}
	ret.Slot = 100
	return ret
}()

var TestingSignedBeaconBlock = func(ks *TestKeySet) *bellatrix.SignedBeaconBlock {
	return &bellatrix.SignedBeaconBlock{
		Message:   TestingBeaconBlock,
		Signature: signBeaconObject(TestingBeaconBlock, types.DomainProposer, ks),
	}
}

var TestingAggregateAndProof = &spec.AggregateAndProof{
	AggregatorIndex: 1,
	SelectionProof:  spec.BLSSignature{},
	Aggregate: &spec.Attestation{
		AggregationBits: bitfield.NewBitlist(128),
		Signature:       spec.BLSSignature{},
		Data:            TestingAttestationData,
	},
}
var TestingWrongAggregateAndProof = func() *spec.AggregateAndProof {
	byts, err := TestingAggregateAndProof.MarshalSSZ()
	if err != nil {
		panic(err.Error())
	}
	ret := &spec.AggregateAndProof{}
	if err := ret.UnmarshalSSZ(byts); err != nil {
		panic(err.Error())
	}
	ret.AggregatorIndex = 100
	return ret
}()

var TestingSignedAggregateAndProof = func(ks *TestKeySet) *spec.SignedAggregateAndProof {
	return &spec.SignedAggregateAndProof{
		Message:   TestingAggregateAndProof,
		Signature: signBeaconObject(TestingAggregateAndProof, types.DomainAggregateAndProof, ks),
	}
}

const (
	TestingDutySlot       = 12
	TestingDutySlot2      = 50
	TestingDutyEpoch      = 0
	TestingDutyEpoch2     = 1
	TestingValidatorIndex = 1

	UnknownDutyType = 100
)

var TestingSyncCommitteeBlockRoot = spec.Root{}
var TestingSyncCommitteeWrongBlockRoot = spec.Root{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var TestingSignedSyncCommitteeBlockRoot = func(ks *TestKeySet) *altair.SyncCommitteeMessage {
	return &altair.SyncCommitteeMessage{
		Slot:            TestingDutySlot,
		BeaconBlockRoot: TestingSyncCommitteeBlockRoot,
		ValidatorIndex:  TestingValidatorIndex,
		Signature:       signBeaconObject(types.SSZBytes(TestingSyncCommitteeBlockRoot[:]), types.DomainSyncCommittee, ks),
	}
}

var TestingContributionProofIndexes = []spec.CommitteeIndex{0, 1, 2}
var TestingContributionProofsSigned = func() []spec.BLSSignature {
	// signed with 3515c7d08e5affd729e9579f7588d30f2342ee6f6a9334acf006345262162c6f
	byts1, _ := hex.DecodeString("b18833bb7549ec33e8ac414ba002fd45bb094ca300bd24596f04a434a89beea462401da7c6b92fb3991bd17163eb603604a40e8dd6781266c990023446776ff42a9313df26a0a34184a590e57fa4003d610c2fa214db4e7dec468592010298bc")
	byts2, _ := hex.DecodeString("9094342c95146554df849dc20f7425fca692dacee7cb45258ddd264a8e5929861469fda3d1567b9521cba83188ffd61a0dbe6d7180c7a96f5810d18db305e9143772b766d368aa96d3751f98d0ce2db9f9e6f26325702088d87f0de500c67c68")
	byts3, _ := hex.DecodeString("a7f88ce43eff3aa8cdd2e3957c5bead4e21353fbecac6079a5398d03019bc45ff7c951785172deee70e9bc5abbc8ca6a0f0441e9d4cc9da74c31121357f7d7c7de9533f6f457da493e3314e22d554ab76613e469b050e246aff539a33807197c")

	ret := make([]spec.BLSSignature, 0)
	for _, byts := range [][]byte{byts1, byts2, byts3} {
		b := spec.BLSSignature{}
		copy(b[:], byts)
		ret = append(ret, b)
	}
	return ret
}()

var TestingSyncCommitteeContributions = []*altair.SyncCommitteeContribution{
	{
		Slot:              TestingDutySlot,
		BeaconBlockRoot:   TestingSyncCommitteeBlockRoot,
		SubcommitteeIndex: 0,
		AggregationBits:   bitfield.NewBitvector128(),
		Signature:         spec.BLSSignature{},
	},
	{
		Slot:              TestingDutySlot,
		BeaconBlockRoot:   TestingSyncCommitteeBlockRoot,
		SubcommitteeIndex: 1,
		AggregationBits:   bitfield.NewBitvector128(),
		Signature:         spec.BLSSignature{},
	},
	{
		Slot:              TestingDutySlot,
		BeaconBlockRoot:   TestingSyncCommitteeBlockRoot,
		SubcommitteeIndex: 2,
		AggregationBits:   bitfield.NewBitvector128(),
		Signature:         spec.BLSSignature{},
	},
}

var TestingSignedSyncCommitteeContributions = func(
	contrib *altair.SyncCommitteeContribution,
	proof spec.BLSSignature,
	ks *TestKeySet) *altair.SignedContributionAndProof {
	msg := &altair.ContributionAndProof{
		AggregatorIndex: TestingValidatorIndex,
		Contribution:    contrib,
		SelectionProof:  proof,
	}
	return &altair.SignedContributionAndProof{
		Message:   msg,
		Signature: signBeaconObject(msg, types.DomainContributionAndProof, ks),
	}
}

var TestingAttesterDuty = &types.Duty{
	Type:                    types.BNRoleAttester,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          3,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

var TestingProposerDuty = &types.Duty{
	Type:                    types.BNRoleProposer,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          3,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

// TestingProposerDutyNextEpoch testing for a second duty start
var TestingProposerDutyNextEpoch = &types.Duty{
	Type:                    types.BNRoleProposer,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot2,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          3,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

var TestingAggregatorDuty = &types.Duty{
	Type:                    types.BNRoleAggregator,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          22,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

// TestingAggregatorDutyNextEpoch testing for a second duty start
var TestingAggregatorDutyNextEpoch = &types.Duty{
	Type:                    types.BNRoleAggregator,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    TestingDutySlot2,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          22,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

var TestingSyncCommitteeDuty = &types.Duty{
	Type:                          types.BNRoleSyncCommittee,
	PubKey:                        TestingValidatorPubKey,
	Slot:                          TestingDutySlot,
	ValidatorIndex:                TestingValidatorIndex,
	CommitteeIndex:                3,
	CommitteesAtSlot:              36,
	CommitteeLength:               128,
	ValidatorCommitteeIndex:       11,
	ValidatorSyncCommitteeIndices: TestingContributionProofIndexes,
}

var TestingSyncCommitteeContributionDuty = &types.Duty{
	Type:                          types.BNRoleSyncCommitteeContribution,
	PubKey:                        TestingValidatorPubKey,
	Slot:                          TestingDutySlot,
	ValidatorIndex:                TestingValidatorIndex,
	CommitteeIndex:                3,
	CommitteesAtSlot:              36,
	CommitteeLength:               128,
	ValidatorCommitteeIndex:       11,
	ValidatorSyncCommitteeIndices: TestingContributionProofIndexes,
}

// TestingSyncCommitteeContributionNexEpochDuty testing for a second duty start
var TestingSyncCommitteeContributionNexEpochDuty = &types.Duty{
	Type:                          types.BNRoleSyncCommitteeContribution,
	PubKey:                        TestingValidatorPubKey,
	Slot:                          TestingDutySlot2,
	ValidatorIndex:                TestingValidatorIndex,
	CommitteeIndex:                3,
	CommitteesAtSlot:              36,
	CommitteeLength:               128,
	ValidatorCommitteeIndex:       11,
	ValidatorSyncCommitteeIndices: TestingContributionProofIndexes,
}

var TestingUnknownDutyType = &types.Duty{
	Type:                    UnknownDutyType,
	PubKey:                  TestingValidatorPubKey,
	Slot:                    12,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          22,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

var TestingWrongDutyPK = &types.Duty{
	Type:                    types.BNRoleAttester,
	PubKey:                  TestingWrongValidatorPubKey,
	Slot:                    12,
	ValidatorIndex:          TestingValidatorIndex,
	CommitteeIndex:          3,
	CommitteesAtSlot:        36,
	CommitteeLength:         128,
	ValidatorCommitteeIndex: 11,
}

//func blsSigFromHex(str string) spec.BLSSignature {
//	byts, _ := hex.DecodeString(str)
//	ret := spec.BLSSignature{}
//	copy(ret[:], byts)
//	return ret
//}

type TestingBeaconNode struct {
	BroadcastedRoots             []spec.Root
	syncCommitteeAggregatorRoots map[string]bool
}

func NewTestingBeaconNode() *TestingBeaconNode {
	return &TestingBeaconNode{
		BroadcastedRoots: []spec.Root{},
	}
}

// SetSyncCommitteeAggregatorRootHexes FOR TESTING ONLY!! sets which sync committee aggregator roots will return true for aggregator
func (bn *TestingBeaconNode) SetSyncCommitteeAggregatorRootHexes(roots map[string]bool) {
	bn.syncCommitteeAggregatorRoots = roots
}

// GetBeaconNetwork returns the beacon network the node is on
func (bn *TestingBeaconNode) GetBeaconNetwork() types.BeaconNetwork {
	return types.BeaconTestNetwork
}

// GetAttestationData returns attestation data by the given slot and committee index
func (bn *TestingBeaconNode) GetAttestationData(slot spec.Slot, committeeIndex spec.CommitteeIndex) (*spec.AttestationData, error) {
	return TestingAttestationData, nil
}

// SubmitAttestation submit the attestation to the node
func (bn *TestingBeaconNode) SubmitAttestation(attestation *spec.Attestation) error {
	r, _ := attestation.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

// GetBeaconBlock returns beacon block by the given slot and committee index
func (bn *TestingBeaconNode) GetBeaconBlock(slot spec.Slot, graffiti, randao []byte) (*bellatrix.BeaconBlock, error) {
	return TestingBeaconBlock, nil
}

// SubmitBeaconBlock submit the block to the node
func (bn *TestingBeaconNode) SubmitBeaconBlock(block *bellatrix.SignedBeaconBlock) error {
	r, _ := block.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

// GetBlindedBeaconBlock returns blinded beacon block by the given slot and committee index
func (bn *TestingBeaconNode) GetBlindedBeaconBlock(slot spec.Slot, committeeIndex spec.CommitteeIndex, graffiti, randao []byte) (*v1.BlindedBeaconBlock, error) {
	return TestingBlindedBeaconBlock, nil
}

// SubmitBlindedBeaconBlock submit the blinded block to the node
func (bn *TestingBeaconNode) SubmitBlindedBeaconBlock(block *v1.SignedBlindedBeaconBlock) error {
	r, _ := block.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

// SubmitAggregateSelectionProof returns an AggregateAndProof object
func (bn *TestingBeaconNode) SubmitAggregateSelectionProof(slot spec.Slot, committeeIndex spec.CommitteeIndex, committeeLength uint64, index spec.ValidatorIndex, slotSig []byte) (*spec.AggregateAndProof, error) {
	return TestingAggregateAndProof, nil
}

// SubmitSignedAggregateSelectionProof broadcasts a signed aggregator msg
func (bn *TestingBeaconNode) SubmitSignedAggregateSelectionProof(msg *spec.SignedAggregateAndProof) error {
	r, _ := msg.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

// GetSyncMessageBlockRoot returns beacon block root for sync committee
func (bn *TestingBeaconNode) GetSyncMessageBlockRoot(slot spec.Slot) (spec.Root, error) {
	return TestingSyncCommitteeBlockRoot, nil
}

// SubmitSyncMessage submits a signed sync committee msg
func (bn *TestingBeaconNode) SubmitSyncMessage(msg *altair.SyncCommitteeMessage) error {
	r, _ := msg.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

// IsSyncCommitteeAggregator returns tru if aggregator
func (bn *TestingBeaconNode) IsSyncCommitteeAggregator(proof []byte) (bool, error) {
	if len(bn.syncCommitteeAggregatorRoots) != 0 {
		if val, found := bn.syncCommitteeAggregatorRoots[hex.EncodeToString(proof)]; found {
			return val, nil
		}
		return false, nil
	}
	return true, nil
}

// SyncCommitteeSubnetID returns sync committee subnet ID from subcommittee index
func (bn *TestingBeaconNode) SyncCommitteeSubnetID(index spec.CommitteeIndex) (uint64, error) {
	// each subcommittee index correlates to TestingContributionProofRoots by index
	return uint64(index), nil
}

// GetSyncCommitteeContribution returns
func (bn *TestingBeaconNode) GetSyncCommitteeContribution(slot spec.Slot, subnetID uint64) (*altair.SyncCommitteeContribution, error) {
	return TestingSyncCommitteeContributions[subnetID], nil
}

// SubmitSignedContributionAndProof broadcasts to the network
func (bn *TestingBeaconNode) SubmitSignedContributionAndProof(contribution *altair.SignedContributionAndProof) error {
	r, _ := contribution.HashTreeRoot()
	bn.BroadcastedRoots = append(bn.BroadcastedRoots, r)
	return nil
}

func (bn *TestingBeaconNode) DomainData(epoch spec.Epoch, domain spec.DomainType) (spec.Domain, error) {
	// epoch is used to calculate fork version, here we hard code it
	return types.ComputeETHDomain(domain, types.GenesisForkVersion, types.GenesisValidatorsRoot)
}
