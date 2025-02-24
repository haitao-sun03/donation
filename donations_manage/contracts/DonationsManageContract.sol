// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {NFT} from "./nft.sol";
import "@chainlink/contracts/src/v0.8/automation/interfaces/AutomationCompatibleInterface.sol";

/**
 * @title donates contract
 * @author Zhang haitao
 */
contract DonationsManageContract is AutomationCompatibleInterface {
    address public owner;

    /// 定义活动性质的枚举
    enum CampaignNature {
        Charity,
        Education,
        Health,
        Environment,
        Other
    }

    /// 定义受益对象的枚举
    enum Beneficiary {
        Individuals,
        NonProfitOrganizations,
        Communities,
        Other
    }

    /// 定义活动状态的枚举
    enum CampaignStatus {
        Pending,    /// 活动未开始
        Active,     /// 活动正在进行
        Completed,   /// 活动已关闭（可能是达成目标或时间到期）
        Cancelled   /// 活动已取消
    }

    ///募捐活动实体
    struct Campaign {
        string title;
        uint goal;
        uint totalDonated;
        uint startTime;
        uint endTime;
        CampaignNature nature; // 活动性质
        Beneficiary beneficiary; // 受益对象
        string purpose; // 活动目的
        string expectedImpact; // 预期影响
        mapping(address => uint) donations;
        CampaignStatus status;
        address starter;
    }

    IERC20 token;
    NFT nft;

    mapping(uint => Campaign) public campaigns;
    /// 存储活动ID，用于活动状态更新
    uint[] public campaignsToCheck;
    uint public campaignCount;

    modifier onlyStarter(uint campaignId) {
        Campaign storage campaign =  campaigns[campaignId];
        require(msg.sender == campaign.starter,"not starter");
        _;
    }

    event StartCampaign(
        uint indexed id,
        address starter,
        string title,
        uint goal,
        uint startTime,
        uint endTime,
        CampaignNature nature, // 新增字段
        Beneficiary beneficiary, // 新增字段
        string purpose, // 活动目的
        string expectedImpact, // 预期影响
        CampaignStatus status
    );
    event Donate(uint indexed id,address donater,uint amount);
    event CancellCampaign(uint indexed id,address caller,uint time);
    event CompletedCampaign(uint indexed id,address caller,uint time);
    event Refund(uint indexed id,address refunder,uint refundAmount,uint time);
    event Withdraw(uint indexed id,address withdrawer,uint withdrawAmount,uint time);
    event ActiveCampaign(uint indexed id,address caller,uint time);

    constructor(address _token,address _nft) {
        token = IERC20(_token);
        nft = NFT(_nft);
    }

    function createCampaign(
        string memory _title,
        uint _goal,
        uint _startTime,
        uint _endTime,
        CampaignNature _nature, // 新增参数
        Beneficiary _beneficiary, // 新增参数
        string memory _purpose, // 新增参数
        string memory _expectedImpact // 新增参数
    ) external {
        require(bytes(_title).length > 0, "Title cannot be empty");
        require(_goal > 0,"Invalid goal");
        require(_startTime < _endTime,"_startTime is later than _endTime");
        require(block.timestamp < _endTime,"Invalid endTime");
        // 充当campaignId
        campaignCount++;
        Campaign storage campaign = campaigns[campaignCount];
        campaign.title = _title;
        campaign.goal = _goal;
        campaign.startTime = _startTime;
        campaign.endTime = _endTime;
        campaign.nature = _nature; // 新增赋值
        campaign.beneficiary = _beneficiary; // 新增赋值
        campaign.purpose = _purpose; // 新增赋值
        campaign.expectedImpact = _expectedImpact; // 新增赋值
        campaign.status = block.timestamp > _startTime ? CampaignStatus.Active : CampaignStatus.Pending;
        campaign.starter = msg.sender;
        campaignsToCheck.push(campaignCount);
        emit StartCampaign(campaignCount, msg.sender, _title, _goal, _startTime, _endTime, _nature, _beneficiary, _purpose, _expectedImpact, campaign.status);
    }

    /// 手动将活动状态变为active,有些网络不支持chainlink automation，需要手动触发
    function activeCampaign(uint _campaignId) external onlyStarter(_campaignId) {
        Campaign storage compaign = campaigns[_campaignId];
        require(compaign.status == CampaignStatus.Pending,"the campaign is not pending,can not active");
        compaign.status = CampaignStatus.Active;
        emit ActiveCampaign(_campaignId, msg.sender, block.timestamp);
    }

    /// starter主动取消活动，另一种情况为到了endTime，goal没有达到，会自动置为cancelled
    function cancellCampaign(uint _campaignId) external onlyStarter(_campaignId) {
        Campaign storage compaign = campaigns[_campaignId];
        require(compaign.status != CampaignStatus.Completed,"the campaign is completed,can not cancell");
        compaign.status = CampaignStatus.Cancelled;
        emit CancellCampaign(_campaignId,msg.sender,block.timestamp);
    }

    /// 若在活动时间内，提前reach goal，starter可手动关闭
    function completedCampaign(uint _campaignId) external onlyStarter(_campaignId) {
        Campaign storage compaign = campaigns[_campaignId];
        require(compaign.status == CampaignStatus.Active ,"the campaign is not active ,can not completed");
        require(compaign.totalDonated >= compaign.goal ,"goal is not reach");
        compaign.status = CampaignStatus.Completed;
        emit CompletedCampaign(_campaignId,msg.sender,block.timestamp);
    }

    function checkUpkeep(bytes calldata) external view override returns  (bool upkeepNeeded, bytes memory performData) {
        upkeepNeeded = false;
        uint campaignsToCheckLength = campaignsToCheck.length;
        uint[] memory campaignsToProcess = new uint[](campaignsToCheckLength);
        uint processCount = 0;

        for (uint i = 0; i < campaignsToCheckLength; i++) {
            Campaign storage campaign = campaigns[campaignsToCheck[i]];
            if (campaign.status == CampaignStatus.Cancelled || campaign.status == CampaignStatus.Completed) {
                continue;
            }

            /**
            * 3种情况需要自动控制募捐活动状态：
            *   达到了设置的活动时间内，活动状态为Pending，需要更改为Active
            *   活动开始后，只要达到目标，活动状态更改为Completed
            *   超出活动时间后，还没有达到目标，活动需要更改为Cancelled
            */
            if((block.timestamp > campaign.startTime && block.timestamp < campaign.endTime && campaign.status == CampaignStatus.Pending)
            || (block.timestamp > campaign.startTime && campaign.totalDonated >= campaign.goal)
            || (block.timestamp >= campaign.endTime && campaign.totalDonated < campaign.goal)) {
                campaignsToProcess[processCount] = campaignsToCheck[i];
                processCount++;
            }
        }
        
        if (processCount > 0) {
            upkeepNeeded = true;
            performData = abi.encode(sliceArray(campaignsToProcess, processCount));
        } 
    }

    function sliceArray(uint[] memory arr, uint length) private pure returns (uint[] memory) {
        uint[] memory result = new uint[](length);
        for (uint i = 0; i < length; i++) {
            result[i] = arr[i];
        }
        return result;
    }

    function performUpkeep(bytes calldata performData) external override {
       uint[] memory campaignsToProcess = abi.decode(performData, (uint[]));
        for (uint i = 0; i < campaignsToProcess.length; i++) {
            uint campaignId = campaignsToProcess[i];
            Campaign storage campaign = campaigns[campaignId];
            // 若在自动运行过程中，管理员手动处理取消或关闭
            if (campaign.status == CampaignStatus.Cancelled || campaign.status == CampaignStatus.Completed) {
                continue ;
            }

            if(block.timestamp > campaign.startTime && block.timestamp < campaign.endTime && campaign.status == CampaignStatus.Pending) {
                campaign.status = CampaignStatus.Active;
                emit ActiveCampaign(campaignId, msg.sender, block.timestamp);
            }

            if (block.timestamp >= campaign.endTime && campaign.totalDonated < campaign.goal) {
                campaign.status = CampaignStatus.Cancelled;
                emit CancellCampaign(campaignId,msg.sender,block.timestamp);
                continue ;
            } 
            if (block.timestamp > campaign.startTime && campaign.totalDonated >= campaign.goal) {
                campaign.status = CampaignStatus.Completed;
                emit CompletedCampaign(campaignId, msg.sender, block.timestamp);
                continue ;
            }
        }
    }

    /// 活动取消后，捐赠者可退款
    function refundToken(uint _campaignId) external {
        Campaign storage compaign = campaigns[_campaignId];
        require(compaign.status == CampaignStatus.Cancelled,"campain not cancelled,can not refund");
        mapping(address => uint) storage donation = compaign.donations;
        uint refund = donation[msg.sender];
        if (refund != 0) {
            donation[msg.sender] = 0;
            compaign.totalDonated -= refund;
            bool success = token.transfer(msg.sender, refund);
            require(success,"refund transfer failed");
            emit Refund(_campaignId,msg.sender,refund,block.timestamp);
        }
    }

    /// 捐赠者调用，用于给某活动捐赠
    function donate(uint _campaignId,uint _amount) external {
        Campaign storage campaign = campaigns[_campaignId];
        require(campaign.status == CampaignStatus.Active,"the campaign is not start or ended");
        
        campaign.donations[msg.sender] += _amount;
        campaign.totalDonated += _amount;

        bool success = token.transferFrom(msg.sender, address(this), _amount);
        require(success,"donate transfer fail");

        emit Donate(_campaignId, msg.sender, _amount);
    }

    /// 当活动达到目标，处于Completed状态，可取款
    function withdraw(uint _campaignId) external onlyStarter(_campaignId) {
        Campaign storage campaign = campaigns[_campaignId];
        require(campaign.status == CampaignStatus.Completed,"the campaign is not Completed");
        uint donated = campaign.totalDonated;
        campaign.totalDonated = 0;
        bool success = token.transfer(msg.sender, donated);
        require(success,"withdraw transfer failed");

        emit Withdraw(_campaignId,msg.sender,donated,block.timestamp);
    }

    function getCampaignDetails(uint _campaignId) external view returns (string memory title, uint goal, uint totalDonated,uint startTime,uint endTime,CampaignStatus status,address starter) {
        Campaign storage campaign = campaigns[_campaignId];
        title = campaign.title;
        goal = campaign.goal;
        totalDonated = campaign.totalDonated;
        startTime = campaign.startTime;
        endTime = campaign.endTime;
        status = campaign.status;
        starter = campaign.starter;
    }
}
