module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const token = await deployments.get("Token");
  const nft = await deployments.get("NFT");

  const donationsManageContract = await deploy("DonationsManageContract", {
    from: deployer,
    args: [token.address, nft.address],
    log: true,
  });
  
  console.log("DonationsManageContract verify start");
  await hre.run("verify:verify", {
    address: donationsManageContract.address,
    constructorArguments: [
      token.address,
      nft.address,
    ],
  });
  console.log("DonationsManageContract verify success");
};

module.exports.tags = ["DonationsManage"];
