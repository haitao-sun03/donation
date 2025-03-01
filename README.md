# EquaSeed Donation

## Overview

Welcome to AquaSeed, a web3-based donation platform is a full-stack application that allows users to manage and donate to various campaigns. It consists of three main components:
1. **Frontend**: A React application that provides a user interface for interacting with the donation system.
2. **Backend**: A Go application that handles business logic, database interactions, and API requests.
3. **Smart Contracts**: Solidity contracts deployed on the Ethereum blockchain for managing donations and campaigns.

   - **DonationsManageContract**: 0x6Cf23549b2678027E4A1dDdC34c59c6255d05D13
   - **NFT**: 0x0A7Ceb2B9707123ceD34B0f6e444Cc5562bEA3DD
   - **Token**: 0x8D192A2f68700AEAeA4F5D3ADd666198e2047A81

    **contracts above have been deployed on linea sepolia testnet**

## Project Structure

- **backend-end/**: Contains the Go backend code.
  - `controllers/`: Handlers for HTTP requests.
  - `db/`: Database access objects.
  - `event/`: Event listeners.
  - `model/`: Data models.
  - `routers/`: HTTP routers.
  - `utils/`: Utility functions.
  - `vo/`: Value objects.
  - `config/`: Configuration files.
  - `abi/`: ABI files for smart contracts.
  - `main.go`: Entry point for the backend application.

- **donations_manage/**: Contains Solidity smart contracts and deployment scripts.
  - `contracts/`: Solidity contract files.
  - `deploy/`: Deployment scripts.
  - `deployments/`: Deployment artifacts.
  - `hardhat.config.js`: Hardhat configuration file.
  - `README.md`: Existing README for smart contracts.

- **front-end/**: Contains the React frontend code.
  - `src/`: Source code for the React application.
  - `public/`: Public assets.
  - `package.json`: Project metadata and dependencies.
  - `package-lock.json`: Dependency lock file.

## Setup Instructions

### Backend

1. **Install Go**: Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).
2. **Install Dependencies**: Navigate to the `backend-end/` directory and run:
   ```sh
   go mod tidy
   ```
3. **Run the Backend**: Start the backend server by running:
   ```sh
   $env:ENV="dev";go run main.go
   $env:ENV="test";go run main.go
   ```

### Smart Contracts

1. **Install Node.js and npm**: Ensure you have Node.js and npm installed on your system. You can download them from [nodejs.org](https://nodejs.org/).
2. **Install Dependencies**: Navigate to the `donations_manage/` directory and run:
   ```sh
   npm install
   ```
3. **Set env-enc password and key-value**: Set env-enc password and key-value(PRIVATE_KEY,SEPOLIA_RPC_URL,LINEA_SEPOLIA_RPC_URL,LINEA_API_KEY) By commands below:
   ```sh
   npx env-enc set-pwd
   npx env-enc set
   ```
4. **Compile Contracts**: Compile the Solidity contracts by running:
   ```sh
   npx hardhat compile
   ```
5. **Deploy Contracts**: Deploy the contracts to a local Ethereum network (e.g., Hardhat Network) by running:
   ```sh
   npx hardhat deploy --network lineaSepolia
   ```

### Frontend

1. **Install Node.js and npm**: Ensure you have Node.js and npm installed on your system. You can download them from [nodejs.org](https://nodejs.org/).
2. **Install Dependencies**: Navigate to the `front-end/` directory and run:
   ```sh
   npm install
   ```
3. **Run the Frontend**: Start the frontend development server by running:
   ```sh
   npm start
   ```

## Key Components

### Backend

- **Controllers**: Handle HTTP requests and interact with the database and smart contracts.
- **Database**: Manages campaign and donation data.
- **Smart Contracts**: Interact with the Ethereum blockchain for managing donations and campaigns.

### Smart Contracts

- **DonationsManageContract.sol**: Manages the creation and management of donation campaigns.
- **NFT.sol**: Manages non-fungible tokens (NFTs) associated with donations.
- **Token.sol**: Manages fungible tokens used for donations.

### Frontend

- **ConnectWallet.jsx**: Component for connecting to a user's Ethereum wallet.
- **DonationHome.jsx**: Component for displaying donation campaigns.
- **DonationsManage.jsx**: Component for managing donation campaigns.

## Contributing

Contributions are welcome! Please follow these steps to contribute to the project:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your fork.
5. Open a pull request to the main repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
