require("@nomicfoundation/hardhat-toolbox");
require("hardhat-deploy");
require("@chainlink/env-enc").config();

const PRIVATE_KEY = process.env.PRIVATE_KEY
const SEPOLIA_RPC_URL = process.env.SEPOLIA_RPC_URL
const LINEA_SEPOLIA_RPC_URL = process.env.LINEA_SEPOLIA_RPC_URL

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
    }
  },
};
