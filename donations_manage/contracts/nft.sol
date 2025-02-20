// SPDX-License-Identifier: MIT
// Compatible with OpenZeppelin Contracts ^5.0.0
pragma solidity ^0.8.22;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NFT is ERC721, ERC721Enumerable, ERC721URIStorage, ERC721Burnable, Ownable {
    uint256 private _nextTokenId;

    mapping(string=>string) public levelToMetaDataURI;

    // 定义一个结构体来存储键值对
    struct LevelMetaDataURI {
        string level;
        string uri;
    }

    // 事件定义，使用数组来传递映射内容
    event NFTDeployed(LevelMetaDataURI[] levelToMetaDataURIs);
    event NFTMinted(address to,uint256 tokenId,string level);

    constructor(string memory tokenName,string memory tokenSymol)
        ERC721(tokenName, tokenSymol)
        Ownable(msg.sender) {
            levelToMetaDataURI["Diamond"] = "ipfs://QmQwZbQNvfY9dHfW1UsH2zGPXuFsgPUNnKgBVs5LWHxokf";
            levelToMetaDataURI["Gold"] = "ipfs://QmULKjJgJJU42po2W6EWRQ68KH1sYBfTL5HeGtjo6QKiBc";
            levelToMetaDataURI["Silver"] = "ipfs://QmVHysuCAvFFKsmZvAFGzrtbsc5MJ3x58mrQR54PcJC7xg";
        }
    
    // 添加一个函数来触发NFTDeployed事件
    function triggerNFTDeployedEvent() public onlyOwner {
         // 将mapping转换为数组
        LevelMetaDataURI[] memory uris = new LevelMetaDataURI[](3);
        uris[0] = LevelMetaDataURI("Diamond", levelToMetaDataURI["Diamond"]);
        uris[1] = LevelMetaDataURI("Gold", levelToMetaDataURI["Gold"]);
        uris[2] = LevelMetaDataURI("Silver", levelToMetaDataURI["Silver"]);
        // 发送事件
        emit NFTDeployed(uris);
    }

    function safeMint(address to,string memory level) public {
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
        _setTokenURI(tokenId,  levelToMetaDataURI[level]);
        emit NFTMinted(to, tokenId, level);
    }

    // The following functions are overrides required by Solidity.

    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721, ERC721Enumerable)
        returns (address) {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(address account, uint128 value)
        internal
        override(ERC721, ERC721Enumerable) {
        super._increaseBalance(account, value);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory) {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721Enumerable, ERC721URIStorage)
        returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}