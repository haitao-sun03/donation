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
import { parseJwt, isTokenExpiringSoon } from "./utils/jwt";

function App() {
  const [isLoading, setIsLoading] = useState(true);
  const [connected, setConnected] = useState(false);
  const [account, setAccount] = useState(null);
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);
  const [balance, setBalance] = useState(0);
  const [symbol, setSymbol] = useState("");
  const [error, setError] = useState(null);
  const accountsChangedListenerAdded = useRef(false);
  const connectedAccountsRef = useRef([]); // 新增 ref 用于同步状态

  const [renewalTimer, setRenewalTimer] = useState(null);

  

  const signMessage = async (currentAccount) => {
    console.log("currentAccount:", currentAccount);
    let jwt = localStorage.getItem("jwt_" + currentAccount);
    let refreshJwt = localStorage.getItem("jwt_refresh_" + currentAccount);
    if (jwt && refreshJwt) {
      return;
    }

    let nonce;
    const requestBody = {
      currentAccount: currentAccount,
    };

    const dataLogin = await request("/user/auth/nonce", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify(requestBody),
    });

    if (dataLogin.code === 200) {
      nonce = dataLogin.data;
    } else {
      setError(dataLogin.msg || "Failed to Get nonce :" + dataLogin.msg);
    }

    let signature;
    console.log("======");
    const message = `please login,confirm the signature that do not cost gas\n nonce: ${nonce}`;
    try {
      signature = await window.ethereum.request({
        method: "personal_sign",
        params: [message, currentAccount],
      });
    } catch (error) {
      console.error("Signature error:", error);
    }
    console.log("!!!!!!");

    const requestBodyJwt = {
      currentAccount: currentAccount,
      signature: signature,
    };

    const dataJwt = await request("/user/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify(requestBodyJwt),
    });

    if (dataJwt.code === 200) {
      jwt = dataJwt.data[0];
      refreshJwt = dataJwt.data[1];
    } else {
      setError(dataJwt.msg || "Failed to Get jwt : " + dataJwt.msg);
    }
    console.log(`jwt_${currentAccount} :`, jwt);
    console.log(`jwt_refresh_${currentAccount} :`, refreshJwt);
    localStorage.setItem(`jwt_${currentAccount}`, jwt);
    localStorage.setItem(`jwt_refresh_${currentAccount}`, refreshJwt);
  };

  // 续期核心方法
  const silentRenew = async () => {
    if (!account) return false;

    const currentAccount = account.toLowerCase();
    const refreshToken = localStorage.getItem(`jwt_refresh_${currentAccount}`);

    try {
      const requestBodyJwt = {
        currentAccount: currentAccount,
        refreshJwt: refreshToken,
      };
      // 尝试使用refreshToken续期
      const { data } = await request("/user/auth/renewJwt", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${refreshToken}`,
        },
        body: JSON.stringify(requestBodyJwt),
      });

      localStorage.setItem(`jwt_${currentAccount}`, data);
      return true;
    } catch (error) {
      console.log("静默续期失败，降级到签名续期:", error);
      return signatureBasedRenew(currentAccount);
    }
  };

  // 签名续期降级方案
  const signatureBasedRenew = async (address) => {
    try {
      // 获取新nonce
      const { data: nonce } = await request("/user/auth/nonce", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({ currentAccount: address }),
      });
      const message = `please login,confirm the signature that do not cost gas\n nonce: ${nonce}`;
      // 触发签名（会弹窗）
      const signature = await window.ethereum.request({
        method: "personal_sign",
        params: [message, address],
      });

      // 获取新token
      const { data: tokens } = await request("/user/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({ currentAccount: address, signature }),
      });

      localStorage.setItem(`jwt_${address}`, tokens[0]);
      localStorage.setItem(`jwt_refresh_${address}`, tokens[1]);
      return true;
    } catch (error) {
      console.error("续期失败:", error);
      return false;
    }
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
        connectedAccountsRef.current = accounts;
        // 调用 signMessage 函数
        console.log("@@@");
        // await signMessage(address);
      } catch (error) {
        console.error("Connection error:", error);
        setAccount(null);
        setProvider(null);
        setSigner(null);
        setConnected(false);
        connectedAccountsRef.current = [];
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
        console.log("Attempting to remove:", `jwt_${addr}`); // 调试日志
        localStorage.removeItem(`jwt_${addr}`);
        localStorage.removeItem(`jwt_refresh_${addr}`);
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
          localStorage.removeItem(`jwt_refresh_${addr}`);
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
      connectedAccountsRef.current = [];
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

  // 新增的续期检测 useEffect
  useEffect(() => {
    if (!connected || !account) return;

    const setupRenewalCheck = async () => {
      const currentAccount = account.toLowerCase();

      // 5min内，等待用户处理登录签名
      const checkJWT = async () => {
        const maxWaitForSign = 5 * 60 * 1000; 
        let waitTime = 0;
        while (true) {
          const accessToken = localStorage.getItem(`jwt_${currentAccount}`);

          // 找到 token 立即返回
          if (accessToken) return accessToken;

          // 未找到时等待 1000ms（非阻塞）
          await new Promise((resolve) => setTimeout(resolve, 1000));
          waitTime += 1000;
          if(waitTime > maxWaitForSign) {
            return null;
          }
        }
      };

      // 使用方式
      const accessToken = await checkJWT();
      console.log(
        "setupRenewalCheck accessToken is : ",
        currentAccount,
        accessToken
      );
      // 防御性检测
      if (!accessToken) {
        console.warn("No access token found,renew not trigger!");
        return;
      }

      const payload = parseJwt(accessToken);
      if (!payload?.exp) {
        console.error("Invalid JWT structure:", payload);
        return;
      }
      // accessToken有效剩余时间
      const expiresInMs = payload.exp * 1000 - Date.now();
      // 5分钟前触发
      const renewThreshold = 5 * 60 * 1000;
      const delay = Math.max(expiresInMs - renewThreshold, 0);

      if (delay > 0) {
        console.log(`renew dalay time is ${delay/1000/60} min for ${currentAccount}`)
        const timer = setTimeout(async () => {
          // 续期
          silentRenew().then((success) => {
            if (success) setupRenewalCheck(); // 续期成功后，递归调用进行下次续期
          });
        }, delay);

        setRenewalTimer(timer);
      } else {
        // 立即续期
        const success = await silentRenew();
        if (success) await setupRenewalCheck();
      }
    };

    (async () => {
      await setupRenewalCheck();
    })();

    // 清理函数
    return () => {
      if (renewalTimer) {
        clearTimeout(renewalTimer);
        setRenewalTimer(null);
      }
    };
  }, [account, connected]);

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
