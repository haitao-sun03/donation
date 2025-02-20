import React, { useState } from "react";
import {
  Box,
  TextField,
  Button,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Grid,
} from "@mui/material";
import { ethers } from "ethers";
import { DateTimePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import DonationsManageContract from "../contracts/DonationsManageContract.json"; // 确保路径正确
import TxDialog from "./TxDialog";

const AddCampaign = ({ contractAddress }) => {
  const [title, setTitle] = useState("");
  const [goal, setGoal] = useState("");
  const [startTime, setStartTime] = useState(null);
  const [endTime, setEndTime] = useState(null);
  const [nature, setNature] = useState("0");
  const [beneficiary, setBeneficiary] = useState("0");
  const [purpose, setPurpose] = useState("");
  const [expectedImpact, setExpectedImpact] = useState("");
  const [error, setError] = useState(null);
  const [txDialogOpen, setTxDialogOpen] = useState(false);
  const [txStatus, setTxStatus] = useState("");
  const [txHash, setTxHash] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!window.ethereum) {
      alert("Please install MetaMask!");
      return;
    }

    const provider = new ethers.BrowserProvider(window.ethereum);
    const signer = await provider.getSigner();
    const contract = new ethers.Contract(
      contractAddress,
      DonationsManageContract.abi,
      signer
    );

    try {
      setTxDialogOpen(true);
      setTxStatus("Preparing create campaign transaction...");

      const tx = await contract.createCampaign(
        title,
        ethers.parseUnits(goal, 18), // 假设目标以18位小数表示
        Math.floor(startTime.getTime() / 1000), // 转换为时间戳
        Math.floor(endTime.getTime() / 1000), // 转换为时间戳
        nature,
        beneficiary,
        purpose,
        expectedImpact
      );

      setTxHash(tx.hash);
      setTxStatus(
        "Waiting for create campaign transaction to finish on-chain..."
      );
      await tx.wait();
      setTxStatus("Create campaign transaction completed successfully!");

      // 重置表单字段为默认值
      setTitle("");
      setGoal("");
      setStartTime(null);
      setEndTime(null);
      setNature("0"); // 重置为默认值
      setBeneficiary("0"); // 重置为默认值
      setPurpose("");
      setExpectedImpact("");
    } catch (error) {
      console.error("Error creating campaign:", error);
      setTxStatus("Failed to create campaign.");
    } finally {
      // 交易完成后关闭弹窗
      setTimeout(() => {
        setTxDialogOpen(false);
      }, 2000); // 2秒后关闭弹窗
    }
  };

  return (
    <LocalizationProvider dateAdapter={AdapterDateFns}>
      <Box
        component="form"
        onSubmit={handleSubmit}
        sx={{
          p: 4,
          backgroundColor: "rgba(31,41,55,0.9)",
          borderRadius: "16px",
          boxShadow: 3,
          maxWidth: "600px",
          margin: "0 auto",
        }}
      >
        <Grid container spacing={3}>
          <Grid item xs={12}>
            <TextField
              label="Title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              fullWidth
              required
              variant="outlined"
              placeholder="Enter campaign title"
              sx={{
                backgroundColor: "rgba(55, 65, 81, 0.5)", // Changed to a darker, semi-transparent gray
                borderRadius: "8px",
                "& .MuiInputBase-input": { color: "#FFFFFF" }, // Changed text color to white for better visibility
                "& .MuiInputLabel-root": { color: "#9CA3AF" }, // Changed label color to a lighter gray
                "& .MuiOutlinedInput-root": {
                  "& fieldset": { borderColor: "#4A5568" }, // Adjusted border color to match theme
                  "&:hover fieldset": { borderColor: "#6B7280" }, // Adjusted hover border color
                  "&.Mui-focused fieldset": { borderColor: "#4A90E2" },
                },
              }}
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              label="Goal"
              value={goal}
              onChange={(e) => setGoal(e.target.value)}
              fullWidth
              required
              variant="outlined"
              placeholder="Enter campaign goal (e.g: 1000 AquaCoin)"
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
          </Grid>
          <Grid item xs={12} sm={6}>
            <FormControl fullWidth variant="outlined">
              <InputLabel>Nature</InputLabel>
              <Select
                value={nature}
                onChange={(e) => setNature(e.target.value)}
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
              >
                <MenuItem value="0">Charity</MenuItem>
                <MenuItem value="1">Education</MenuItem>
                <MenuItem value="2">Health</MenuItem>
                <MenuItem value="3">Environment</MenuItem>
                <MenuItem value="4">Other</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12} sm={6}>
            <FormControl fullWidth variant="outlined">
              <InputLabel>Beneficiary</InputLabel>
              <Select
                value={beneficiary}
                onChange={(e) => setBeneficiary(e.target.value)}
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
              >
                <MenuItem value="0">Individuals</MenuItem>
                <MenuItem value="1">NonProfitOrganizations</MenuItem>
                <MenuItem value="2">Communities</MenuItem>
                <MenuItem value="3">Other</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12} sm={6}>
            <DateTimePicker
              label="Start Time"
              value={startTime}
              onChange={(newValue) => setStartTime(newValue)}
              renderInput={(params) => (
                <TextField
                  {...params}
                  fullWidth
                  required
                  variant="outlined"
                  placeholder="Select start time"
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
              )}
            />
          </Grid>
          <Grid item xs={12} sm={6}>
            <DateTimePicker
              label="End Time"
              value={endTime}
              required
              onChange={(newValue) => setEndTime(newValue)}
              renderInput={(params) => (
                <TextField
                  {...params}
                  fullWidth
                  variant="outlined"
                  placeholder="Select end time"
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
              )}
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              label="Purpose"
              value={purpose}
              onChange={(e) => setPurpose(e.target.value)}
              fullWidth
              variant="outlined"
              multiline
              rows={3}
              placeholder="Describe the purpose of the campaign"
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
          </Grid>
          <Grid item xs={12}>
            <TextField
              label="Expected Impact"
              value={expectedImpact}
              onChange={(e) => setExpectedImpact(e.target.value)}
              fullWidth
              variant="outlined"
              multiline
              rows={3}
              placeholder="Describe the expected impact of the campaign"
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
          </Grid>
          <Grid item xs={12}>
            <Button
              type="submit"
              variant="contained"
              fullWidth
              size="large"
              sx={{
                mt: 2,
                backgroundColor: "#4A90E2",
                "&:hover": { backgroundColor: "#357ABD" },
                borderRadius: "8px",
                padding: "12px",
              }}
            >
              Create Campaign
            </Button>
          </Grid>
        </Grid>

        <TxDialog
          txDialogOpen={txDialogOpen}
          txStatus={txStatus}
          txHash={txHash}
          onClose={() => setTxDialogOpen(false)}
        />
      </Box>
    </LocalizationProvider>
  );
};

export default AddCampaign;
