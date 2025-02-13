module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const token = await deployments.get("Token");
  const nft = await deployments.get("NFT");

  await deploy("DonationsManageContract", {
    from: deployer,
    args: [token.address, nft.address],
    log: true,
  });
};

module.exports.tags = ["DonationsManage"];
