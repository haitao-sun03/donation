# EquaSeed Donation

## 概述

EquaSeed是一个基于Web3技术的捐赠平台，是一个全栈应用程序，允许用户管理和向各种活动进行捐赠。它由三个主要组件组成：
1. **前端**：一个React应用程序，提供与捐赠系统交互的用户界面。
2. **后端**：一个Go应用程序，处理业务逻辑、数据库交互和API请求，包含业务服务和用户服务。
3. **智能合约**：部署在以太坊区块链上的Solidity合约，用于管理捐赠和活动。

## 项目结构

### 后端结构

- **backend-end/**: 包含Go后端代码。
  - **business/**: 业务主服务
    - `controllers/`: HTTP请求处理器。
    - `db/`: 数据库访问对象。
    - `event/`: 事件监听器。
    - `model/`: 数据模型。
    - `routers/`: HTTP路由器。
    - `service/`: 业务服务层。
    - `utils/`: 实用工具函数。
    - `vo/`: 值对象。
    - `config/`: 配置文件。
    - `abi/`: 智能合约的ABI文件。
    - `main.go`: 后端应用程序的入口点。
  - **user/**: 用户服务，负责用户注册，登录，颁发token
    - `controllers/`: 用户相关HTTP请求处理器。
    - `db/`: 用户数据库访问对象。
    - `model/`: 用户数据模型。
    - `routers/`: 用户服务HTTP路由器。
    - `service/`: 用户服务层。
    - `utils/`: 用户服务实用工具函数。
    - `vo/`: 用户服务值对象。
    - `main.go`: 用户服务的入口点。

### 智能合约结构

- **donations_manage/**: 包含Solidity智能合约和部署脚本。
  - `contracts/`: Solidity合约文件。
    - `DonationsManageContract.sol`: 管理捐赠活动的创建和管理。
    - `DonationsManageContractV2.sol`: 捐赠管理合约的升级版本。
    - `nft.sol`: 管理与捐赠相关的非同质化代币(NFTs)。
    - `token.sol`: 管理用于捐赠的同质化代币。
    - `temp.sol`: 临时合约文件。
  - `deploy/`: 部署脚本。
  - `deployments/`: 部署构件。
  - `meta_data/`: NFT元数据文件。
  - `hardhat.config.js`: Hardhat配置文件。

### 前端结构

- **front-end/**: 包含React前端代码。
  - `src/`: React应用程序的源代码。
    - `components/`: React组件。
      - `AddCampaign.jsx`: 添加新捐赠活动的组件。
      - `DonationHome.jsx`: 显示捐赠活动的主页组件。
      - `DonationsManage.jsx`: 管理捐赠活动的组件。
      - `NftDisplay.jsx`: 显示NFT的组件。
      - `TxDialog.jsx`: 交易对话框组件。
    - `contracts/`: 合约交互相关代码。
    - `utils/`: 前端实用工具函数。
  - `public/`: 公共资产。
  - `package.json`: 项目元数据和依赖项。

## 设置说明

### 后端

1. **安装Go**: 确保您的系统上安装了Go。您可以从[golang.org](https://golang.org/dl/)下载。
2. **安装依赖项**: 导航到`backend-end/business/`和`backend-end/user/`目录并运行：
   ```sh
   go mod tidy
   ```
3. **运行后端**: 通过运行以下命令启动后端服务器：
   ```sh
   # 业务主服务
   cd backend-end/business
   $env:ENV="dev";go run main.go
   # 或
   $env:ENV="test";go run main.go
   
   # 用户服务
   cd backend-end/user
   $env:ENV="dev";go run main.go
   # 或
   $env:ENV="test";go run main.go
   ```

### 智能合约

1. **安装Node.js和npm**: 确保您的系统上安装了Node.js和npm。您可以从[nodejs.org](https://nodejs.org/)下载。
2. **安装依赖项**: 导航到`donations_manage/`目录并运行：
   ```sh
   npm install
   ```
3. **设置env-enc密码和键值对**: 通过以下命令设置env-enc密码和键值对(PRIVATE_KEY,SEPOLIA_RPC_URL,LINEA_SEPOLIA_RPC_URL,LINEA_API_KEY)：
   ```sh
   npx env-enc set-pwd
   npx env-enc set
   ```
4. **编译合约**: 通过运行以下命令编译Solidity合约：
   ```sh
   npx hardhat compile
   ```
5. **部署合约**: 通过运行以下命令将合约部署到本地以太坊网络（例如，Hardhat Network）：
   ```sh
   npx hardhat deploy --network sepolia
   ```

### 前端

1. **安装Node.js和npm**: 确保您的系统上安装了Node.js和npm。您可以从[nodejs.org](https://nodejs.org/)下载。
2. **安装依赖项**: 导航到`front-end/`目录并运行：
   ```sh
   npm install
   ```
3. **运行前端**: 通过运行以下命令启动前端开发服务器：
   ```sh
   npm run dev
   ```

## 主要组件

### 后端

- **业务主服务**: 处理捐赠活动和交易相关的业务逻辑,通过geth监听链上事件，同步数据至mysql，提高系统性能。
- **用户服务**: 处理用户注册，登录，颁发token，与token续期等
- **数据库**: 管理活动和捐赠，用户，权限数据
- **智能合约交互**: 与以太坊区块链上的智能合约交互，管理捐赠和活动。

### 智能合约

- **DonationsManageContract.sol**: 管理捐赠活动的创建和管理。
- **DonationsManageContractV2.sol**: 捐赠管理合约的升级版本，提供更多功能。
- **nft.sol**: 管理与捐赠相关的非同质化代币(NFTs)，捐赠者可以获得NFT作为奖励。
- **token.sol**: 管理用于捐赠的同质化代币，平台的原生代币。

### 前端

- **AddCampaign.jsx**: 用于创建新的捐赠活动的组件。
- **DonationHome.jsx**: 显示所有可用捐赠活动的主页组件。
- **DonationsManage.jsx**: 用于管理现有捐赠活动的组件。
- **NftDisplay.jsx**: 显示用户拥有的NFT的组件。
- **TxDialog.jsx**: 显示交易状态和确认的对话框组件。

## 功能特点

1. **创建捐赠活动**: 用户可以创建新的捐赠活动，设置目标金额、截止日期和描述。
2. **捐赠**: 用户可以向活动捐赠ETH或平台代币。
3. **NFT奖励**: 捐赠者可以根据捐赠金额获得不同等级的NFT奖励。
4. **活动管理**: 活动创建者可以管理其活动，查看捐赠情况和提取资金。
5. **用户认证**: 用户可以创建账户并登录，管理其捐赠历史和NFT收藏。

## 贡献

欢迎贡献！请按照以下步骤为项目做出贡献：
1. Fork仓库。
2. 为您的功能或错误修复创建一个新分支。
3. 进行更改并提交。
4. 将更改推送到您的fork。
5. 向主仓库提交拉取请求。

## 许可证

本项目根据MIT许可证授权。有关详细信息，请参阅[LICENSE](LICENSE)文件。
