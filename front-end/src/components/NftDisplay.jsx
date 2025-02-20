import React, { useEffect, useState } from 'react';
import { Box, Typography, CircularProgress, Card, CardContent, Grid } from '@mui/material';
import WarningIcon from '@mui/icons-material/Warning';

const NftDisplay = ({ currentAddress }) => {
  const [nfts, setNfts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [nftMetadata, setNftMetadata] = useState({});
  const fileBaseGatewayUrl = 'https://massive-amethyst-tyrannosaurus.myfilebase.com/ipfs/';
  
  // 占位图像的 URL
  const placeholderImage = 'https://example.com/path/to/your/placeholder-image.png'; // 替换为实际的图片链接

  useEffect(() => {
    const fetchNfts = async () => {
      setLoading(true);
      try {
        const response = await fetch('/api/nft/nftGroupByUser', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ user: currentAddress }),
        });

        if (!response.ok) {
          throw new Error('Failed to fetch NFTs');
        }

        const data = await response.json();
        if (data.code === 200) {
          if (data.data === null) {
            setNfts([]);
          } else if (Array.isArray(data.data)) {
            setNfts(data.data);
            await fetchNftMetadata(data.data);
          } else {
            setError('Invalid data format');
          }
        } else {
          setError(data.msg);
        }
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    const fetchNftMetadata = async (nftData) => {
      const metadataPromises = nftData.map(async (nft) => {
        const metadataUrl = nft.MetaData.replace('ipfs://', fileBaseGatewayUrl);
        
        const response = await fetch(metadataUrl);
        if (!response.ok) {
          throw new Error('Failed to fetch NFT metadata');
        }
        const metadata = await response.json();
        return { ...nft, metadata };
      });

      const results = await Promise.all(metadataPromises);
      const metadataMap = {};
      results.forEach((item) => {
        metadataMap[item.NftLevel] = item.metadata;
      });
      setNftMetadata(metadataMap);
    };

    if (currentAddress) {
      fetchNfts();
    }
  }, [currentAddress]);

  return (
    <Box sx={{ p: 2, backgroundColor: 'rgba(30, 41, 59, 0.95)', borderRadius: '8px' }}>
      {loading && <CircularProgress />}
      {error && <Typography color="error">{error}</Typography>}
      {nfts.length === 0 && !loading && (
        <Box sx={{ textAlign: 'center', mt: 4 }}>
          <WarningIcon sx={{ fontSize: 50, color: '#f8fafc' }} />
          <Typography variant="body1" sx={{ color: '#f8fafc', mt: 2 }}>No NFTs found for this user.</Typography>
        </Box>
      )}
      {nfts.map((nftGroup, index) => (
        <Box key={index} sx={{ mb: 4 }}>
          <Typography variant="h5" sx={{ color: '#f8fafc', mb: 2 }}>{nftGroup.NftLevel}</Typography>
          <Grid container spacing={2}>
            {nftGroup.Nfts.map((n, idx) => (
              <Grid item xs={12} sm={6} md={4} key={idx}>
                <Card 
                  sx={{ 
                    background: 'rgba(52, 52, 52, 0.8)', 
                    borderRadius: '12px', 
                    boxShadow: 3,
                    transition: 'transform 0.2s',
                    '&:hover': {
                      transform: 'scale(1.05)',
                    },
                    p: 2
                  }}
                >
                  <CardContent>
                    {nftMetadata[nftGroup.NftLevel] && (
                      <Box sx={{ mt: 2 }}>
                        <Typography variant="h6" sx={{ color: '#f8fafc', mb: 1 }}>{nftMetadata[nftGroup.NftLevel].name}</Typography>
                        <Typography variant="body2" sx={{ color: '#94a3b8', mb: 1 }}>{nftMetadata[nftGroup.NftLevel].description}</Typography>
                        <img 
                          src={nftMetadata[nftGroup.NftLevel].image.replace('ipfs://', fileBaseGatewayUrl)} 
                          alt={nftMetadata[nftGroup.NftLevel].name} 
                          style={{ width: '100%', borderRadius: '8px', marginTop: '8px' }} 
                        />
                      </Box>
                    )}
                     <Typography variant="body2" sx={{ color: '#f8fafc', mb: 1,mt:1 }}>From campaign: {n.CampaignId}</Typography>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>
        </Box>
      ))}
    </Box>
  );
};

export default NftDisplay; 