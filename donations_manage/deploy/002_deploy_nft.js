module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  await deploy("NFT", {
    from: deployer,
    args: [
      "DonationNFT", // tokenName
      "DNFT" // tokenSymbol
    ],
    log: true,
  });
};

module.exports.tags = ["NFT"];
