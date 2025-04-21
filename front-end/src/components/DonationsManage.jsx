import { useInView } from "react-intersection-observer";
import React, { useState, useEffect, useRef } from "react";
import { ethers } from "ethers";
import { contractAddresses } from "../main.jsx";
import DonationsManageContract from "../contracts/DonationsManageContract.json";
import Token from "../contracts/Token.json";
import {
  Container,
  Grid,
  Typography,
  TextField,
  Select,
  MenuItem,
  Button,
  Card,
  CardContent,
  CardActions,
  Alert,
  CircularProgress,
  Box,
  Tabs,
  Tab,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Chip,
} from "@mui/material";
import { request } from '../utils/api';
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import FilterListIcon from "@mui/icons-material/FilterList";
import SearchIcon from "@mui/icons-material/Search";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import NftDisplay from "./NftDisplay";
import AddCampaign from "./AddCampaign";
import WarningIcon from "@mui/icons-material/Warning";
import TxDialog from "./TxDialog";

function formatDuration(seconds) {
  const sec = Number(seconds);
  if (sec < 60) {
    return `${sec} seconds`;
  } else if (sec < 3600) {
    return `${Math.floor(sec / 60)} minutes`;
  } else if (sec < 86400) {
    return `${Math.floor(sec / 3600)} hours`;
  }
  return `${Math.floor(sec / 86400)} days`;
}

// 添加自定义主题
const theme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: "#6366f1",
    },
    secondary: {
      main: "#ec4899",
    },
    background: {
      default: "#0f172a",
      paper: "#1e293b",
    },
    text: {
      primary: "#f8fafc",
      secondary: "#94a3b8",
    },
  },
  components: {
    MuiCard: {
      styleOverrides: {
        root: {
          background: "rgba(30, 41, 59, 0.8)",
          backdropFilter: "blur(12px)",
          border: "1px solid rgba(99, 102, 241, 0.1)",
          borderRadius: "16px",
        },
      },
    },
    MuiDialog: {
      styleOverrides: {
        paper: {
          background: "rgba(30, 41, 59, 0.95)",
          backdropFilter: "blur(12px)",
          border: "1px solid rgba(99, 102, 241, 0.1)",
          borderRadius: "16px",
        },
      },
    },
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: "12px",
          textTransform: "none",
          fontWeight: 600,
          padding: "8px 16px",
        },
        contained: {
          background: "linear-gradient(135deg, #6366f1 0%, #4f46e5 100%)",
          "&:hover": {
            background: "linear-gradient(135deg, #4f46e5 0%, #4338ca 100%)",
          },
        },
      },
    },
    MuiTab: {
      styleOverrides: {
        root: {
          textTransform: "none",
          fontWeight: 600,
        },
      },
    },
  },
});

