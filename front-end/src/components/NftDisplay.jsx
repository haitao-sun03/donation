import React, { useEffect, useState } from "react";
import {
  Box,
  Typography,
  CircularProgress,
  Card,
  CardContent,
  Grid,
  Alert,
  Chip,
} from "@mui/material";
import WarningIcon from "@mui/icons-material/Warning";
import { request } from "../utils/api";

const NftDisplay = ({ currentAddress }) => {
  const [nfts, setNfts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [nftMetadata, setNftMetadata] = useState({});
  const fileBaseGatewayUrl =
    "https://massive-amethyst-tyrannosaurus.myfilebase.com/ipfs/";

  useEffect(() => {
    if (error) {
      const timer = setTimeout(() => {
        setError(null); // 假设你有一个状态管理错误信息
      }, 5000); // 5秒后清除错误
      return () => clearTimeout(timer);
    }
  }, [error]);

  useEffect(() => {
    const fetchNfts = async () => {
      setLoading(true);
      try {
        const requestBody = { user: currentAddress };

        const data = await request("/core/nft/nftGroupByUser", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
          },
          body: JSON.stringify(requestBody),
        });

        if (data.code === 200) {
          if (data.data === null) {
            setNfts([]);
          } else if (Array.isArray(data.data)) {
            setNfts(data.data);
            await fetchNftMetadata(data.data);
          } else {
            setError("Invalid data format");
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
        const metadataUrl = nft.MetaData.replace("ipfs://", fileBaseGatewayUrl);

        const response = await fetch(metadataUrl);
        if (!response.ok) {
          throw new Error("Failed to fetch NFT metadata");
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
    <Box>
      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}
      <Box
        sx={{
          p: 2,
          backgroundColor: "rgba(30, 41, 59, 0.95)",
          borderRadius: "8px",
        }}
      >
        {loading && <CircularProgress />}

        {/* {error && <Typography color="error">{error}</Typography>} */}

        {nfts.length === 0 && !loading && (
          <Box sx={{ textAlign: "center", mt: 4 }}>
            <WarningIcon sx={{ fontSize: 50, color: "#f8fafc" }} />
            <Typography variant="body1" sx={{ color: "#f8fafc", mt: 2 }}>
              No NFTs found for this user.
            </Typography>
          </Box>
        )}
        {nfts.map((nftGroup, index) => (
          <Box key={index} sx={{ mb: 2 }}>
            <Typography variant="h5" sx={{ color: "#f8fafc", mb: 1 }}>
              {nftGroup.NftLevel}
            </Typography>
            <Grid container spacing={2}>
              {nftGroup.Nfts.map((n, idx) => (
                <Grid item xs={12} sm={6} md={4} key={idx}>
                  <Card
                    sx={{
                      background: "rgba(52, 52, 52, 0.9)",
                      borderRadius: "12px",
                      boxShadow: "0 4px 20px rgba(0, 0, 0, 0.2)",
                      transition: "transform 0.2s, box-shadow 0.2s",
                      "&:hover": {
                        transform: "scale(1.05)",
                        boxShadow: "0 8px 40px rgba(0, 0, 0, 0.3)",
                      },
                      p: 1,
                    }}
                  >
                    <CardContent>
                      {nftMetadata[nftGroup.NftLevel] && (
                        <Box sx={{ mt: 1 }}>
                          <Typography
                            variant="h6"
                            sx={{
                              color: "#f8fafc",
                              mb: 1,
                              whiteSpace: "nowrap",
                              overflow: "hidden",
                              textOverflow: "ellipsis",
                            }}
                          >
                            {nftMetadata[nftGroup.NftLevel].name}
                          </Typography>
                          <Typography
                            variant="body2"
                            sx={{
                              color: "#94a3b8",
                              mb: 1,
                            }}
                          >
                            {nftMetadata[nftGroup.NftLevel].description}
                          </Typography>
                          <img
                            src={nftMetadata[nftGroup.NftLevel].image.replace(
                              "ipfs://",
                              fileBaseGatewayUrl
                            )}
                            alt={nftMetadata[nftGroup.NftLevel].name}
                            style={{
                              width: "100%",
                              borderRadius: "8px",
                              marginTop: "8px",
                            }}
                          />
                          <Box sx={{ mt: 1 }}>
                            <Typography
                              variant="body2"
                              sx={{ color: "#f8fafc", mb: 1 }}
                            >
                              Attributes:
                            </Typography>
                            <Box
                              sx={{ display: "flex", flexWrap: "wrap", gap: 1 }}
                            >
                              {nftMetadata[nftGroup.NftLevel].attributes.map(
                                (attr, attrIndex) => (
                                  <Chip
                                    key={attrIndex}
                                    label={`${attr.trait_type}: ${attr.value}`}
                                    sx={{
                                      backgroundColor:
                                        "rgba(156, 163, 175, 0.2)",
                                      color: "#94a3b8",
                                      borderRadius: "8px",
                                    }}
                                  />
                                )
                              )}
                            </Box>
                          </Box>
                        </Box>
                      )}
                      <Typography
                        variant="body2"
                        sx={{ color: "#f8fafc", mb: 1, mt: 1 }}
                      >
                        From campaign: {n.CampaignId}
                      </Typography>
                    </CardContent>
                  </Card>
                </Grid>
              ))}
            </Grid>
          </Box>
        ))}
      </Box>
    </Box>
  );
};

export default NftDisplay;
