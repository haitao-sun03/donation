// import React, { useState, useEffect } from "react";
// import { ethers } from "ethers";
// import ConnectWallet from "./components/ConnectWallet";
// import DonationsManage from "./components/DonationsManage";
// import DonationHome from "./components/DonationHome";
// import "./App.css";
// import { AppBar, Toolbar, Typography, Box, Divider,Button  } from "@mui/material";
// import AccountBalanceWalletIcon from "@mui/icons-material/AccountBalanceWallet";
// import CurrencyExchangeIcon from "@mui/icons-material/CurrencyExchange";

// function App() {
//   const [isLoading, setIsLoading] = useState(true);
//   const [connected, setConnected] = useState(false);
//   const [account, setAccount] = useState(null);
//   const [provider, setProvider] = useState(null);
//   const [signer, setSigner] = useState(null);
//   const [balance, setBalance] = useState(0);
//   const [symbol, setSymbol] = useState("");

//   // 定义回调函数，用于接收余额数据
//   const handleBalanceUpdate = (newBalance) => {
//     setBalance(newBalance);
//   };

//   const handleSymbolUpdate = (newSymbol) => {
//     setSymbol(newSymbol);
//   };

//   const handleConnect = async (wallet) => {
//     if (!wallet) {
//       setAccount(null);
//       setProvider(null);
//       setSigner(null);
//       setConnected(false);
//       return;
//     }

//     try {
//       const address = await wallet.signer.getAddress();
//       const newProvider = new ethers.BrowserProvider(window.ethereum);
//       const newSigner = await newProvider.getSigner();

//       setAccount(address);
//       setProvider(newProvider);
//       setSigner(newSigner);
//       setConnected(true);
//     } catch (error) {
//       console.error("Connection error:", error);
//       setAccount(null);
//       setProvider(null);
//       setSigner(null);
//       setConnected(false);
//     }
//   };

//   // 修改 handleAccountsChanged 函数
//   const handleAccountsChanged = async (accounts) => {
//     try {
//       if (accounts.length > 0) {
//         // 等待 provider 准备就绪
//         await new Promise((resolve) => setTimeout(resolve, 100));

//         // 只有当账户真的改变时才更新状态
//         const web3Provider = new ethers.BrowserProvider(window.ethereum);
//         const newSigner = await web3Provider.getSigner();
//         const address = await newSigner.getAddress();

//         // 检查账户是否真的改变了
//         if (address !== account) {
//           setAccount(address);
//           setProvider(web3Provider);
//           setSigner(newSigner);
//           setConnected(true);
//         }
//       } else {
//         // 断开连接时清除状态
//         setAccount(null);
//         setProvider(null);
//         setSigner(null);
//         setConnected(false);
//       }
//     } catch (error) {
//       console.error("Error handling account change:", error);
//       // 出错时也清除状态
//       setAccount(null);
//       setProvider(null);
//       setSigner(null);
//       setConnected(false);
//     }
//   };

//   // 修改 init 函数
//   useEffect(() => {
//     const init = async () => {
//       if (window.ethereum) {
//         try {
//           const provider = new ethers.BrowserProvider(window.ethereum);
//           const accounts = await provider.listAccounts();

//           if (accounts.length > 0) {
//             const signer = await provider.getSigner();
//             const address = await signer.getAddress();

//             // 只在初始化时设置一次
//             setAccount(address);
//             setProvider(provider);
//             setSigner(signer);
//             setConnected(true);
//           }
//         } catch (error) {
//           console.error("Error initializing wallet:", error);
//         } finally {
//           setIsLoading(false);
//         }

//         // 添加事件监听器
//         window.ethereum.on("accountsChanged", handleAccountsChanged);
//         window.ethereum.on("chainChanged", () => {
//           window.location.reload();
//         });
//       } else {
//         setIsLoading(false);
//       }
//     };

//     init();

//     // 清理事件监听器
//     return () => {
//       if (window.ethereum) {
//         window.ethereum.removeListener(
//           "accountsChanged",
//           handleAccountsChanged
//         );
//         window.ethereum.removeListener("chainChanged", () => {
//           window.location.reload();
//         });
//       }
//     };
//   }, []); // 保持空依赖数组

//   if (isLoading) {
//     return (
//       <div className="app-loading">
//         <div className="loading-spinner"></div>
//         <p>Checking wallet connection...</p>
//       </div>
//     );
//   }

//   return (
//     <div className="app">
//       <AppBar position="sticky" sx={{ backgroundColor: "#1e1b4b" }}>
//         <Toolbar
//           sx={{
//             display: "flex",
//             justifyContent: "space-between",
//             gap: 2, // 添加元素间统一间距
//             flexWrap: "wrap", // 小屏幕自动换行
//           }}
//         >
//           {/* 左边区块 */}
//           <Box
//             sx={{
//               display: "flex",
//               alignItems: "center",
//               gap: 1.5,
//               minWidth: 200,
//             }}
//           >
//             <AccountBalanceWalletIcon
//               fontSize="small"
//               sx={{ color: "#a5b4fc" }}
//             />
//             {/* <Typography variant="subtitle1" noWrap>
//               {connected ? `${account.slice(0, 6)}...${account.slice(-4)}` : "Connect Your Wallet"}
//             </Typography> */}

