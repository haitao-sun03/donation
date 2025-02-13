import React, { useState, useEffect, useRef } from 'react'
import { ethers } from 'ethers'
import { contractAddresses } from '../main.jsx'
import DonationsManageContract from '../contracts/DonationsManageContract.json'
import Token from '../contracts/Token.json'
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
  Chip
} from '@mui/material'
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';

function formatDuration(seconds) {
  const sec = Number(seconds)
  if (sec < 60) {
    return `${sec} seconds`
  } else if (sec < 3600) {
    return `${Math.floor(sec / 60)} minutes`
  } else if (sec < 86400) {
    return `${Math.floor(sec / 3600)} hours`
  }
  return `${Math.floor(sec / 86400)} days`
}

export default function DonationsManage({ provider, signer }) {
  const [campaigns, setCampaigns] = useState([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [txDialogOpen, setTxDialogOpen] = useState(false)
  const [txStatus, setTxStatus] = useState('')
  const [txHash, setTxHash] = useState('')

  const mintTokens = async (e, amount) => {
    try {
      setLoading(true)
      setTxDialogOpen(true)
      setTxStatus('Preparing mint transaction...')

      if (isNaN(amount) || parseFloat(amount) <= 0) {
        throw new Error('Invalid amount')
      }
      
      // Convert amount to wei
      const amountInWei = ethers.parseUnits(amount, 18)
      
      // Execute mint transaction
      const tx = await tokenContract.mint(amountInWei)
      setTxHash(tx.hash)
      setTxStatus('Waiting for mint transaction to finish on-chain...')
      
      await tx.wait()
      setTxStatus('Mint transaction completed successfully')
      setError(null)
      
      // Refresh token balance
      const newBalance = await tokenContract.balanceOf(currentAddress)
      setBalance(ethers.formatUnits(newBalance, 18))
      
      // Clear amount input
      e.target.reset()
    } catch (err) {
      setError(err.message)
      setTxDialogOpen(false)
    } finally {
      setLoading(false)
    }
  }

  const [contract, setContract] = useState(null)
  const [tokenContract, setTokenContract] = useState(null)
  const [currentAddress, setCurrentAddress] = useState('')
  const [balance, setBalance] = useState('0')

  useEffect(() => {
    
    const init = async () => {
      if (provider && signer) {
        console.log("Donation Manage中init,进入了provider && signer")
        try {
          const address = await signer.getAddress()
          setCurrentAddress(ethers.getAddress(address))
          
          const donationsContract = new ethers.Contract(
            contractAddresses.donations,
            DonationsManageContract.abi,
            signer
          )
          const token = new ethers.Contract(
            contractAddresses.token,
            Token.abi,
            signer
          )
          
          setContract(donationsContract)
          setTokenContract(token)
          
          // Get balance and owner
          const balance = await token.balanceOf(address)
          setBalance(ethers.formatUnits(balance, 18))
          
          await fetchCampaigns(donationsContract)
        } catch (error) {
          console.error('Error initializing contracts:', error)
          setError('Failed to initialize contracts')
        }
      } else {
        console.log('No provider or signer')
        setContract(null)
        setTokenContract(null)
        setCampaigns([])
        setError(null)
      }
    }
    
    init()
    
    // Add listener for account changes
    const handleAccountsChanged = async (accounts) => {
      if (provider && signer && accounts.length > 0) {
        try {
          const newSigner = await provider.getSigner()
          const address = await newSigner.getAddress()
          setCurrentAddress(ethers.getAddress(address))
          
          const donationsContract = new ethers.Contract(
            contractAddresses.donations,
            DonationsManageContract.abi,
            newSigner
          )
          const token = new ethers.Contract(
            contractAddresses.token,
            Token.abi,
            newSigner
          )
          setContract(donationsContract)
          setTokenContract(token)
          await fetchCampaigns(donationsContract)
        } catch (error) {
          console.error('Error handling account change:', error)
          setError('Failed to handle account change')
        }
      } else {
        console.log('No provider or signer after account change')
        setContract(null)
        setTokenContract(null)
        setCampaigns([])
        setError(null)
      }
    }
    
    window.ethereum?.on('accountsChanged', handleAccountsChanged)
    
    return () => {
      window.ethereum?.removeListener('accountsChanged', handleAccountsChanged)
    }
  }, [provider, signer])

  const fetchCampaigns = async (contractInstance = contract) => {
    try {
      setLoading(true)
      const campaignCount = await contractInstance.campaignCount()
      const count = Number(campaignCount)
      const campaigns = []
      
      for (let i = 1; i <= count; i++) {
        const campaign = await contractInstance.getCampaignDetails(i)
        campaigns.push({
          id: i, // Store ID as number
          title: campaign.title,
          goal: campaign.goal,
          totalDonated: campaign.totalDonated,
          startTime: campaign.startTime,
          endTime: campaign.endTime,
          status: campaign.status,
          starter: campaign.starter
        })
      }
      
      setCampaigns(campaigns)
      setError(null)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const formRef = useRef(null)

  const createCampaign = async (title, goal, duration, timeUnit) => {
    try {
      setLoading(true)
      setTxDialogOpen(true)
      setTxStatus('Preparing transaction...')
      
      // Add minimum 3 minute delay before campaign starts
      const startTime = Math.floor(Date.now() / 1000)
      const endTime = startTime + (duration * timeUnit)
      
      // Validate time parameters
      if (endTime <= startTime) {
        throw new Error('End time must be after start time')
      }
      if (duration <= 0) {
        throw new Error('Duration must be positive')
      }
      
      const tx = await contract.createCampaign(title, ethers.parseUnits(goal, 18), startTime, endTime)
      setTxHash(tx.hash)
      setTxStatus('Waiting for transaction to finish on-chain...')
      
      await tx.wait()
      setTxStatus('Transaction completed successfully')
      await fetchCampaigns()
      setError(null)
    } catch (err) {
      setError(err.message)
      setTxDialogOpen(false)
    } finally {
      setLoading(false)
    }
  }

  const [donationAmounts, setDonationAmounts] = useState({})

  const donate = async (campaignId) => {
    const amount = (donationAmounts[campaignId] || '').toString()
    if (!amount || isNaN(amount) || parseFloat(amount) <= 0) {
      setError('Please enter a valid donation amount')
      return
    }
    try {
      setLoading(true)
      // Approve token transfer first
      const approveTx = await tokenContract.approve(
        contractAddresses.donations,
        ethers.parseUnits(amount, 18)
      )
      setTxHash(approveTx.hash)
      setTxStatus('Waiting for approval transaction to finish on-chain...')
      setTxDialogOpen(true)
      await approveTx.wait()
      
      // Then donate tokens
      const tx = await contract.donate(campaignId, ethers.parseUnits(amount, 18))
      setTxHash(tx.hash)
      setTxStatus('Waiting for donation transaction to finish on-chain...')
      await tx.wait()
      setTxStatus('Transaction finished')
      await fetchCampaigns()
      setError(null)
      // Clear donation amount for this campaign
      setDonationAmounts(prev => ({
        ...prev,
        [campaignId]: ''
      }))
    } catch (err) {
      setError(err.message)
      setTxDialogOpen(false)
    } finally {
      setLoading(false)
    }
  }

  const refund = async (campaignId) => {
    try {
      setLoading(true)
      const tx = await contract.refundToken(campaignId)
      setTxHash(tx.hash)
      setTxStatus('Waiting for transaction to finish on-chain...')
      setTxDialogOpen(true)
      await tx.wait()
      setTxStatus('Transaction finished')
      await fetchCampaigns()
      setError(null)
    } catch (err) {
      setError(err.message)
      setTxDialogOpen(false)
    } finally {
      setLoading(false)
    }
  }

  const withdraw = async (campaignId) => {
    try {
      setLoading(true)
      const campaign = campaigns.find(c => c.id === campaignId)
      if (!campaign) {
        throw new Error('Campaign not found')
      }
      console.log('Attempting to withdraw from campaign:', campaignId, 'Current address:', currentAddress, 'Campaign starter:', campaign.starter);
      
      // Validate campaign starter matches current address
      if (ethers.getAddress(campaign.starter) !== ethers.getAddress(currentAddress)) {
        throw new Error('Only campaign starter can withdraw funds')
      }
      
      const tx = await contract.withdraw(campaignId)
      setTxHash(tx.hash)
      setTxStatus('Waiting for transaction to finish on-chain...')
      setTxDialogOpen(true)
      await tx.wait()
      setTxStatus('Transaction finished')
      await fetchCampaigns()
      setError(null)
    } catch (err) {
      setError(err.message)
      setTxDialogOpen(false)
    } finally {
      setLoading(false)
    }
  }

  const completeCampaign = async (campaignId) => {
    try {
      setLoading(true)
      const tx = await contract.completedCampaign(campaignId)
      setTxHash(tx.hash)
      setTxStatus('Waiting for transaction to finish on-chain...')
      setTxDialogOpen(true)
      await tx.wait()
      setTxStatus('Transaction finished')
      await fetchCampaigns()
      setError(null)
    } catch (err) {
      setError(err.message)
      setTxDialogOpen(false)
    } finally {
      setLoading(false)
    }
  }

  const cancelCampaign = async (campaignId) => {
    try {
      setLoading(true)
      const tx = await contract.cancellCampaign(campaignId)
      setTxHash(tx.hash)
      setTxStatus('Waiting for transaction to finish on-chain...')
      setTxDialogOpen(true)
      await tx.wait()
      setTxStatus('Transaction finished')
      await fetchCampaigns()
      setError(null)
    } catch (err) {
      setError(err.message)
      setTxDialogOpen(false)
    } finally {
      setLoading(false)
    }
  }


  const [tabValue, setTabValue] = useState(0)

  const handleTabChange = (event, newValue) => {
    setTabValue(newValue)
  }

  // Refresh campaigns when tab changes or contract updates
  useEffect(() => {
    if (contract && tabValue == 2) {
      fetchCampaigns()
    }
  }, [tabValue, contract])

  return (
      <Container maxWidth="lg" sx={{
        py: 6,
        background: 'linear-gradient(152deg, rgba(17,24,39,0.9) 0%, rgba(31,41,55,0.9) 100%)',
        borderRadius: '24px',
        backdropFilter: 'blur(20px)',
        boxShadow: '0px 4px 30px rgba(99, 102, 241, 0.2)',
        border: '1px solid rgba(99, 102, 241, 0.2)',
        my: 4,
        position: 'relative',
        overflow: 'hidden',
        '&:before': {
          content: '""',
          position: 'absolute',
          width: '200%',
          height: '200%',
          top: '-50%',
          left: '-50%',
          background: 'radial-gradient(circle, rgba(99,102,241,0.1) 0%, transparent 60%)',
          animation: 'rotate 20s linear infinite',
          '@keyframes rotate': {
            '0%': { transform: 'rotate(0deg)' },
            '100%': { transform: 'rotate(360deg)' }
          }
        }
      }}>
      <Typography variant="h3" component="h1" gutterBottom sx={{ 
        fontWeight: 700,
        color: 'text.primary',
        mb: 4,
        textAlign: 'center',
        background: 'linear-gradient(90deg, #6366F1 0%, #8B5CF6 100%)',
        WebkitBackgroundClip: 'text',
        WebkitTextFillColor: 'transparent'
      }}>
        Manage Donations
      </Typography>

      <Tabs value={tabValue} onChange={handleTabChange} sx={{ 
        mb: 4,
        '& .MuiTabs-indicator': {
          height: 3,
          background: 'linear-gradient(90deg, #6366F1 0%, #8B5CF6 100%)',
          borderRadius: '2px'
        },
        '& .MuiTab-root': {
          fontSize: '1rem',
          fontWeight: 600,
          textTransform: 'none',
          minWidth: 120,
          mx: 1,
          color: 'text.secondary',
          transition: 'all 0.2s ease',
          '&:hover': {
            color: '#6366F1',
            transform: 'translateY(-2px)'
          },
          '&.Mui-selected': {
            color: '#6366F1',
            fontWeight: 700
          }
        }
      }}>
        <Tab label="Mint Tokens" />
        <Tab label="Add Campaign" />
        <Tab label="Show Campaigns" />
      </Tabs>

      <Box sx={{ mb: 4 }}>
        {tabValue === 0 && (
          <Box>
            <Typography variant="h6" color="text.secondary" gutterBottom>
            Mint tokens for the current account (for convenience in experience)
            </Typography>
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <AccountBalanceIcon fontSize="small" />
              <Typography variant="body1" gutterBottom>
                {balance} Tokens
              </Typography>
            </Box>
          </Box>
        )}
        {tabValue === 1 && (
          <Typography variant="h6" color="text.secondary" gutterBottom>
            Create a new donation campaign
          </Typography>
        )}
        {tabValue === 2 && (
          <Typography variant="h6" color="text.secondary" gutterBottom>
            View and manage existing campaigns
          </Typography>
        )}
      </Box>

      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}

      {loading && (
        <Box sx={{ display: 'flex', justifyContent: 'center', my: 2 }}>
          <CircularProgress />
        </Box>
      )}

      {tabValue === 0 && (
        <Card sx={{ 
          mb: 4, 
          boxShadow: 3,
          backgroundColor: 'rgba(31,41,55,0.9)',
          border: '1px solid rgba(99,102,241,0.3)'
        }}>
          <CardContent sx={{ color: '#ffffff' }}>
            <Grid container spacing={3} component="form" onSubmit={async (e) => {
              e.preventDefault()
              const formData = new FormData(e.target)
              const amount = formData.get('amount')
              await mintTokens(e, amount)
            }}>
              <Grid item xs={12}>
                <TextField
                  fullWidth
                  name="amount"
                  type="number"
                  label="Amount"
                  placeholder="Enter amount to mint"
                  inputProps={{ min: 1 }}
                  required
                  variant="outlined"
                  InputLabelProps={{ shrink: true }}
                  sx={{ mb: 2 }}
                />
                <Button 
                  type="submit" 
                  variant="contained" 
                  fullWidth 
                  disabled={loading}
                >
                  {loading ? <CircularProgress size={24} /> : 'Mint Tokens'}
                </Button>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
      )}

      {tabValue === 1 && (
      <Card sx={{ 
        mb: 4, 
        boxShadow: 3,
        backgroundColor: 'rgba(31,41,55,0.9)',
        border: '1px solid rgba(99,102,241,0.3)'
      }}>
        <CardContent sx={{ color: '#ffffff' }}>
          <Grid 
            container
            spacing={3}
            component="form"
            ref={formRef}
            sx={{
              '& .MuiTextField-root': {
                marginBottom: '16px'
              },
              '& .MuiGrid-item': {
                paddingTop: '16px'
              }
            }}
            onSubmit={async (e) => {
              e.preventDefault()
              const formData = new FormData(e.target)
              const title = formData.get('title')
              const goal = formData.get('goal')
              const duration = formData.get('duration')
              const timeUnit = parseInt(formData.get('timeUnit'))
              await createCampaign(title, goal, duration, timeUnit)
              formRef.current.reset()
            }}>
            <Grid item xs={12} sx={{ mb: 3 }}>
              
              <Grid container spacing={3}>
                <Grid item xs={12}>
                  <Box sx={{ mb: 3 }}>
                    <Typography variant="subtitle1" gutterBottom sx={{ fontWeight: 500, color: 'text.primary', mb: 1 }}>
                      Campaign Title
                    </Typography>
                    <TextField
                      fullWidth
                      name="title"
                      label="Title"
                      placeholder="Enter campaign title"
                      required
                      variant="outlined"
                      InputLabelProps={{ shrink: true }}
                      sx={{ mb: 2 }}
                    />

                  </Box>
                </Grid>

                <Grid item xs={12}>
                  <Box sx={{ mb: 3 }}>
                    <Typography variant="subtitle1" gutterBottom sx={{ fontWeight: 500, color: 'text.primary', mb: 1 }}>
                      Funding Goal
                    </Typography>
                    <TextField
                      fullWidth
                      name="goal"
                      type="number"
                      label="Amount (Tokens)"
                      placeholder="Enter funding goal"
                      inputProps={{ min: 1 }}
                      required
                      variant="outlined"
                      InputLabelProps={{ shrink: true }}
                      sx={{ mb: 2 }}
                    />
                  </Box>
                </Grid>

                <Grid item xs={12}>
                  <Box sx={{ mb: 3 }}>
                    <Typography variant="subtitle1" gutterBottom sx={{ fontWeight: 500, color: 'text.primary', mb: 1 }}>
                      Campaign Duration
                    </Typography>
                    <Box sx={{ display: 'flex', gap: 2, alignItems: 'center' }}>
                      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 1 }}>
                        <Typography variant="subtitle2" color="text.secondary">
                          Duration
                        </Typography>
                        <Box sx={{ display: 'flex', gap: 2, alignItems: 'center' }}>
                          <TextField
                            name="duration"
                            type="number"
                            placeholder="Enter duration"
                            inputProps={{ min: 1 }}
                            required
                            variant="outlined"
                            InputLabelProps={{ shrink: true }}
                            sx={{ 
                              width: '180px',
                              mb: 2
                            }}
                          />
                          <Select
                            name="timeUnit"
                            defaultValue={86400}
                            required
                            variant="outlined"
                            sx={{
                              width: '180px',
                              backgroundColor: 'background.paper',
                              '& .MuiSelect-select': {
                                padding: '14.5px 14px',
                                height: '56px'
                              },
                              '& .MuiOutlinedInput-notchedOutline': {
                                borderColor: 'rgba(0, 0, 0, 0.23)'
                              }
                            }}
                          >
                            <MenuItem value={1}>Seconds</MenuItem>
                            <MenuItem value={60}>Minutes</MenuItem>
                            <MenuItem value={3600}>Hours</MenuItem>
                            <MenuItem value={86400}>Days</MenuItem>
                          </Select>
                        </Box>
                      </Box>
                    </Box>
                  </Box>
                </Grid>
              </Grid>
            </Grid>
            <Grid item xs={12}>
              <Button type="submit" variant="contained" fullWidth>
                Create Campaign
              </Button>
            </Grid>
          </Grid>
        </CardContent>
      </Card>
      )}

      {tabValue === 2 && (

      <Grid container spacing={3}>
        {campaigns.map((campaign, index) => (
          <Grid item xs={12} sm={6} md={4} key={index}>
            <Card sx={{
              background: 'rgba(31,41,55,0.6)',
              border: '1px solid rgba(99,102,241,0.2)',
              boxShadow: '0px 4px 30px rgba(99, 102, 241, 0.1)',
              transition: 'all 0.3s ease',
              '&:hover': {
                transform: 'translateY(-4px)',
                boxShadow: '0px 8px 40px rgba(99, 102, 241, 0.3)'
              }
            }}>
              <CardContent>
                <Typography variant="h5" component="h3" gutterBottom sx={{
                  fontWeight: 600,
                  color: 'text.primary',
                  background: 'linear-gradient(90deg, #6366F1 0%, #8B5CF6 100%)',
                  WebkitBackgroundClip: 'text',
                  WebkitTextFillColor: 'transparent',
                  mb: 3
                }}>
                  {campaign.title} 
                  <Typography variant="body2" component="span" sx={{ 
                    color: 'text.secondary',
                    ml: 1,
                    verticalAlign: 'middle'
                  }}>
                    (ID: {campaign.id})
                  </Typography>
                </Typography>
                <Typography variant="body1" gutterBottom>
                  Goal: {ethers.formatUnits(campaign.goal, 18)} Tokens
                </Typography>
                <Typography variant="body1" gutterBottom>
                  Raised: {ethers.formatUnits(campaign.totalDonated, 18)} Tokens
                </Typography>
                <Typography variant="body1" gutterBottom>
                  Duration: {formatDuration(campaign.endTime - campaign.startTime)}
                </Typography>
                <Typography variant="body1" gutterBottom>
                  Start Time: {new Date(Number(campaign.startTime) * 1000).toLocaleString()}
                </Typography>
                <Box sx={{ mb: 2 }}>
                  <Chip 
                    label={['Pending', 'Active', 'Completed', 'Cancelled'][campaign.status]}
                    sx={{
                      backgroundColor: ['warning.light', 'success.light', 'info.light', 'error.light'][campaign.status],
                      color: 'white',
                      fontWeight: 600
                    }}
                  />
                </Box>
                <Typography variant="body2" color="text.secondary">
                  Starter: {campaign.starter.slice(0, 6)}...{campaign.starter.slice(-4)}
                </Typography>
              </CardContent>
              <CardActions>
                <TextField
                  fullWidth
                  type="number"
                  label="Donation Amount"
                  value={donationAmounts[campaign.id] ?? ''}
                  onChange={(e) => {
                    const value = e.target.value
                    setDonationAmounts(prev => ({
                      ...prev,
                      [campaign.id]: value
                    }))
                  }}
                  inputProps={{ min: 0.01, step: 0.01 }}
                  sx={{ 
                    mb: 2,
                    '& .MuiInputBase-root': {
                      height: '56px',
                      '& input[type=number]': {
                        padding: '14.5px 14px',
                        MozAppearance: 'textfield',
                        '&::-webkit-outer-spin-button, &::-webkit-inner-spin-button': {
                          WebkitAppearance: 'none',
                          margin: 0
                        }
                      }
                    },
                    '& .MuiInputLabel-root': {
                      transform: 'translate(14px, 9px) scale(1)',
                      '&.Mui-focused': {
                        transform: 'translate(14px, -9px) scale(0.75)'
                      },
                      '&.MuiFormLabel-filled': {
                        transform: 'translate(14px, -9px) scale(0.75)'
                      }
                    },
                    '& .MuiOutlinedInput-input': {
                      padding: '14.5px 14px'
                    }
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
                    Number(campaign.status) !== 1 || // Enable only for active campaigns
                    ethers.getAddress(campaign.starter) === ethers.getAddress(currentAddress) // Disable for campaign starter
                  }
                  sx={{
                    '&.Mui-disabled': {
                      backgroundColor: 'rgba(99, 102, 241, 0.12)',
                      color: 'rgba(99, 102, 241, 0.5)',
                      border: '1px solid rgba(99, 102, 241, 0.2)'
                    }
                  }}
                >
                  {loading ? <CircularProgress size={24} /> : 'Donate Tokens'}
                </Button>
              </CardActions>
              <CardActions>
                <Grid container spacing={1}>
                  <Grid item xs={6}>
                    <Button
                      variant="outlined"
                      fullWidth
                      onClick={() => refund(campaign.id)}
                      disabled={
                        loading ||
                        ethers.getAddress(campaign.starter) === ethers.getAddress(currentAddress) ||
                        Number(campaign.status) !== 3 || // Only allow refunds for Completed campaigns
                        Number(campaign.totalDonated) === 0
                      }
                    >
                      {loading ? <CircularProgress size={24} /> : 'Refund'}
                    </Button>
                  </Grid>
                  <Grid item xs={6}>
                    <Button
                      variant="outlined"
                      fullWidth
                      onClick={() => withdraw(index + 1)}
                      disabled={
                        loading ||
                        ethers.getAddress(campaign.starter) !== ethers.getAddress(currentAddress) ||
                        Number(campaign.status) !== 2 || // Only allow withdrawals for completed campaigns
                        Number(campaign.totalDonated) === 0
                      }
                    >
                      {loading ? <CircularProgress size={24} /> : 'Withdraw'}
                    </Button>
                  </Grid>
                  <Grid item xs={6}>
                    <Button
                      variant="outlined"
                      fullWidth
                      onClick={() => completeCampaign(campaign.id)}
                      disabled={
                        loading ||
                        ethers.getAddress(campaign.starter) !== ethers.getAddress(currentAddress) ||
                        Number(campaign.status) !== 1 // Only allow closing active campaigns
                      }
                    >
                      {loading ? <CircularProgress size={24} /> : 'Complete'}
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
                        ethers.getAddress(campaign.starter) !== ethers.getAddress(currentAddress) ||
                        Number(campaign.status) === 2 || Number(campaign.status) === 3 //Completed or Cancelled can not Cancel again
                      }
                    >
                      {loading ? <CircularProgress size={24} /> : 'Cancel'}
                    </Button>
                  </Grid>
                </Grid>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
      )}
      {/* Transaction Dialog */}
      <Dialog open={txDialogOpen} onClose={() => setTxDialogOpen(false)}>
        <DialogTitle>Transaction Status</DialogTitle>
        <DialogContent>
          <DialogContentText>
            {txStatus}
          </DialogContentText>
          {txHash && (
            <Typography variant="body2" color="text.secondary" sx={{ mt: 2 }}>
              Transaction Hash: {txHash}
            </Typography>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setTxDialogOpen(false)}>Close</Button>
        </DialogActions>
      </Dialog>
    </Container>
  )
}
