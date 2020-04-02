package blockchain

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/smartcontractkit/external-initiator/store"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var (
	substrateTestMetadataHex = "0x6d657461092c1853797374656d011853797374656d34304163636f756e744e6f6e636501010130543a3a4163636f756e74496420543a3a496e646578001000000000047c2045787472696e73696373206e6f6e636520666f72206163636f756e74732e3845787472696e736963436f756e7400000c753332040004b820546f74616c2065787472696e7369637320636f756e7420666f72207468652063757272656e7420626c6f636b2e4c416c6c45787472696e73696373576569676874000018576569676874040004150120546f74616c2077656967687420666f7220616c6c2065787472696e736963732070757420746f6765746865722c20666f72207468652063757272656e7420626c6f636b2e40416c6c45787472696e736963734c656e00000c753332040004410120546f74616c206c656e6774682028696e2062797465732920666f7220616c6c2065787472696e736963732070757420746f6765746865722c20666f72207468652063757272656e7420626c6f636b2e24426c6f636b4861736801010138543a3a426c6f636b4e756d6265721c543a3a48617368008000000000000000000000000000000000000000000000000000000000000000000498204d6170206f6620626c6f636b206e756d6265727320746f20626c6f636b206861736865732e3445787472696e736963446174610101010c7533321c5665633c75383e000400043d012045787472696e73696373206461746120666f72207468652063757272656e7420626c6f636b20286d61707320616e2065787472696e736963277320696e64657820746f206974732064617461292e184e756d626572010038543a3a426c6f636b4e756d6265721000000000040901205468652063757272656e7420626c6f636b206e756d626572206265696e672070726f6365737365642e205365742062792060657865637574655f626c6f636b602e28506172656e744861736801001c543a3a4861736880000000000000000000000000000000000000000000000000000000000000000004702048617368206f66207468652070726576696f757320626c6f636b2e3845787472696e73696373526f6f7401001c543a3a486173688000000000000000000000000000000000000000000000000000000000000000000415012045787472696e7369637320726f6f74206f66207468652063757272656e7420626c6f636b2c20616c736f2070617274206f662074686520626c6f636b206865616465722e1844696765737401002c4469676573744f663c543e040004f020446967657374206f66207468652063757272656e7420626c6f636b2c20616c736f2070617274206f662074686520626c6f636b206865616465722e184576656e747301008c5665633c4576656e745265636f72643c543a3a4576656e742c20543a3a486173683e3e040004a0204576656e7473206465706f736974656420666f72207468652063757272656e7420626c6f636b2e284576656e74436f756e740100284576656e74496e646578100000000004b820546865206e756d626572206f66206576656e747320696e2074686520604576656e74733c543e60206c6973742e2c4576656e74546f706963730102010828291c543a3a48617368845665633c28543a3a426c6f636b4e756d6265722c204576656e74496e646578293e010400342501204d617070696e67206265747765656e206120746f7069632028726570726573656e74656420627920543a3a486173682920616e64206120766563746f72206f6620696e646578657394206f66206576656e747320696e2074686520603c4576656e74733c543e3e60206c6973742e002d0120546865206669727374206b657920736572766573206e6f20707572706f73652e2054686973206669656c64206973206465636c6172656420617320646f75626c655f6d6170206a757374a820666f7220636f6e76656e69656e6365206f66207573696e67206072656d6f76655f707265666978602e00510120416c6c20746f70696320766563746f727320686176652064657465726d696e69737469632073746f72616765206c6f636174696f6e7320646570656e64696e67206f6e2074686520746f7069632e2054686973450120616c6c6f7773206c696768742d636c69656e747320746f206c6576657261676520746865206368616e67657320747269652073746f7261676520747261636b696e67206d656368616e69736d20616e64e420696e2063617365206f66206368616e67657320666574636820746865206c697374206f66206576656e7473206f6620696e7465726573742e004d01205468652076616c756520686173207468652074797065206028543a3a426c6f636b4e756d6265722c204576656e74496e646578296020626563617573652069662077652075736564206f6e6c79206a7573744d012074686520604576656e74496e64657860207468656e20696e20636173652069662074686520746f70696320686173207468652073616d6520636f6e74656e7473206f6e20746865206e65787420626c6f636b0101206e6f206e6f74696669636174696f6e2077696c6c20626520747269676765726564207468757320746865206576656e74206d69676874206265206c6f73742e011c2866696c6c5f626c6f636b0004210120412062696720646973706174636820746861742077696c6c20646973616c6c6f7720616e79206f74686572207472616e73616374696f6e20746f20626520696e636c756465642e1872656d61726b041c5f72656d61726b1c5665633c75383e046c204d616b6520736f6d65206f6e2d636861696e2072656d61726b2e387365745f686561705f7061676573041470616765730c75363404fc2053657420746865206e756d626572206f6620706167657320696e2074686520576562417373656d626c7920656e7669726f6e6d656e74277320686561702e207365745f636f6465040c6e65771c5665633c75383e04482053657420746865206e657720636f64652e2c7365745f73746f7261676504146974656d73345665633c4b657956616c75653e046c2053657420736f6d65206974656d73206f662073746f726167652e306b696c6c5f73746f7261676504106b657973205665633c4b65793e0478204b696c6c20736f6d65206974656d732066726f6d2073746f726167652e2c6b696c6c5f70726566697804187072656669780c4b6579041501204b696c6c20616c6c2073746f72616765206974656d7320776974682061206b657920746861742073746172747320776974682074686520676976656e207072656669782e01084045787472696e7369635375636365737304304469737061746368496e666f049420416e2065787472696e73696320636f6d706c65746564207375636365737366756c6c792e3c45787472696e7369634661696c6564083444697370617463684572726f72304469737061746368496e666f045420416e2065787472696e736963206661696c65642e000c4c526571756972655369676e65644f726967696e004452657175697265526f6f744f726967696e003c526571756972654e6f4f726967696e002454696d657374616d70012454696d657374616d70080c4e6f77010024543a3a4d6f6d656e7420000000000000000004902043757272656e742074696d6520666f72207468652063757272656e7420626c6f636b2e24446964557064617465010010626f6f6c040004b420446964207468652074696d657374616d7020676574207570646174656420696e207468697320626c6f636b3f01040c736574040c6e6f7748436f6d706163743c543a3a4d6f6d656e743e245820536574207468652063757272656e742074696d652e00590120546869732063616c6c2073686f756c6420626520696e766f6b65642065786163746c79206f6e63652070657220626c6f636b2e2049742077696c6c2070616e6963206174207468652066696e616c697a6174696f6ed82070686173652c20696620746869732063616c6c206861736e2774206265656e20696e766f6b656420627920746861742074696d652e004501205468652074696d657374616d702073686f756c642062652067726561746572207468616e207468652070726576696f7573206f6e652062792074686520616d6f756e74207370656369666965642062794420604d696e696d756d506572696f64602e00d820546865206469737061746368206f726967696e20666f7220746869732063616c6c206d7573742062652060496e686572656e74602e0004344d696e696d756d506572696f6424543a3a4d6f6d656e7420b80b00000000000010690120546865206d696e696d756d20706572696f64206265747765656e20626c6f636b732e204265776172652074686174207468697320697320646966666572656e7420746f20746865202a65787065637465642a20706572696f64690120746861742074686520626c6f636b2070726f64756374696f6e206170706172617475732070726f76696465732e20596f75722063686f73656e20636f6e73656e7375732073797374656d2077696c6c2067656e6572616c6c79650120776f726b2077697468207468697320746f2064657465726d696e6520612073656e7369626c6520626c6f636b2074696d652e20652e672e20466f7220417572612c2069742077696c6c20626520646f75626c6520746869737020706572696f64206f6e2064656661756c742073657474696e67732e00104175726100000000001c4772616e647061013c4772616e64706146696e616c6974791c2c417574686f726974696573010034417574686f726974794c6973740400102c20444550524543415445440061012054686973207573656420746f2073746f7265207468652063757272656e7420617574686f72697479207365742c20776869636820686173206265656e206d6967726174656420746f207468652077656c6c2d6b6e6f776e94204752414e4450415f415554484f52495445535f4b455920756e686173686564206b65792e14537461746501006c53746f72656453746174653c543a3a426c6f636b4e756d6265723e04000490205374617465206f66207468652063757272656e7420617574686f72697479207365742e3450656e64696e674368616e676500008c53746f72656450656e64696e674368616e67653c543a3a426c6f636b4e756d6265723e040004c42050656e64696e67206368616e67653a20287369676e616c65642061742c207363686564756c6564206368616e6765292e284e657874466f72636564000038543a3a426c6f636b4e756d626572040004bc206e65787420626c6f636b206e756d6265722077686572652077652063616e20666f7263652061206368616e67652e1c5374616c6c656400008028543a3a426c6f636b4e756d6265722c20543a3a426c6f636b4e756d626572290400049020607472756560206966207765206172652063757272656e746c79207374616c6c65642e3043757272656e7453657449640100145365744964200000000000000000085d0120546865206e756d626572206f66206368616e6765732028626f746820696e207465726d73206f66206b65797320616e6420756e6465726c79696e672065636f6e6f6d696320726573706f6e736962696c697469657329c420696e20746865202273657422206f66204772616e6470612076616c696461746f72732066726f6d2067656e657369732e30536574496453657373696f6e0001011453657449643053657373696f6e496e64657800040004c1012041206d617070696e672066726f6d206772616e6470612073657420494420746f2074686520696e646578206f6620746865202a6d6f737420726563656e742a2073657373696f6e20666f7220776869636820697473206d656d62657273207765726520726573706f6e7369626c652e0104487265706f72745f6d69736265686176696f72041c5f7265706f72741c5665633c75383e0464205265706f727420736f6d65206d69736265686176696f722e010c384e6577417574686f7269746965730434417574686f726974794c6973740490204e657720617574686f726974792073657420686173206265656e206170706c6965642e1850617573656400049c2043757272656e7420617574686f726974792073657420686173206265656e207061757365642e1c526573756d65640004a02043757272656e7420617574686f726974792073657420686173206265656e20726573756d65642e00001c496e6469636573011c496e6469636573082c4e657874456e756d53657401003c543a3a4163636f756e74496e6465781000000000047c20546865206e657874206672656520656e756d65726174696f6e207365742e1c456e756d5365740101013c543a3a4163636f756e74496e646578445665633c543a3a4163636f756e7449643e00040004582054686520656e756d65726174696f6e20736574732e010001043c4e65774163636f756e74496e64657808244163636f756e744964304163636f756e74496e64657810882041206e6577206163636f756e7420696e646578207761732061737369676e65642e0005012054686973206576656e74206973206e6f7420747269676765726564207768656e20616e206578697374696e6720696e64657820697320726561737369676e65646020746f20616e6f7468657220604163636f756e744964602e00002042616c616e636573012042616c616e6365731434546f74616c49737375616e6365010028543a3a42616c616e6365400000000000000000000000000000000004982054686520746f74616c20756e6974732069737375656420696e207468652073797374656d2e1c56657374696e6700010130543a3a4163636f756e744964ac56657374696e675363686564756c653c543a3a42616c616e63652c20543a3a426c6f636b4e756d6265723e00040004d820496e666f726d6174696f6e20726567617264696e67207468652076657374696e67206f66206120676976656e206163636f756e742e2c4672656542616c616e636501010130543a3a4163636f756e74496428543a3a42616c616e63650040000000000000000000000000000000002c9c20546865202766726565272062616c616e6365206f66206120676976656e206163636f756e742e004101205468697320697320746865206f6e6c792062616c616e63652074686174206d61747465727320696e207465726d73206f66206d6f7374206f7065726174696f6e73206f6e20746f6b656e732e204974750120616c6f6e65206973207573656420746f2064657465726d696e65207468652062616c616e6365207768656e20696e2074686520636f6e747261637420657865637574696f6e20656e7669726f6e6d656e742e205768656e207468697355012062616c616e63652066616c6c732062656c6f77207468652076616c7565206f6620604578697374656e7469616c4465706f736974602c207468656e20746865202763757272656e74206163636f756e74272069733d012064656c657465643a207370656369666963616c6c7920604672656542616c616e6365602e20467572746865722c2074686520604f6e4672656542616c616e63655a65726f602063616c6c6261636b450120697320696e766f6b65642c20676976696e672061206368616e636520746f2065787465726e616c206d6f64756c657320746f20636c65616e2075702064617461206173736f636961746564207769746854207468652064656c65746564206163636f756e742e00750120606672616d655f73797374656d3a3a4163636f756e744e6f6e63656020697320616c736f2064656c657465642069662060526573657276656442616c616e63656020697320616c736f207a65726f2028697420616c736f2067657473150120636f6c6c617073656420746f207a65726f2069662069742065766572206265636f6d6573206c657373207468616e20604578697374656e7469616c4465706f736974602e3c526573657276656442616c616e636501010130543a3a4163636f756e74496428543a3a42616c616e63650040000000000000000000000000000000002c75012054686520616d6f756e74206f66207468652062616c616e6365206f66206120676976656e206163636f756e7420746861742069732065787465726e616c6c792072657365727665643b20746869732063616e207374696c6c206765749c20736c61736865642c20627574206765747320736c6173686564206c617374206f6620616c6c2e006d0120546869732062616c616e63652069732061202772657365727665272062616c616e63652074686174206f746865722073756273797374656d732075736520696e206f7264657220746f2073657420617369646520746f6b656e732501207468617420617265207374696c6c20276f776e65642720627920746865206163636f756e7420686f6c6465722c20627574207768696368206172652073757370656e6461626c652e007501205768656e20746869732062616c616e63652066616c6c732062656c6f77207468652076616c7565206f6620604578697374656e7469616c4465706f736974602c207468656e2074686973202772657365727665206163636f756e7427b42069732064656c657465643a207370656369666963616c6c792c2060526573657276656442616c616e6365602e00650120606672616d655f73797374656d3a3a4163636f756e744e6f6e63656020697320616c736f2064656c6574656420696620604672656542616c616e63656020697320616c736f207a65726f2028697420616c736f2067657473190120636f6c6c617073656420746f207a65726f2069662069742065766572206265636f6d6573206c657373207468616e20604578697374656e7469616c4465706f736974602e29144c6f636b7301010130543a3a4163636f756e744964b05665633c42616c616e63654c6f636b3c543a3a42616c616e63652c20543a3a426c6f636b4e756d6265723e3e00040004b820416e79206c6971756964697479206c6f636b73206f6e20736f6d65206163636f756e742062616c616e6365732e0110207472616e736665720810646573748c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f757263651476616c75654c436f6d706163743c543a3a42616c616e63653e64d8205472616e7366657220736f6d65206c697175696420667265652062616c616e636520746f20616e6f74686572206163636f756e742e00090120607472616e73666572602077696c6c207365742074686520604672656542616c616e636560206f66207468652073656e64657220616e642072656365697665722e21012049742077696c6c2064656372656173652074686520746f74616c2069737375616e6365206f66207468652073797374656d2062792074686520605472616e73666572466565602e1501204966207468652073656e6465722773206163636f756e742069732062656c6f7720746865206578697374656e7469616c206465706f736974206173206120726573756c74b4206f6620746865207472616e736665722c20746865206163636f756e742077696c6c206265207265617065642e00190120546865206469737061746368206f726967696e20666f7220746869732063616c6c206d75737420626520605369676e65646020627920746865207472616e736163746f722e002c2023203c7765696768743e3101202d20446570656e64656e74206f6e20617267756d656e747320627574206e6f7420637269746963616c2c20676976656e2070726f70657220696d706c656d656e746174696f6e7320666f72cc202020696e70757420636f6e6669672074797065732e205365652072656c617465642066756e6374696f6e732062656c6f772e6901202d20497420636f6e7461696e732061206c696d69746564206e756d626572206f6620726561647320616e642077726974657320696e7465726e616c6c7920616e64206e6f20636f6d706c657820636f6d7075746174696f6e2e004c2052656c617465642066756e6374696f6e733a0051012020202d2060656e737572655f63616e5f77697468647261776020697320616c776179732063616c6c656420696e7465726e616c6c792062757420686173206120626f756e64656420636f6d706c65786974792e2d012020202d205472616e7366657272696e672062616c616e63657320746f206163636f756e7473207468617420646964206e6f74206578697374206265666f72652077696c6c206361757365d420202020202060543a3a4f6e4e65774163636f756e743a3a6f6e5f6e65775f6163636f756e746020746f2062652063616c6c65642edc2020202d2052656d6f76696e6720656e6f7567682066756e64732066726f6d20616e206163636f756e742077696c6c20747269676765725901202020202060543a3a4475737452656d6f76616c3a3a6f6e5f756e62616c616e6365646020616e642060543a3a4f6e4672656542616c616e63655a65726f3a3a6f6e5f667265655f62616c616e63655f7a65726f602e49012020202d20607472616e736665725f6b6565705f616c6976656020776f726b73207468652073616d652077617920617320607472616e73666572602c206275742068617320616e206164646974696f6e616cf82020202020636865636b207468617420746865207472616e736665722077696c6c206e6f74206b696c6c20746865206f726967696e206163636f756e742e00302023203c2f7765696768743e2c7365745f62616c616e63650c0c77686f8c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f75726365206e65775f667265654c436f6d706163743c543a3a42616c616e63653e306e65775f72657365727665644c436f6d706163743c543a3a42616c616e63653e349420536574207468652062616c616e636573206f66206120676976656e206163636f756e742e00210120546869732077696c6c20616c74657220604672656542616c616e63656020616e642060526573657276656442616c616e63656020696e2073746f726167652e2069742077696c6c090120616c736f2064656372656173652074686520746f74616c2069737375616e6365206f66207468652073797374656d202860546f74616c49737375616e636560292e190120496620746865206e65772066726565206f722072657365727665642062616c616e63652069732062656c6f7720746865206578697374656e7469616c206465706f7369742c01012069742077696c6c20726573657420746865206163636f756e74206e6f6e63652028606672616d655f73797374656d3a3a4163636f756e744e6f6e636560292e00b420546865206469737061746368206f726967696e20666f7220746869732063616c6c2069732060726f6f74602e002c2023203c7765696768743e80202d20496e646570656e64656e74206f662074686520617267756d656e74732ec4202d20436f6e7461696e732061206c696d69746564206e756d626572206f6620726561647320616e64207772697465732e302023203c2f7765696768743e38666f7263655f7472616e736665720c18736f757263658c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f7572636510646573748c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f757263651476616c75654c436f6d706163743c543a3a42616c616e63653e0851012045786163746c7920617320607472616e73666572602c2065786365707420746865206f726967696e206d75737420626520726f6f7420616e642074686520736f75726365206163636f756e74206d61792062652c207370656369666965642e4c7472616e736665725f6b6565705f616c6976650810646573748c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f757263651476616c75654c436f6d706163743c543a3a42616c616e63653e1851012053616d6520617320746865205b607472616e73666572605d2063616c6c2c206275742077697468206120636865636b207468617420746865207472616e736665722077696c6c206e6f74206b696c6c2074686540206f726967696e206163636f756e742e00bc20393925206f66207468652074696d6520796f752077616e74205b607472616e73666572605d20696e73746561642e00c4205b607472616e73666572605d3a207374727563742e4d6f64756c652e68746d6c236d6574686f642e7472616e73666572010c284e65774163636f756e7408244163636f756e7449641c42616c616e6365046c2041206e6577206163636f756e742077617320637265617465642e345265617065644163636f756e7404244163636f756e744964045c20416e206163636f756e7420776173207265617065642e205472616e7366657210244163636f756e744964244163636f756e7449641c42616c616e63651c42616c616e636504b0205472616e7366657220737563636565646564202866726f6d2c20746f2c2076616c75652c2066656573292e0c484578697374656e7469616c4465706f73697428543a3a42616c616e636540f401000000000000000000000000000004d420546865206d696e696d756d20616d6f756e7420726571756972656420746f206b65657020616e206163636f756e74206f70656e2e2c5472616e7366657246656528543a3a42616c616e636540000000000000000000000000000000000494205468652066656520726571756972656420746f206d616b652061207472616e736665722e2c4372656174696f6e46656528543a3a42616c616e63654000000000000000000000000000000000049c205468652066656520726571756972656420746f2063726561746520616e206163636f756e742e00485472616e73616374696f6e5061796d656e74012042616c616e63657304444e6578744665654d756c7469706c6965720100284d756c7469706c69657220000000000000000000000008485472616e73616374696f6e426173654665653042616c616e63654f663c543e400000000000000000000000000000000004dc205468652066656520746f206265207061696420666f72206d616b696e672061207472616e73616374696f6e3b2074686520626173652e485472616e73616374696f6e427974654665653042616c616e63654f663c543e4001000000000000000000000000000000040d01205468652066656520746f206265207061696420666f72206d616b696e672061207472616e73616374696f6e3b20746865207065722d6279746520706f7274696f6e2e00105375646f01105375646f040c4b6579010030543a3a4163636f756e74496480000000000000000000000000000000000000000000000000000000000000000004842054686520604163636f756e74496460206f6620746865207375646f206b65792e010c107375646f042070726f706f73616c40426f783c543a3a50726f706f73616c3e2839012041757468656e7469636174657320746865207375646f206b657920616e64206469737061746368657320612066756e6374696f6e2063616c6c20776974682060526f6f7460206f726967696e2e00d020546865206469737061746368206f726967696e20666f7220746869732063616c6c206d757374206265205f5369676e65645f2e002c2023203c7765696768743e20202d204f2831292e64202d204c696d697465642073746f726167652072656164732e60202d204f6e6520444220777269746520286576656e74292ed4202d20556e6b6e6f776e20776569676874206f662064657269766174697665206070726f706f73616c6020657865637574696f6e2e302023203c2f7765696768743e1c7365745f6b6579040c6e65778c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f757263652475012041757468656e74696361746573207468652063757272656e74207375646f206b657920616e6420736574732074686520676976656e204163636f756e7449642028606e6577602920617320746865206e6577207375646f206b65792e00d020546865206469737061746368206f726967696e20666f7220746869732063616c6c206d757374206265205f5369676e65645f2e002c2023203c7765696768743e20202d204f2831292e64202d204c696d697465642073746f726167652072656164732e44202d204f6e65204442206368616e67652e302023203c2f7765696768743e1c7375646f5f6173080c77686f8c3c543a3a4c6f6f6b7570206173205374617469634c6f6f6b75703e3a3a536f757263652070726f706f73616c40426f783c543a3a50726f706f73616c3e2c51012041757468656e7469636174657320746865207375646f206b657920616e64206469737061746368657320612066756e6374696f6e2063616c6c207769746820605369676e656460206f726967696e2066726f6d44206120676976656e206163636f756e742e00d020546865206469737061746368206f726967696e20666f7220746869732063616c6c206d757374206265205f5369676e65645f2e002c2023203c7765696768743e20202d204f2831292e64202d204c696d697465642073746f726167652072656164732e60202d204f6e6520444220777269746520286576656e74292ed4202d20556e6b6e6f776e20776569676874206f662064657269766174697665206070726f706f73616c6020657865637574696f6e2e302023203c2f7765696768743e010c1453756469640410626f6f6c04602041207375646f206a75737420746f6f6b20706c6163652e284b65794368616e67656404244163636f756e74496404f020546865207375646f6572206a757374207377697463686564206964656e746974793b20746865206f6c64206b657920697320737570706c6965642e285375646f4173446f6e650410626f6f6c04602041207375646f206a75737420746f6f6b20706c6163652e000024436861696e6c696e6b0124436861696e6c696e6b0424536f6d657468696e6700000c75333204000001043073656e645f726571756573740424736f6d657468696e670c753332000104344f7261636c655265717565737404244163636f756e7449640000003854656d706c6174654d6f64756c65013854656d706c6174654d6f64756c650424536f6d657468696e6700000c753332040000010430646f5f736f6d657468696e670424736f6d657468696e670c7533320001043c536f6d657468696e6753746f726564080c753332244163636f756e7449640000006052616e646f6d6e657373436f6c6c656374697665466c6970016052616e646f6d6e657373436f6c6c656374697665466c6970043852616e646f6d4d6174657269616c0100305665633c543a3a486173683e04000c610120536572696573206f6620626c6f636b20686561646572732066726f6d20746865206c61737420383120626c6f636b73207468617420616374732061732072616e646f6d2073656564206d6174657269616c2e2054686973610120697320617272616e67656420617320612072696e672062756666657220776974682060626c6f636b5f6e756d626572202520383160206265696e672074686520696e64657820696e746f20746865206056656360206f664420746865206f6c6465737420686173682e0100000000"
	substrateTestAddr1       = "0xd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d"
	substrateTestAddr2       = "0x8eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48"
)

