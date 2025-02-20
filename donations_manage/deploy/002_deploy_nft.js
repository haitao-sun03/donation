module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const deployment = await deploy("NFT", {
    from: deployer,
    args: [
      "EcoSeed", // tokenName
      "EcoSeed" // tokenSymbol
    ],
    log: true,
  });
  // 等待一段时间，让后端服务有时间更新和启动
  await new Promise(resolve => setTimeout(resolve, 10000)); // 等待30秒，可以根据需要调整时间

  // 获取合约实例后，调用triggerNFTDeployedEvent方法
  const nftContract = await ethers.getContractAt("NFT", deployment.address);
  await nftContract.triggerNFTDeployedEvent();
  console.log("NFTDeployed event has been triggered.");
};

module.exports.tags = ["NFT"];
