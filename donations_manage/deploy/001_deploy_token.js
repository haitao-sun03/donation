module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  await deploy("Token", {
    from: deployer,
    args: [
      deployer, // initialOwner
      "AquaCoin", // _name
      "AquaCoin" // _symbol
    ],
    log: true,
  });
};

module.exports.tags = ["Token"];