func TestCreateSubstrateManager(t *testing.T) {
	addr1, err := types.NewAddressFromHexAccountID(substrateTestAddr1)
	require.NoError(t, err)
	addr2, err := types.NewAddressFromHexAccountID(substrateTestAddr2)
	require.NoError(t, err)

	tests := []struct {
		name string
		args store.SubstrateSubscription
		want []types.Address
	}{
		{
			"adds valid addresses",
			store.SubstrateSubscription{
				AccountIds: []string{substrateTestAddr1, substrateTestAddr2},
			},
			[]types.Address{addr1, addr2},
		},
		{
			"disregards invalid addresses",
			store.SubstrateSubscription{
				AccountIds: []string{"not a valid address", substrateTestAddr2},
			},
			[]types.Address{addr2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSubstrateManager(store.Subscription{Substrate: tt.args}); !reflect.DeepEqual(got.filter.Address, tt.want) {
				t.Errorf("CreateSubstrateManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstrateManager_GetTestJson(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{
			"returns getMetadata JSON-RPC message",
			[]byte(`{"jsonrpc":"2.0","id":1,"method":"state_getMetadata"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SubstrateManager{}
			if got := sm.GetTestJson(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTestJson() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestSubstrateManager_GetTriggerJson(t *testing.T) {
	var metadata types.Metadata
	require.NoError(t, types.DecodeFromHexString(substrateTestMetadataHex, &metadata))

	tests := []struct {
		name   string
		filter substrateFilter
		want   []byte
	}{
		{
			"generates storagekey from metadata",
			substrateFilter{},
			[]byte(`{"jsonrpc":"2.0","id":1,"method":"state_subscribeStorage","params":[["0x26aa394eea5630e07c48ae0c9558cef780d41e5e16056765bc8461851072c9d7"]]}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SubstrateManager{
				filter: tt.filter,
				meta:   &metadata,
			}
			if got := sm.GetTriggerJson(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTriggerJson() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_convertStringArrayToKV(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want map[string]string
	}{
		{
			"converts standard keys,values",
			[]string{"test", "val", "test2", "val2"},
			map[string]string{
				"test":  "val",
				"test2": "val2",
			},
		},
		{
			"skips empty keys",
			[]string{"test", "val", "", "val2", "test3", "val3"},
			map[string]string{
				"test":  "val",
				"test3": "val3",
			},
		},
		{
			"skips empty values",
			[]string{"test", "val", "key2", "", "test3", "val3", "", "", "test5", "val5"},
			map[string]string{
				"test":  "val",
				"test3": "val3",
				"test5": "val5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStringArrayToKV(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertStringArrayToKV() = %v, want %v", got, tt.want)
			}
		})
	}
}
