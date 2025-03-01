require("@nomicfoundation/hardhat-toolbox");
require("hardhat-deploy");
require("@chainlink/env-enc").config();

const PRIVATE_KEY = process.env.PRIVATE_KEY
const SEPOLIA_RPC_URL = process.env.SEPOLIA_RPC_URL
const LINEA_SEPOLIA_RPC_URL = process.env.LINEA_SEPOLIA_RPC_URL
const LINEA_API_KEY = process.env.LINEA_API_KEY

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  namedAccounts: {
    deployer: {
      default: 0, // first account from mnemonic
    },
  },
  networks: {
    localhost: {
      chainId: 31337,
    },
    sepolia: {
      gasPrice: "auto",
      accounts : [PRIVATE_KEY],
      url : SEPOLIA_RPC_URL,
      chainId :11155111,
      blockConfirmations : 6,
    },
    lineaSepolia: {
      gasPrice: "auto",
      accounts: [process.env.PRIVATE_KEY],
      url: LINEA_SEPOLIA_RPC_URL,
      chainId: 59141,
      blockConfirmations: 6,
      verify: {
        etherscan: {
          // 显式指定验证参数
          apiUrl: "https://api.lineascan.build/api"
        }
      },
    }
  },
  etherscan: {
    // enabled: false // 禁用 Etherscan
    apiKey: {
      //linea的apikey主网，测试网都需要
      lineaSepolia: LINEA_API_KEY
    },
    customChains: [
      {
        network: "lineaSepolia",
        chainId: 59141,
        urls: {
          apiURL: "https://api-sepolia.lineascan.build/api", // 测试网专用接口
          browserURL: "https://explorer.sepolia.linea.build"
        }
      }
    ]
  },
  // 启用 Sourcify 验证（可选）
  sourcify: {
    enabled: true, // 启用 Sourcify 验证
    // force: true // 强制仅使用 Sourcify
  },
};
