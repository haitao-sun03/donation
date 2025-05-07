require("@nomicfoundation/hardhat-toolbox");
require("hardhat-deploy");
require("@chainlink/env-enc").config();
require('@nomicfoundation/hardhat-ethers');
require('@openzeppelin/hardhat-upgrades');

const PRIVATE_KEY = process.env.PRIVATE_KEY
const SEPOLIA_RPC_URL = process.env.SEPOLIA_RPC_URL
const LINEA_SEPOLIA_RPC_URL = process.env.LINEA_SEPOLIA_RPC_URL
const LINEA_API_KEY = process.env.LINEA_API_KEY
const SEPOLIA_API_KEY = process.env.SEPOLIA_API_KEY

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
      accounts: [PRIVATE_KEY],
      url: LINEA_SEPOLIA_RPC_URL,
      chainId: 59141,
      blockConfirmations: 6,
    }
  },
  etherscan: {
    // enabled: false // 禁用 Etherscan
    apiKey: {
      sepolia: SEPOLIA_API_KEY,
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
      },
      {
        network: "sepolia",
        chainId: 11155111,
        urls: {
          apiURL: "https://api-sepolia.etherscan.io/api", // 测试网专用接口
          browserURL: "https://sepolia.etherscan.io",
        }
      }
    ],
    timeout: 60000 // 默认是 30000（30秒）
  },
  // 启用 Sourcify 验证（可选）
  sourcify: {
    enabled: true, // 启用 Sourcify 验证
    // force: true // 强制仅使用 Sourcify
  },
};
