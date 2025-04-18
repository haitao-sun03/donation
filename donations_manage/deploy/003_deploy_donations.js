module.exports = async ({ getNamedAccounts, deployments }) => {
  const { deployer } = await getNamedAccounts();
  const { get, save } = deployments;

  const token = await get("Token");
  const nft = await get("NFT");
  const DonationsManageContractFactory = await ethers.getContractFactory("DonationsManageContract");

  // 1 部署代理合约并初始化
  const donationsManageContractProxy = await hre.upgrades.deployProxy(
    DonationsManageContractFactory,
    [token.address, nft.address], // 初始化参数
    { 
      deployer: deployer,
      initializer: "initialize", // 指定初始化函数
      kind: 'uups' // 指定 UUPS 模式
    }
  );

  // 等待部署确认
  await donationsManageContractProxy.waitForDeployment();
  const proxyAddress = await donationsManageContractProxy.getAddress()

  // 获取逻辑合约地址
  const implAddress = await hre.upgrades.erc1967.getImplementationAddress(proxyAddress);

  // 打印关键信息
  console.log("================ Deployment Info ================");
  console.log("Proxy Contract Address:", proxyAddress);
  console.log("Logic Contract Address:", implAddress);
  console.log("=================================================");

  // 2 手动验证逻辑合约
  // console.log("Verifying logic contract...");
  // await hre.run("verify:verify", {
  //   address: implAddress,
  //   constructorArguments: [], // 无构造函数参数
  // });

  // 3 保存代理合约信息到部署记录
  await save("DonationsManageContractProxy", {
    address: proxyAddress,
    abi: DonationsManageContractFactory.interface.format('json'), // 必须包含 ABI
  });
};

module.exports.tags = ["DonationsManage"];
