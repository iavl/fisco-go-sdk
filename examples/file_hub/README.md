

# 配置文件(config.toml)

```toml
[Network]
#type rpc or channel
Type="rpc"
CAFile="./sdk/gm/ca.crt"
Cert="./sdk/gm/sdk.crt"
Key="./sdk/gm/sdk.key"
[[Network.Connection]]
NodeURL="121.196.203.127:8545"
GroupID=1
[Account]
# only support PEM format for now
KeyFile="./accounts_gm/0x1d29e94559b887e32020074c190522d5354f121b.pem"

[Chain]
ChainID=1
SMCrypto=false
```


## Account

- KeyFile：节点签发交易时所使用的私钥，PEM格式，支持国密和非国密。

请使用[get_account.sh](https://github.com/FISCO-BCOS/console/blob/master/tools/get_account.sh)和[get_gm_account.sh](https://github.com/FISCO-BCOS/console/blob/master/tools/get_gm_account.sh)脚本生成。使用方式[参考这里](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/manual/account.html)。

如果想使用Go-SDK代码生成，请[参考这里](doc/README.md#环境配置#外部账户)。

# Package功能使用

以下的示例是通过`import`的方式来使用`go-sdk`，如引入RPC控制台库:

```go
import "github.com/iavl/fisco-go-sdk/client"
```

## Solidity合约编译为Go文件

在利用SDK进行项目开发时，对智能合约进行操作时需要将Solidity智能合约利用go-sdk的`abigen`工具转换为`Go`文件代码。整体上主要包含了五个流程：

- 准备需要编译的智能合约
- 配置好相应版本的`solc`编译器
- 构建go-sdk的合约编译工具`abigen`
- 编译生成go文件
- 使用生成的go文件进行合约调用

下面的内容作为一个示例进行使用介绍。

1.样例智能合约 FileHub.sol

2.安装对应版本的[`solc`编译器](https://github.com/ethereum/solidity/releases/tag/v0.4.25)，目前FISCO BCOS默认的`solc`编译器版本为0.4.25。

```bash
# 如果是国密则添加-g选项
bash tools/download_solc.sh -v 0.6.10 -g
```

3.构建`go-sdk`的代码生成工具`abigen`

```bash
# 下面指令在go-sdk目录下操作，编译生成abigen工具
go build ./cmd/abigen
```

4.编译生成go文件，先利用`solc`将合约文件生成`abi`和`bin`文件，以前面所提供的`FileHub.sol`为例：

```bash
./solc-0.6.10-gm --bin --abi -o ./examples/file_hub ./examples/file_hub/FileHub.sol
```

```bash
./abigen --bin ./examples/file_hub/FileHub.bin --abi ./examples/file_hub/FileHub.abi --pkg store --type FileHub --out ./examples/file_hub/FileHub.go --smcrypto=true
```

最后file_hub目录下面存在以下文件：

```bash
FileHub.abi FileHub.bin FileHub.go  FileHub.sol
```
