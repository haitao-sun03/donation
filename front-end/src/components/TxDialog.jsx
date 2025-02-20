import React, { useState } from "react";
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
} from "@mui/material";
import { Typography } from "@mui/material";
import Button from "@mui/material/Button";


export default function TxDialog({ txDialogOpen, txStatus, txHash, onClose }) {
  return (
    <Dialog open={txDialogOpen} onClose={onClose}>
      <DialogTitle>Transaction Status</DialogTitle>
      <DialogContent>
        <DialogContentText>{txStatus}</DialogContentText>
        {txHash && (
          <Typography variant="body2" color="text.secondary" sx={{ mt: 2 }}>
            Transaction Hash: {txHash}
          </Typography>
        )}
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Close</Button>
      </DialogActions>
    </Dialog>
  );
}
