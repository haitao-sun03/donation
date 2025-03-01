import React, { useState, useEffect, useRef } from "react";
import { ethers } from "ethers";
import ConnectWallet from "./components/ConnectWallet";
import DonationsManage from "./components/DonationsManage";
import DonationHome from "./components/DonationHome";
import "./App.css";
import {
  AppBar,
  Toolbar,
  Typography,
  Box,
  Divider,
  Button,
} from "@mui/material";
import AccountBalanceWalletIcon from "@mui/icons-material/AccountBalanceWallet";
import CurrencyExchangeIcon from "@mui/icons-material/CurrencyExchange";
import { request } from "./utils/api";

function App() {
  const [isLoading, setIsLoading] = useState(true);
  const [connected, setConnected] = useState(false);
  const [account, setAccount] = useState(null);
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);
  const [balance, setBalance] = useState(0);
  const [symbol, setSymbol] = useState("");
  const accountsChangedListenerAdded = useRef(false);
  const connectedAccountsRef = useRef([]); // 新增 ref 用于同步状态

  const signMessage = async (currentAccount) => {
    console.log("currentAccount:", currentAccount);
    let jwt = localStorage.getItem("jwt_" + currentAccount);
    if (jwt) {
      return;
    }

    let nonce;
    const requestBody = {
      currentAccount: currentAccount,
    };
    const response = await fetch("/api/auth/nonce", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify(requestBody),
    });

    if (!response.ok) {
      const errorText = await response.text();
      console.error("Response not ok:", response.status, errorText);
      throw new Error(`Get nonce error! status: ${response.status}`);
    }

    const data = await response.json();
    console.log("Raw response data:", data);

    if (data.code === 200) {
      nonce = data.data;
    } else {
      setError(data.msg || "Failed to Get nonce : " + data.msg);
    }
    let signature;
    try {
      signature = await window.ethereum.request({
        method: "personal_sign",
        params: [nonce, currentAccount],
      });
    } catch (error) {
      console.error("Signature error:", error);
    }

    const requestBodyJwt = {
      currentAccount: currentAccount,
      signature: signature,
    };

    const dataJwt = await request("/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify(requestBodyJwt),
    });

    if (dataJwt.code === 200) {
      jwt = dataJwt.data;
    } else {
      setError(dataJwt.msg || "Failed to Get jwt : " + dataJwt.msg);
    }
    console.log("jwt_" + currentAccount.tol + " :", jwt);
    localStorage.setItem("jwt_" + currentAccount, jwt);
  };

  // 定义回调函数，用于接收余额数据
  const handleBalanceUpdate = (newBalance) => {
    setBalance(newBalance);
  };

  const handleSymbolUpdate = (newSymbol) => {
    setSymbol(newSymbol);
  };

  const handleConnect = async () => {
    console.log("handleConnect is called");
    if (window.ethereum) {
      try {
        // 请求用户授权连接钱包
        const accounts = await window.ethereum.request({
          method: "eth_requestAccounts",
        });

        // 初始化 provider 和 signer
        const provider = new ethers.BrowserProvider(window.ethereum);
        const signer = await provider.getSigner();
        const address = await signer.getAddress();

        // 更新状态
        setAccount(address);
        setProvider(provider);
        setSigner(signer);
        setConnected(true);
        console.log("Connected accounts in handleConnect:", accounts);
        // 初始化 connectedAccounts
        connectedAccountsRef.current = accounts
        // 调用 signMessage 函数
        console.log("@@@");
        // await signMessage(address);
      } catch (error) {
        console.error("Connection error:", error);
        setAccount(null);
        setProvider(null);
        setSigner(null);
        setConnected(false);
        connectedAccountsRef.current = []
      }
    } else {
      console.error("Ethereum provider not found");
      alert("Please install MetaMask or another Ethereum wallet.");
    }
  };

  // 修改 handleAccountsChanged 函数
  const handleAccountsChanged = async (accounts) => {
    console.log("handleAccountsChanged is called");
    console.log("current accounts:", accounts);
    console.log("before accounts:", connectedAccountsRef.current);
    const prevAccounts = connectedAccountsRef.current; // 从 ref 获取最新值

    // 检测断开的账户
    const disconnectedAccounts = prevAccounts.filter(
      (addr) => !accounts.includes(addr)
    );
    console.log("Disconnected accounts:", disconnectedAccounts);
    if (disconnectedAccounts.length > 0) {
      disconnectedAccounts.forEach((addr) => {
        console.log('Attempting to remove:', `jwt_${addr}`); // 调试日志
        localStorage.removeItem(`jwt_${addr}`);
        console.log(`Removed JWT for ${addr}`);
      });
    }

    // 更新 connectedAccounts
    connectedAccountsRef.current = accounts;
    console.log(
      "setConnectedAccounts in handleAccountsChanged",
      connectedAccountsRef.current
    );
    try {
      if (accounts.length > 0) {
        // 等待 provider 准备就绪
        await new Promise((resolve) => setTimeout(resolve, 100));

        // 只有当账户真的改变时才更新状态
        const web3Provider = new ethers.BrowserProvider(window.ethereum);
        const newSigner = await web3Provider.getSigner();
        const address = await newSigner.getAddress();

        // 检查账户是否真的改变了
        if (address !== account) {
          setAccount(address);
          setProvider(web3Provider);
          setSigner(newSigner);
          setConnected(true);

          // 调用 signMessage 函数
          await signMessage(address.toLowerCase());
        }
      } else {
        // 断开连接时清除状态
        console.log("All accounts disconnected");
        setAccount(null);
        setProvider(null);
        setSigner(null);
        setConnected(false);
        connectedAccountsRef.current.forEach((addr) => {
          localStorage.removeItem(`jwt_${addr}`);
        });
      }
    } catch (error) {
      console.error("Error handling account change:", error);
      // 出错时也清除状态
      setAccount(null);
      setProvider(null);
      setSigner(null);
      setConnected(false);
      connectedAccountsRef.current.forEach((addr) => {
        localStorage.removeItem(`jwt_${addr}`);
      });
      connectedAccountsRef.current = []

    }
  };

  // 修改 init 函数
  useEffect(() => {
    const init = async () => {
      if (window.ethereum) {
        try {
          const provider = new ethers.BrowserProvider(window.ethereum);
          // const accounts = await provider.listAccounts();
          const accounts = await window.ethereum.request({
            method: "eth_accounts",
          });

          if (accounts.length > 0) {
            const signer = await provider.getSigner();
            const address = await signer.getAddress();

            // 只在初始化时设置一次
            setAccount(address);
            setProvider(provider);
            setSigner(signer);
            setConnected(true);
            connectedAccountsRef.current = accounts; // 初始化时设置 ref
            console.log("Initialized connectedAccounts in init:", accounts);
          } else {
            connectedAccountsRef.current = [];
          }
        } catch (error) {
          console.error("Error initializing wallet:", error);
          setConnectedAccounts([]);
        } finally {
          setIsLoading(false);
        }

        // 添加事件监听器
        if (!accountsChangedListenerAdded.current) {
          window.ethereum.on("accountsChanged", handleAccountsChanged);
          console.log("accountsChanged is added");
          accountsChangedListenerAdded.current = true;
        }
        window.ethereum.on("chainChanged", () => {
          window.location.reload();
        });
      } else {
        setIsLoading(false);
      }
    };

    init();

    // 清理事件监听器
    return () => {
      if (window.ethereum) {
        window.ethereum.removeListener(
          "accountsChanged",
          handleAccountsChanged
        );
        window.ethereum.removeListener("chainChanged", () => {
          window.location.reload();
        });
      }
    };
  }, []); // 保持空依赖数组

  if (isLoading) {
    return (
      <div className="app-loading">
        <div className="loading-spinner"></div>
        <p>Checking wallet connection...</p>
      </div>
    );
  }

  return (
    <div className="app">
      <AppBar position="sticky" sx={{ backgroundColor: "#1e1b4b" }}>
        <Toolbar
          sx={{
            display: "flex",
            justifyContent: "space-between",
            gap: 2, // 添加元素间统一间距
            flexWrap: "wrap", // 小屏幕自动换行
          }}
        >
          {/* 左边区块 */}
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              gap: 1.5,
              minWidth: 200,
            }}
          >
            <AccountBalanceWalletIcon
              fontSize="small"
              sx={{ color: "#a5b4fc" }}
            />
            {connected ? (
              <Typography variant="subtitle1" noWrap>
                {`${account.slice(0, 6)}...${account.slice(-4)}`}
              </Typography>
            ) : (
              <Button
                variant="contained"
                onClick={handleConnect} // 绑定点击事件
                sx={{
                  textTransform: "none", // 禁用大写转换
                  backgroundColor: "#6366f1", // 设置按钮颜色
                  "&:hover": {
                    backgroundColor: "#4f46e5", // 设置悬停颜色
                  },
                }}
              >
                Connect Your Wallet
              </Button>
            )}
          </Box>

          {/* 右边区块 */}
          {connected && (
            <Box
              sx={{
                display: "flex",
                alignItems: "center",
                gap: 2,
                color: "#c7d2fe",
              }}
            >
              {/* 添加优雅的分隔线 */}
              <Divider
                orientation="vertical"
                flexItem
                sx={{ bgcolor: "#4f46e5" }}
              />

              <Box sx={{ display: "flex", alignItems: "center", gap: 1 }}>
                <CurrencyExchangeIcon fontSize="small" />
                <Typography variant="subtitle2" sx={{ fontWeight: 500 }}>
                  {balance}
                  <Box component="span" sx={{ color: "#818cf8", ml: 0.5 }}>
                    {symbol}
                  </Box>
                </Typography>
              </Box>
            </Box>
          )}
        </Toolbar>
      </AppBar>
      <main>
        {connected ? (
          <DonationsManage
            provider={provider}
            signer={signer}
            onBalanceUpdate={handleBalanceUpdate}
            onSymbolUpdate={handleSymbolUpdate}
          />
        ) : (
          <DonationHome />
        )}
      </main>
    </div>
  );
}

export default App;