export default function DonationsManage({
  provider,
  signer,
  onBalanceUpdate,
  onSymbolUpdate,
}) {
  const [campaigns, setCampaigns] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [txDialogOpen, setTxDialogOpen] = useState(false);
  const [txStatus, setTxStatus] = useState("");
  const [txHash, setTxHash] = useState("");
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(3);
  const [filterStarter, setFilterStarter] = useState("");
  const [filterStatus, setFilterStatus] = useState(-1);
  const [donationDetailsOpen, setDonationDetailsOpen] = useState(false);
  const [currentDonationDetails, setCurrentDonationDetails] = useState([]);
  const [selectedCampaignId, setSelectedCampaignId] = useState(null);
  const [hasMore, setHasMore] = useState(true);
  const [isLoadingMore, setIsLoadingMore] = useState(false);
  const [donationPage, setDonationPage] = useState(1);
  const [hasMoreDonations, setHasMoreDonations] = useState(true);
  const [loadingMoreDonations, setLoadingMoreDonations] = useState(false);

  const { ref: donationLoadMoreRef, inView } = useInView({
    root: null,
    threshold: 0.5,
    rootMargin: "0px",
    triggerOnce: false,
    skip: !donationDetailsOpen,
  });

  useEffect(() => {
    if (
      donationDetailsOpen &&
      inView &&
      hasMoreDonations &&
      !loadingMoreDonations &&
      selectedCampaignId
    ) {
      const timer = setTimeout(() => {
        console.log("Loading more donations...", {
          inView,
          hasMoreDonations,
          loadingMoreDonations,
          selectedCampaignId,
          currentPage: donationPage,
          isDialogOpen: donationDetailsOpen,
        });
        fetchDonationDetails(selectedCampaignId, true);
      }, 100);

      return () => clearTimeout(timer);
    }
  }, [inView, donationDetailsOpen]);

  useEffect(() => {
    if (error) {
      const timer = setTimeout(() => {
        setError(null); // 假设你有一个状态管理错误信息
      }, 5000); // 5秒后清除错误
      return () => clearTimeout(timer);
    }
  }, [error]);

  const mintTokens = async (e, amount) => {
    try {
      setLoading(true);
      setTxDialogOpen(true);
      setTxStatus("Preparing mint transaction...");

      if (isNaN(amount) || parseFloat(amount) <= 0) {
        throw new Error("Invalid amount");
      }

      // Convert amount to wei
      const amountInWei = ethers.parseUnits(amount, 18);

      // Execute mint transaction
      const tx = await tokenContract.mint(amountInWei);
      setTxHash(tx.hash);
      setTxStatus("Waiting for mint transaction to finish on-chain...");

      await tx.wait();
      setTxStatus("Mint transaction completed successfully");
      setError(null);

      // Refresh token balance
      const newBalance = await tokenContract.balanceOf(currentAddress);
      if (onBalanceUpdate) {
        onBalanceUpdate(ethers.formatUnits(newBalance, 18));
      }

      // Clear amount input
      e.target.reset();
    } catch (err) {
      setError(err.message);
      setTxDialogOpen(false);
    } finally {
      setLoading(false);
    }
  };

  const [contract, setContract] = useState(null);
  const [tokenContract, setTokenContract] = useState(null);
  const [currentAddress, setCurrentAddress] = useState("");
  const [tokenSymbol, setTokenSymbol] = useState("");

  useEffect(() => {
    const init = async () => {
      if (provider && signer) {
        try {
          const address = await signer.getAddress();
          setCurrentAddress(ethers.getAddress(address));

          const donationsContract = new ethers.Contract(
            contractAddresses.donations,
            DonationsManageContract.abi,
            signer
          );
          const token = new ethers.Contract(
            contractAddresses.token,
            Token.abi,
            signer
          );

          setContract(donationsContract);
          setTokenContract(token);

          const balance = await token.balanceOf(address);
          const formattedBalance = ethers.formatUnits(balance, 18);
          const symbol = await token.symbol();
          setTokenSymbol(symbol);

          // 调用回调函数，将余额传递给父组件
          if (onBalanceUpdate) {
            onBalanceUpdate(formattedBalance);
          }
          if (onSymbolUpdate) {
            onSymbolUpdate(symbol);
          }
        } catch (error) {
          console.error("Error initializing contracts:", error);
          setError("Failed to initialize contracts");
        }
      } else {
        console.log("No provider or signer");
        setContract(null);
        setTokenContract(null);
        setCampaigns([]);
        setError(null);
        setCurrentAddress("");
        setBalance("0");
      }
    };

    init();
  }, [provider, signer, onBalanceUpdate]);

  const fetchCampaigns = async (isLoadMore = false) => {
    try {
      console.log("isLoadMore", isLoadMore);

      if (!isLoadMore) {
        setLoading(true);
        // setCampaigns([]);
      } else {
        setIsLoadingMore(true);
      }

      const requestBody = {
        page: isLoadMore ? page + 1 : 1,
        pageSize: pageSize,
        starter: filterStarter || undefined,
        status: filterStatus,
      };

      console.log("Request body:", requestBody);

      const data = await request("/core/campaign/list", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify(requestBody),
      });

      if (data.code === 200) {
        const mappedCampaigns = data.data.map((campaign) => ({
          id: Number(campaign.CampaignID),
          title: campaign.Title,
          goal: campaign.Goal || "0",
          totalDonated: campaign.TotalDonated || "0",
          startTime: new Date(campaign.StartTime).getTime() / 1000,
          endTime: new Date(campaign.EndTime).getTime() / 1000,
          status: Number(campaign.Status),
          starter: campaign.Starter,
          isWithdraw: Number(campaign.IsWithdraw),
          nature: Number(campaign.Nature),
          beneficiary: Number(campaign.Beneficiary),
          purpose: campaign.Purpose,
          expectedImpact: campaign.ExpectedImpact,
        }));

        // 检查是否还有更多数据
        setHasMore(mappedCampaigns.length === pageSize);

        if (isLoadMore) {
          setCampaigns((prev) => [...prev, ...mappedCampaigns]);
          setPage((prev) => prev + 1);
        } else {
          setCampaigns(mappedCampaigns);
          setPage(1);
        }
      } else {
        setError(data.msg || "Failed to fetch campaigns");
      }
    } catch (err) {
      console.error("Fetch error:", err);
      setError(err.message || "Failed to fetch campaigns");
    } finally {
      if (!isLoadMore) {
        setLoading(false);
      } else {
        setIsLoadingMore(false);
      }
    }
  };

  const [donationAmounts, setDonationAmounts] = useState({});

  const donate = async (campaignId) => {
    const amount = (donationAmounts[campaignId] || "").toString();
    if (!amount || isNaN(amount) || parseFloat(amount) <= 0) {
      setError("Please enter a valid donation amount");
      return;
    }
    try {
      setLoading(true);
      // Approve token transfer first
      const approveTx = await tokenContract.approve(
        contractAddresses.donations,
        ethers.parseUnits(amount, 18)
      );
      setTxHash(approveTx.hash);
      setTxStatus("Waiting for approval transaction to finish on-chain...");
      setTxDialogOpen(true);
      await approveTx.wait();
      setTxStatus("Approval transaction finished");
      // Then donate tokens
      const tx = await contract.donate(
        campaignId,
        ethers.parseUnits(amount, 18)
      );
      setTxHash(tx.hash);
      setTxStatus("Waiting for donate transaction to finish on-chain...");
      await tx.wait();
      setTxStatus("Donate transaction finished");
      const newBalance = await tokenContract.balanceOf(currentAddress);
      if (onBalanceUpdate) {
        onBalanceUpdate(ethers.formatUnits(newBalance, 18));
      }
      await fetchCampaigns();
      
      setError(null);
      // Clear donation amount for this campaign
      setDonationAmounts((prev) => ({
        ...prev,
        [campaignId]: "",
      }));


      // 等待一段时间，让后端服务有足够时间更新权限
      await new Promise((resolve) => setTimeout(resolve, 3000)); // 等待3秒

      const requestBody = {
        currentAccount: signer.address,
      };

      console.log("Jwt refresh request body:", requestBody);

      const data = await request("/user/auth/refreshJWT", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify(requestBody),
      });

      if (data.code === 200) {
        let jwts = data.data;
        // 将 address 和 signature 发送给后端
        console.log(`update jwt_${signer.address.toLowerCase()} :`, jwts[0]);
        localStorage.setItem(`jwt_${signer.address.toLowerCase()}`, jwts[0]);

        console.log(`update jwt_refresh_${signer.address.toLowerCase()} :`, jwts[1]);
        localStorage.setItem(`jwt_refresh_${signer.address.toLowerCase()}`, jwts[1]);
      } else {
        setError(data.msg || "Failed to refresh jwt");
      }

    } catch (err) {
      setError(err.message);
      setTxDialogOpen(false);
    } finally {
      setLoading(false);
    }
  };

  const refund = async (campaignId) => {
    try {
      setLoading(true);
      const tx = await contract.refundToken(campaignId);
      setTxHash(tx.hash);
      setTxStatus("Waiting for refund transaction to finish on-chain...");
      setTxDialogOpen(true);
      await tx.wait();
      setTxStatus("Refund transaction finished");

      let newBalance
      // 重新加载数据
      await Promise.all([
        fetchCampaigns(), // 刷新活动列表
        newBalance = await tokenContract.balanceOf(currentAddress),
        fetchDonationDetails(campaignId), // 刷新捐赠详情
      ]);

      if (onBalanceUpdate) {
        onBalanceUpdate(ethers.formatUnits(newBalance, 18));
      }

      setError(null);
    } catch (err) {
      setError(err.message);
      setTxDialogOpen(false);
    } finally {
      setLoading(false);
    }
  };

  const withdraw = async (campaignId) => {
    try {
      setLoading(true);
      const campaign = campaigns.find((c) => c.id === campaignId);
      if (!campaign) {
        throw new Error("Campaign not found");
      }
      if (!campaign.starter || !currentAddress) {
        throw new Error("Invalid address");
      }
      console.log(
        "Attempting to withdraw from campaign:",
        campaignId,
        "Current address:",
        currentAddress,
        "Campaign starter:",
        campaign.starter
      );

      // Validate campaign starter matches current address
      if (
        ethers.getAddress(campaign.starter) !==
        ethers.getAddress(currentAddress)
      ) {
        throw new Error("Only campaign starter can withdraw funds");
      }

      const tx = await contract.withdraw(campaignId);
      setTxHash(tx.hash);
      setTxStatus("Waiting for withdraw transaction to finish on-chain...");
      setTxDialogOpen(true);
      await tx.wait();
      setTxStatus("Withdraw transaction finished");
      const newBalance = await tokenContract.balanceOf(currentAddress);
      if (onBalanceUpdate) {
        onBalanceUpdate(ethers.formatUnits(newBalance, 18));
      }
      await fetchCampaigns();
      setError(null);
    } catch (err) {
      setError(err.message);
      setTxDialogOpen(false);
    } finally {
      setLoading(false);
    }
  };

  const completeCampaign = async (campaignId) => {
    try {
      setLoading(true);
      const tx = await contract.completedCampaign(campaignId);
      setTxHash(tx.hash);
      setTxStatus("Waiting for complete transaction to finish on-chain...");
      setTxDialogOpen(true);
      await tx.wait();
      setTxStatus("Complete transaction finished");
      await new Promise((resolve) => setTimeout(resolve, 2000)); // 等待 2000 毫秒
      await fetchCampaigns();
      setError(null);
    } catch (err) {
      setError(err.message);
      setTxDialogOpen(false);
    } finally {
      setLoading(false);
    }
  };

  const cancelCampaign = async (campaignId) => {
    try {
      setLoading(true);
      const tx = await contract.cancellCampaign(campaignId);
      setTxHash(tx.hash);
      setTxStatus("Waiting for cancel transaction to finish on-chain...");
      setTxDialogOpen(true);
      await tx.wait();
      setTxStatus("Cancel transaction finished");
      await fetchCampaigns();
      setError(null);
    } catch (err) {
      setError(err.message);
      setTxDialogOpen(false);
    } finally {
      setLoading(false);
    }
  };

  const activeCampaign = async (campaignId) => {
    try {
      setLoading(true);
      const tx = await contract.activeCampaign(campaignId);
      setTxHash(tx.hash); 
      setTxStatus("Waiting for avtive transaction to finish on-chain...");
      setTxDialogOpen(true);
      await tx.wait();
      setTxStatus("Avtive transaction finished");
      await fetchCampaigns();
      setError(null);
    } catch (err) {
      setError(err.message);
      setTxDialogOpen(false);
    } finally {
      setLoading(false);
    }
  };

  const fetchDonationDetails = async (campaignId, isLoadMore = false) => {
    try {
      if (!isLoadMore) {
        setLoading(true);
        setDonationPage(1);
      } else {
        setLoadingMoreDonations(true);
      }

      const requestBody = {
        page: isLoadMore ? donationPage + 1 : 1,
        pageSize: 10,
        campaignId: campaignId,
      };

      console.log("Fetching donations:", {
        isLoadMore,
        currentPage: donationPage,
        requestBody,
      });

      const data = await request('/core/donation/list', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
        body: JSON.stringify(requestBody),
      });
  
      console.log("Response data:", data);

      if (data.code === 200) {
        const donations = data.data || [];
        console.log("Received donations:", {
          count: donations.length,
          totalCount: data.count,
          donations,
        });

        const mappedDonations = donations.map((donation) => ({
          donor: donation.Donor,
          amount: donation.Amount || "0",
          isRefund: Boolean(donation.IsRefund),
          time: donation.CreatedAt,
          mintLevel: donation.MintLevel,
        }));

        console.log("Mapped donations:", {
          count: mappedDonations.length,
          currentPage: donationPage,
          mappedDonations,
        });

        const currentTotal = isLoadMore
          ? currentDonationDetails.length + mappedDonations.length
          : mappedDonations.length;
        setHasMoreDonations(currentTotal < data.count);

        if (isLoadMore) {
          setCurrentDonationDetails((prev) => [...prev, ...mappedDonations]);
          setDonationPage((prev) => prev + 1);
        } else {
          setCurrentDonationDetails(mappedDonations);
        }

        setSelectedCampaignId(campaignId);
        setDonationDetailsOpen(true);
      } else {
        setError(data.msg || "Failed to fetch donation details");
      }
    } catch (err) {
      console.error("Fetch error:", err);
      setError(err.message || "Failed to fetch donation details");
    } finally {
      if (!isLoadMore) {
        setLoading(false);
      } else {
        setLoadingMoreDonations(false);
      }
    }
  };

  const [tabValue, setTabValue] = useState(0);

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };

  // 添加一个单独的 effect 来处理 tabValue 变化
  useEffect(() => {
    console.log("tabValue", tabValue);
    if (tokenContract && tabValue === 0) {
      const getTokenInfo = async () => {
        // Get balance and owner
        const balance = await tokenContract.balanceOf(currentAddress);
        const symbol = await tokenContract.symbol();
        setTokenSymbol(symbol);
      };
      getTokenInfo();
    }

    if (contract && tabValue === 2) {
      const getTokenInfo = async () => {
        setPage(1); // 重置页面
        setPageSize(3); // 确保 pageSize 被重置
        await fetchCampaigns();
      };
      getTokenInfo();
    }
  }, [tabValue]);

  const isAddressMatch = (addr1, addr2) => {
    if (!addr1 || !addr2) return false;
    try {
      return ethers.getAddress(addr1) === ethers.getAddress(addr2);
    } catch (err) {
      console.error("Error comparing addresses:", err);
      return false;
    }
  };

  const handleLoadMore = () => {
    if (!isLoadingMore && hasMore) {
      fetchCampaigns(true);
    }
  };

  const handleLoadMoreDonations = () => {
    if (!loadingMoreDonations && hasMoreDonations && selectedCampaignId) {
      fetchDonationDetails(selectedCampaignId, true);
    }
  };

  // 占位图像的 URL
  const placeholderImage =
    "https://example.com/path/to/your/placeholder-image.png"; // 替换为实际的图片链接

  return (
    <ThemeProvider theme={theme}>
      <Box
        sx={{
          minHeight: "100vh",
          background: "linear-gradient(135deg, #0f172a 0%, #1e1b4b 100%)",
          pt: 3,
        }}
      >
        <Container
          maxWidth="lg"
          sx={{
            py: 6,
            background:
              "linear-gradient(152deg, rgba(17,24,39,0.9) 0%, rgba(31,41,55,0.9) 100%)",
            borderRadius: "24px",
            backdropFilter: "blur(20px)",
            boxShadow: "0px 4px 30px rgba(99, 102, 241, 0.2)",
            border: "1px solid rgba(99, 102, 241, 0.2)",
            my: 4,
            position: "relative",
            overflow: "hidden",
            "&:before": {
              content: '""',
              position: "absolute",
              width: "200%",
              height: "200%",
              top: "-50%",
              left: "-50%",
              background:
                "radial-gradient(circle, rgba(99,102,241,0.1) 0%, transparent 60%)",
              animation: "rotate 20s linear infinite",
              "@keyframes rotate": {
                "0%": { transform: "rotate(0deg)" },
                "100%": { transform: "rotate(360deg)" },
              },
            },
          }}
        >
          <Typography
            variant="h3"
            component="h1"
            gutterBottom
            sx={{
              fontWeight: 700,
              color: "text.primary",
              mb: 4,
              textAlign: "center",
              background: "linear-gradient(90deg, #6366F1 0%, #8B5CF6 100%)",
              WebkitBackgroundClip: "text",
              WebkitTextFillColor: "transparent",
            }}
          >
            Manage Donations
          </Typography>

          <Tabs
            value={tabValue}
            onChange={handleTabChange}
            sx={{
              mb: 4,
              "& .MuiTabs-indicator": {
                height: 3,
                background: "linear-gradient(90deg, #6366F1 0%, #8B5CF6 100%)",
                borderRadius: "2px",
              },
              "& .MuiTab-root": {
                fontSize: "1rem",
                fontWeight: 600,
                textTransform: "none",
                minWidth: 120,
                mx: 1,
                color: "text.secondary",
                transition: "all 0.2s ease",
                "&:hover": {
                  color: "#6366F1",
                  transform: "translateY(-2px)",
                },
                "&.Mui-selected": {
                  color: "#6366F1",
                  fontWeight: 700,
                },
              },
            }}
          >
            <Tab label="Mint Tokens" />
            <Tab label="Create Campaign" />
            <Tab label="Manage Campaigns" />
            <Tab label="User NFTs" />
          </Tabs>

          {error && (
            <Alert severity="error" sx={{ mb: 2 }}>
              {error}
            </Alert>
          )}

          {loading && (
            <Box sx={{ display: "flex", justifyContent: "center", my: 2 }}>
              <CircularProgress />
            </Box>
          )}

          <Box sx={{ mb: 4 }}>
            {tabValue === 0 && (
              <Box>
                <Typography variant="h6" color="text.secondary" gutterBottom>
                  {`To make the functionality more convenient to experience, mint
                  ${tokenSymbol} for the current user.`}
                </Typography>

                <Box
                  component="form"
                  onSubmit={async (e) => {
                    e.preventDefault();
                    const formData = new FormData(e.target);
                    const amount = formData.get("amount");
                    await mintTokens(e, amount);
                  }}
                  sx={{
                    p: 4,
                    backgroundColor: "rgba(31,41,55,0.9)",
                    borderRadius: "16px",
                    boxShadow: 3,
                    maxWidth: "600px",
                    margin: "0 auto",
                  }}
                >
                  <TextField
                    fullWidth
                    name="amount"
                    type="number"
                    label="Amount"
                    placeholder="Enter amount to mint"
                    inputProps={{ min: 1 }}
                    required
                    variant="outlined"
                    sx={{
                      backgroundColor: "rgba(55, 65, 81, 0.5)",
                      borderRadius: "8px",
                      "& .MuiInputBase-input": { color: "#FFFFFF" },
                      "& .MuiInputLabel-root": { color: "#9CA3AF" },
                      "& .MuiOutlinedInput-root": {
                        "& fieldset": { borderColor: "#4A5568" },
                        "&:hover fieldset": { borderColor: "#6B7280" },
                        "&.Mui-focused fieldset": { borderColor: "#4A90E2" },
                      },
                    }}
                  />
                  <Button
                    type="submit"
                    variant="contained"
                    fullWidth
                    disabled={loading}
                  >
                    {loading ? <CircularProgress size={24} /> : "Mint Tokens"}
                  </Button>
                </Box>
              </Box>
            )}

            {tabValue === 1 && (
              <AddCampaign contractAddress={contractAddresses.donations} />
            )}

            {tabValue === 2 && (
              <>
                <Box sx={{ mb: 3 }}>
                  <Grid container spacing={2} alignItems="center">
                    <Grid item xs={12} sm={4}>
                      <TextField
                        fullWidth
                        label="Filter by Starter"
                        value={filterStarter}
                        onChange={(e) => setFilterStarter(e.target.value)}
                        InputProps={{
                          startAdornment: <SearchIcon />,
                        }}
                      />
                    </Grid>
                    <Grid item xs={12} sm={4}>
                      <Select
                        fullWidth
                        value={filterStatus}
                        onChange={(e) =>
                          setFilterStatus(Number(e.target.value))
                        }
                        displayEmpty
                        startAdornment={<FilterListIcon />}
                      >
                        <MenuItem value={-1}>All Status</MenuItem>
                        <MenuItem value={0}>Pending</MenuItem>
                        <MenuItem value={1}>Active</MenuItem>
                        <MenuItem value={2}>Completed</MenuItem>
                        <MenuItem value={3}>Cancelled</MenuItem>
                      </Select>
                    </Grid>
                    <Grid item xs={12} sm={4}>
                      <Button
                        variant="contained"
                        onClick={() => fetchCampaigns()}
                        startIcon={<SearchIcon />}
                      >
                        Search
                      </Button>
                    </Grid>
                  </Grid>
                </Box>

                <Box sx={{ textAlign: "center", mt: 4 }}>
                  {campaigns.length === 0 && !loading && (
                    <Box sx={{ textAlign: "center", mt: 4 }}>
                      <WarningIcon sx={{ fontSize: 50, color: "#f8fafc" }} />
                      <Typography
                        variant="body1"
                        sx={{ color: "#f8fafc", mt: 2 }}
                      >
                        No campaigns found.
                      </Typography>
                    </Box>
                  )}
                </Box>

                <Grid container spacing={3}>
                  {campaigns.map((campaign, index) => (
                    <Grid item xs={12} sm={6} md={4} key={index}>
                      <Card
                        sx={{
                          background: "rgba(31,41,55,0.6)",
                          border: "1px solid rgba(99,102,241,0.2)",
                          boxShadow: "0px 4px 30px rgba(99, 102, 241, 0.1)",
                          transition: "all 0.3s ease",
                          "&:hover": {
                            transform: "translateY(-4px)",
                            boxShadow: "0px 8px 40px rgba(99, 102, 241, 0.3)",
                          },
                        }}
                      >
                        <CardContent>
                          <Typography
                            variant="h5"
                            component="h3"
                            gutterBottom
                            sx={{
                              fontWeight: 600,
                              color: "text.primary",
                              background:
                                "linear-gradient(90deg, #6366F1 0%, #8B5CF6 100%)",
                              WebkitBackgroundClip: "text",
                              WebkitTextFillColor: "transparent",
                              mb: 3,
                            }}
                          >
                            {campaign.title}
                            <Typography
                              variant="body2"
                              component="span"
                              sx={{
                                color: "text.secondary",
                                ml: 1,
                                verticalAlign: "middle",
                              }}
                            >
                              (ID: {campaign.id})
                            </Typography>
                          </Typography>
                          <Typography variant="body1" gutterBottom>
                            Goal:{" "}
                            {campaign.goal
                              ? ethers.formatUnits(campaign.goal, 18)
                              : "0"}{" "}
                          </Typography>
                          <Typography variant="body1" gutterBottom>
                            Raised:{" "}
                            {campaign.totalDonated
                              ? ethers.formatUnits(campaign.totalDonated, 18)
                              : "0"}{" "}
                          </Typography>

                          <Typography variant="body1" gutterBottom>
                            Time:{" "}
                            {new Date(
                              Number(campaign.startTime) * 1000
                            ).toLocaleString()}{" "}
                            {" ~ "}
                            {new Date(
                              Number(campaign.endTime) * 1000
                            ).toLocaleString()}
                          </Typography>

                          <Typography variant="body1" gutterBottom>
                            Starter:{" "}
                            {campaign.starter
                              ? `${campaign.starter.slice(
                                  0,
                                  6
                                )}...${campaign.starter.slice(-4)}`
                              : "Unknown"}
                          </Typography>

                          <Typography variant="body1" gutterBottom>
                            Nature:{" "}
                            {
                              [
                                "Charity",
                                "Education",
                                "Health",
                                "Environment",
                                "Other",
                              ][campaign.nature]
                            }
                          </Typography>

                          <Typography variant="body1" gutterBottom>
                            Beneficiary:{" "}
                            {
                              [
                                "Individuals",
                                "NonProfitOrganizations",
                                "Communities",
                                "Other",
                              ][campaign.beneficiary]
                            }
                          </Typography>
                          <Typography variant="body1" gutterBottom>
                            Purpose: {campaign.purpose}
                          </Typography>

                          <Typography variant="body1" gutterBottom>
                            Expected Impact: {campaign.expectedImpact}
                          </Typography>

                          <Box sx={{ mb: 2, mt: 2 }}>
                            <Chip
                              label={
                                ["Pending", "Active", "Completed", "Cancelled"][
                                  campaign.status
                                ]
                              }
                              sx={{
                                backgroundColor: [
                                  "warning.light",
                                  "success.light",
                                  "info.light",
                                  "error.light",
                                ][campaign.status],
                                color: "white",
                                fontWeight: 600,
                              }}
                            />
                          </Box>

                          <Box sx={{ mt: 2, display: "flex", gap: 1 }}>
                            <Chip
                              label={
                                campaign.isWithdraw === 1
                                  ? "Withdraw"
                                  : "Not Withdraw"
                              }
                              sx={{
                                backgroundColor: [
                                  "primary.main",
                                  "secondary.main",
                                ][campaign.isWithdraw],
                                color: "white",
                                fontWeight: 600,
                              }}
                            />
                          </Box>
                        </CardContent>
                        <CardActions>
                          <TextField
                            fullWidth
                            type="number"
                            label="Donation Amount"
                            value={donationAmounts[campaign.id] ?? ""}
                            onChange={(e) => {
                              const value = e.target.value;
                              setDonationAmounts((prev) => ({
                                ...prev,
                                [campaign.id]: value,
                              }));
                            }}
                            inputProps={{ min: 0.01, step: 0.01 }}
                            sx={{
                              mb: 2,
                              "& .MuiInputBase-root": {
                                height: "56px",
                                "& input[type=number]": {
                                  padding: "14.5px 14px",
                                  MozAppearance: "textfield",
                                  "&::-webkit-outer-spin-button, &::-webkit-inner-spin-button":
                                    {
                                      WebkitAppearance: "none",
                                      margin: 0,
                                    },
                                },
                              },
                              "& .MuiInputLabel-root": {
                                transform: "translate(14px, 9px) scale(1)",
                                "&.Mui-focused": {
                                  transform:
                                    "translate(14px, -9px) scale(0.75)",
                                },
                                "&.MuiFormLabel-filled": {
                                  transform:
                                    "translate(14px, -9px) scale(0.75)",
                                },
                              },
                              "& .MuiOutlinedInput-input": {
                                padding: "14.5px 14px",
                              },
                            }}
                            disabled={loading}
                          />
                          <Button
                            variant="contained"
                            fullWidth
                            onClick={() => donate(campaign.id)}
                            disabled={
                              loading ||
                              !(Number(donationAmounts[campaign.id]) > 0) ||
                              Number(campaign.status) !== 1 ||
                              !campaign.starter ||
                              !currentAddress ||
                              isAddressMatch(campaign.starter, currentAddress)
                            }
                            sx={{
                              "&.Mui-disabled": {
                                backgroundColor: "rgba(99, 102, 241, 0.12)",
                                color: "rgba(99, 102, 241, 0.5)",
                                border: "1px solid rgba(99, 102, 241, 0.2)",
                              },
                            }}
                          >
                            {loading ? (
                              <CircularProgress size={24} />
                            ) : (
                              "Donate"
                            )}
                          </Button>
                        </CardActions>
                        <CardActions>
                          <Grid container spacing={1}>
                            <Grid item xs={6}>
                              <Button
                                variant="outlined"
                                fullWidth
                                onClick={() => withdraw(campaign.id)}
                                disabled={
                                  loading ||
                                  !campaign.starter ||
                                  !currentAddress ||
                                  !isAddressMatch(
                                    campaign.starter,
                                    currentAddress
                                  ) ||
                                  Number(campaign.status) !== 2 ||
                                  Number(campaign.isWithdraw) !== 0 ||
                                  Number(campaign.totalDonated) === 0
                                }
                              >
                                {loading ? (
                                  <CircularProgress size={24} />
                                ) : (
                                  "Withdraw"
                                )}
                              </Button>
                            </Grid>
                            <Grid item xs={6}>
                              <Button
                                variant="outlined"
                                fullWidth
                                onClick={() => completeCampaign(campaign.id)}
                                disabled={
                                  loading ||
                                  !campaign.starter ||
                                  !currentAddress ||
                                  !isAddressMatch(
                                    campaign.starter,
                                    currentAddress
                                  ) ||
                                  Number(campaign.status) !== 1
                                }
                              >
                                {loading ? (
                                  <CircularProgress size={24} />
                                ) : (
                                  "Complete"
                                )}
                              </Button>
                            </Grid>
                            <Grid item xs={6}>
                              <Button
                                variant="outlined"
                                color="error"
                                fullWidth
                                onClick={() => cancelCampaign(campaign.id)}
                                disabled={
                                  loading ||
                                  !campaign.starter ||
                                  !currentAddress ||
                                  !isAddressMatch(
                                    campaign.starter,
                                    currentAddress
                                  ) ||
                                  Number(campaign.status) === 2 ||
                                  Number(campaign.status) === 3
                                }
                              >
                                {loading ? (
                                  <CircularProgress size={24} />
                                ) : (
                                  "Cancel"
                                )}
                              </Button>
                            </Grid>
                            <Grid item xs={6}>
                              <Button
                                variant="outlined"
                                color="success"
                                fullWidth
                                onClick={() => activeCampaign(campaign.id)}
                                disabled={
                                  loading ||
                                  !campaign.starter ||
                                  !currentAddress ||
                                  !isAddressMatch(
                                    campaign.starter,
                                    currentAddress
                                  ) ||
                                  Number(campaign.status) !== 0
                                }
                              >
                                {loading ? (
                                  <CircularProgress size={24} />
                                ) : (
                                  "Active"
                                )}
                              </Button>
                            </Grid>
                          </Grid>
                        </CardActions>
                        <CardActions>
                          <Button
                            variant="outlined"
                            fullWidth
                            onClick={() => fetchDonationDetails(campaign.id)}
                          >
                            View Donations
                          </Button>
                        </CardActions>
                      </Card>
                    </Grid>
                  ))}
                </Grid>

                {/* 添加加载更多按钮 */}
                {hasMore && (
                  <Box
                    sx={{ display: "flex", justifyContent: "center", mt: 4 }}
                  >
                    <Button
                      variant="outlined"
                      onClick={handleLoadMore}
                      disabled={isLoadingMore}
                      sx={{
                        minWidth: 200,
                        py: 1,
                        background: "rgba(31,41,55,0.6)",
                        borderColor: "rgba(99,102,241,0.3)",
                        "&:hover": {
                          borderColor: "#6366F1",
                          background: "rgba(31,41,55,0.8)",
                        },
                      }}
                    >
                      {isLoadingMore ? (
                        <CircularProgress size={24} />
                      ) : (
                        "Load More Campaigns"
                      )}
                    </Button>
                  </Box>
                )}

                {/* 捐款详情对话框 */}
                <Dialog
                  open={donationDetailsOpen}
                  onClose={() => {
                    setDonationDetailsOpen(false);
                    setTimeout(() => {
                      setCurrentDonationDetails([]);
                      setDonationPage(1);
                      setHasMoreDonations(true);
                      setSelectedCampaignId(null);
                    }, 300);
                  }}
                  maxWidth="md"
                  fullWidth
                  PaperProps={{
                    sx: {
                      background: "rgba(30, 41, 59, 0.95)",
                      backdropFilter: "blur(12px)",
                      border: "1px solid rgba(99, 102, 241, 0.1)",
                      borderRadius: "16px",
                      boxShadow: "0 8px 32px rgba(0, 0, 0, 0.4)",
                      color: "#f8fafc",
                    },
                  }}
                >
                  <DialogTitle
                    sx={{
                      borderBottom: "1px solid rgba(99, 102, 241, 0.1)",
                      background:
                        "linear-gradient(90deg, rgba(31,41,55,0.95) 0%, rgba(17,24,39,0.95) 100%)",
                      color: "#f8fafc",
                      fontSize: "1.25rem",
                      fontWeight: 600,
                    }}
                  >
                    Donation Details for Campaign #{selectedCampaignId}
                  </DialogTitle>
                  <DialogContent sx={{ p: 0 }}>
                    <Box sx={{ height: "60vh" }}>
                      <TableContainer
                        sx={{
                          height: "100%",
                          overflow: "auto",
                          "&::-webkit-scrollbar": {
                            width: "8px",
                          },
                          "&::-webkit-scrollbar-track": {
                            background: "rgba(31,41,55,0.3)",
                          },
                          "&::-webkit-scrollbar-thumb": {
                            background: "rgba(99,102,241,0.3)",
                            borderRadius: "4px",
                            "&:hover": {
                              background: "rgba(99,102,241,0.5)",
                            },
                          },
                        }}
                      >
                        <Table stickyHeader>
                          <TableHead>
                            <TableRow>
                              <TableCell
                                sx={{
                                  background: "rgba(31,41,55,0.95)",
                                  color: "#94a3b8",
                                  fontWeight: 600,
                                  borderBottom:
                                    "1px solid rgba(99,102,241,0.1)",
                                }}
                              >
                                Donor
                              </TableCell>
                              <TableCell
                                align="center"
                                sx={{
                                  background: "rgba(31,41,55,0.95)",
                                  color: "#94a3b8",
                                  fontWeight: 600,
                                  borderBottom:
                                    "1px solid rgba(99,102,241,0.1)",
                                }}
                              >
                                Amount
                              </TableCell>
                              <TableCell
                                align="center"
                                sx={{
                                  background: "rgba(31,41,55,0.95)",
                                  color: "#94a3b8",
                                  fontWeight: 600,
                                  borderBottom:
                                    "1px solid rgba(99,102,241,0.1)",
                                }}
                              >
                                Status
                              </TableCell>
                              <TableCell
                                align="center"
                                sx={{
                                  background: "rgba(31,41,55,0.95)",
                                  color: "#94a3b8",
                                  fontWeight: 600,
                                  borderBottom:
                                    "1px solid rgba(99,102,241,0.1)",
                                }}
                              >
                                Time
                              </TableCell>
                              <TableCell
                                align="center"
                                sx={{
                                  background: "rgba(31,41,55,0.95)",
                                  color: "#94a3b8",
                                  fontWeight: 600,
                                  borderBottom:
                                    "1px solid rgba(99,102,241,0.1)",
                                }}
                              >
                                Refund
                              </TableCell>
                              <TableCell
                                align="center"
                                sx={{
                                  background: "rgba(31,41,55,0.95)",
                                  color: "#94a3b8",
                                  fontWeight: 600,
                                  borderBottom:
                                    "1px solid rgba(99,102,241,0.1)",
                                }}
                              >
                                NFT Level
                              </TableCell>
                            </TableRow>
                          </TableHead>
                          <TableBody>
                            {currentDonationDetails.map((donation, index) => (
                              <TableRow
                                key={index}
                                sx={{
                                  background: "transparent",
                                  "&:hover": {
                                    background: "rgba(99,102,241,0.1)",
                                  },
                                  transition: "background-color 0.2s ease",
                                }}
                              >
                                <TableCell
                                  sx={{
                                    color: "#f8fafc",
                                    borderBottom:
                                      "1px solid rgba(99,102,241,0.1)",
                                  }}
                                >
                                  {`${donation.donor.slice(
                                    0,
                                    6
                                  )}...${donation.donor.slice(-4)}`}
                                </TableCell>
                                <TableCell
                                  align="center"
                                  sx={{
                                    color: "#f8fafc",
                                    borderBottom:
                                      "1px solid rgba(99,102,241,0.1)",
                                  }}
                                >
                                  {donation.amount
                                    ? ethers.formatUnits(donation.amount, 18)
                                    : 0}
                                </TableCell>
                                <TableCell
                                  align="center"
                                  sx={{
                                    borderBottom:
                                      "1px solid rgba(99,102,241,0.1)",
                                  }}
                                >
                                  <Chip
                                    label={
                                      donation.isRefund
                                        ? "Refunded"
                                        : "Not Refunded"
                                    }
                                    sx={{
                                      backgroundColor: donation.isRefund
                                        ? "rgba(52,211,153,0.2)"
                                        : "rgba(251,191,36,0.2)",
                                      color: donation.isRefund
                                        ? "#34D399"
                                        : "#FBB224",
                                      borderRadius: "8px",
                                      fontWeight: 600,
                                      fontSize: "0.75rem",
                                    }}
                                  />
                                </TableCell>
                                <TableCell
                                  align="center"
                                  sx={{
                                    color: "#f8fafc",
                                    borderBottom:
                                      "1px solid rgba(99,102,241,0.1)",
                                  }}
                                >
                                  {new Date(donation.time).toLocaleString()}
                                </TableCell>
                                <TableCell
                                  align="center"
                                  sx={{
                                    borderBottom:
                                      "1px solid rgba(99,102,241,0.1)",
                                  }}
                                >
                                  <Button
                                    variant="contained"
                                    color="secondary"
                                    size="small"
                                    onClick={() => refund(selectedCampaignId)}
                                    disabled={
                                      loading ||
                                      donation.isRefund ||
                                      !campaigns.find(
                                        (c) => c.id === selectedCampaignId
                                      )?.status ||
                                      campaigns.find(
                                        (c) => c.id === selectedCampaignId
                                      )?.status !== 3 ||
                                      !currentAddress ||
                                      !donation.donor ||
                                      !isAddressMatch(
                                        donation.donor,
                                        currentAddress
                                      )
                                    }
                                    sx={{
                                      minWidth: "80px",
                                      fontSize: "0.75rem",
                                      py: 0.5,
                                    }}
                                  >
                                    {loading ? (
                                      <CircularProgress size={20} />
                                    ) : (
                                      "Refund"
                                    )}
                                  </Button>
                                </TableCell>
                                <TableCell
                                  align="center"
                                  sx={{
                                    color: "#f8fafc",
                                    borderBottom:
                                      "1px solid rgba(99,102,241,0.1)",
                                  }}
                                >
                                  <Chip
                                    label={donation.mintLevel || "Not Minted"}
                                    sx={{
                                      backgroundColor: donation.mintLevel
                                        ? "rgba(52,211,153,0.2)"
                                        : "rgba(156,163,175,0.2)",
                                      color: donation.mintLevel
                                        ? "#34D399"
                                        : "#9CA3AF",
                                      borderRadius: "8px",
                                      fontWeight: 600,
                                      fontSize: "0.75rem",
                                    }}
                                  />
                                </TableCell>
                              </TableRow>
                            ))}
                          </TableBody>
                        </Table>
                        {hasMoreDonations && (
                          <Box
                            ref={donationLoadMoreRef}
                            sx={{
                              py: 2,
                              textAlign: "center",
                              background: "transparent",
                            }}
                          >
                            {loadingMoreDonations ? (
                              <CircularProgress
                                size={24}
                                sx={{ color: "rgba(99,102,241,0.8)" }}
                              />
                            ) : (
                              <Typography
                                variant="body2"
                                sx={{
                                  color: "#94a3b8",
                                  fontSize: "0.875rem",
                                }}
                              >
                                Scroll to load more
                              </Typography>
                            )}
                          </Box>
                        )}
                      </TableContainer>
                    </Box>
                  </DialogContent>
                  <DialogActions
                    sx={{
                      borderTop: "1px solid rgba(99,102,241,0.1)",
                      background:
                        "linear-gradient(90deg, rgba(31,41,55,0.95) 0%, rgba(17,24,39,0.95) 100%)",
                      px: 3,
                      py: 2,
                    }}
                  >
                    <Button
                      onClick={() => {
                        setDonationDetailsOpen(false);
                        setTimeout(() => {
                          setCurrentDonationDetails([]);
                          setDonationPage(1);
                          setHasMoreDonations(true);
                          setSelectedCampaignId(null);
                        }, 300);
                      }}
                      sx={{
                        background: "rgba(99,102,241,0.1)",
                        color: "#f8fafc",
                        borderRadius: "8px",
                        px: 3,
                        "&:hover": {
                          background: "rgba(99,102,241,0.2)",
                        },
                      }}
                    >
                      Close
                    </Button>
                  </DialogActions>
                </Dialog>
              </>
            )}

            {tabValue === 3 && <NftDisplay currentAddress={currentAddress} />}
          </Box>

          <TxDialog
            txDialogOpen={txDialogOpen}
            txStatus={txStatus}
            txHash={txHash}
            onClose={() => setTxDialogOpen(false)}
          />
        </Container>
      </Box>
    </ThemeProvider>
  );
}