//             {connected ? (
//               <Typography variant="subtitle1" noWrap>
//                 {`${account.slice(0, 6)}...${account.slice(-4)}`}
//               </Typography>
//             ) : (
//               <Button
//                 variant="contained"
//                 onClick={handleConnect} // 绑定点击事件
//                 sx={{
//                   textTransform: "none", // 禁用大写转换
//                   backgroundColor: "#6366f1", // 设置按钮颜色
//                   "&:hover": {
//                     backgroundColor: "#4f46e5", // 设置悬停颜色
//                   },
//                 }}
//               >
//                 Connect Your Wallet
//               </Button>
//             )}
//           </Box>

//           {/* 右边区块 */}
//           {connected && (
//             <Box
//               sx={{
//                 display: "flex",
//                 alignItems: "center",
//                 gap: 2,
//                 color: "#c7d2fe",
//               }}
//             >
//               {/* 添加优雅的分隔线 */}
//               <Divider
//                 orientation="vertical"
//                 flexItem
//                 sx={{ bgcolor: "#4f46e5" }}
//               />

//               <Box sx={{ display: "flex", alignItems: "center", gap: 1 }}>
//                 <CurrencyExchangeIcon fontSize="small" />
//                 <Typography variant="subtitle2" sx={{ fontWeight: 500 }}>
//                   {balance}
//                   <Box component="span" sx={{ color: "#818cf8", ml: 0.5 }}>
//                     {symbol}
//                   </Box>
//                 </Typography>
//               </Box>
//             </Box>
//           )}
//         </Toolbar>
//       </AppBar>
//       <main>
//         {connected ? (
//           <DonationsManage
//             provider={provider}
//             signer={signer}
//             onBalanceUpdate={handleBalanceUpdate}
//             onSymbolUpdate={handleSymbolUpdate}
//           />
//         ) : (
//           <DonationHome />
//         )}
//       </main>
//     </div>
//   );
// }

// export default App;
import React, { useState, useEffect } from "react";
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

function App() {
  const [isLoading, setIsLoading] = useState(true);
  const [connected, setConnected] = useState(false);
  const [account, setAccount] = useState(null);
  const [provider, setProvider] = useState(null);
  const [signer, setSigner] = useState(null);
  const [balance, setBalance] = useState(0);
  const [symbol, setSymbol] = useState("");

  // 定义回调函数，用于接收余额数据
  const handleBalanceUpdate = (newBalance) => {
    setBalance(newBalance);
  };

  const handleSymbolUpdate = (newSymbol) => {
    setSymbol(newSymbol);
  };

  const handleConnect = async () => {
    if (window.ethereum) {
      try {
        // 请求用户授权连接钱包
        await window.ethereum.request({ method: "eth_requestAccounts" });

        // 初始化 provider 和 signer
        const provider = new ethers.BrowserProvider(window.ethereum);
        const signer = await provider.getSigner();
        const address = await signer.getAddress();

        // 更新状态
        setAccount(address);
        setProvider(provider);
        setSigner(signer);
        setConnected(true);
      } catch (error) {
        console.error("Connection error:", error);
        setAccount(null);
        setProvider(null);
        setSigner(null);
        setConnected(false);
      }
    } else {
      console.error("Ethereum provider not found");
      alert("Please install MetaMask or another Ethereum wallet.");
    }
  };

  // 修改 handleAccountsChanged 函数
  const handleAccountsChanged = async (accounts) => {
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
        }
      } else {
        // 断开连接时清除状态
        setAccount(null);
        setProvider(null);
        setSigner(null);
        setConnected(false);
      }
    } catch (error) {
      console.error("Error handling account change:", error);
      // 出错时也清除状态
      setAccount(null);
      setProvider(null);
      setSigner(null);
      setConnected(false);
    }
  };

  // 修改 init 函数
  useEffect(() => {
    const init = async () => {
      if (window.ethereum) {
        try {
          const provider = new ethers.BrowserProvider(window.ethereum);
          const accounts = await provider.listAccounts();

          if (accounts.length > 0) {
            const signer = await provider.getSigner();
            const address = await signer.getAddress();

            // 只在初始化时设置一次
            setAccount(address);
            setProvider(provider);
            setSigner(signer);
            setConnected(true);
          }
        } catch (error) {
          console.error("Error initializing wallet:", error);
        } finally {
          setIsLoading(false);
        }

        // 添加事件监听器
        window.ethereum.on("accountsChanged", handleAccountsChanged);
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