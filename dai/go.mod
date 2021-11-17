module github.com/PaddlePaddle/PaddleDTX/dai

go 1.16

require (
	github.com/PaddlePaddle/PaddleDTX/crypto v0.0.0
	github.com/PaddlePaddle/PaddleDTX/xdb v0.0.0
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/xuperchain/xuperchain v0.0.0-20210208123615-2d08ff11de3e
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf
	google.golang.org/grpc v1.41.0
)

replace github.com/PaddlePaddle/PaddleDTX/crypto => icode.baidu.com/baidu/blockchain/PaddleDTX/crypto v0.0.0-20211116114245-2c492c5a07ae

replace github.com/PaddlePaddle/PaddleDTX/xdb => icode.baidu.com/baidu/blockchain/PaddleDTX/xdb v0.0.0-20211116114245-2c492c5a07ae