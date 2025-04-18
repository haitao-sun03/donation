export const parseJwt = (token) => {
  try {
    console.log("parseJwt,jwt is :", token)
    const payload = JSON.parse(atob(token.split(".")[1]));
    console.log("parseJwt : ",payload);
    return payload;
  } catch {
    return null;
  }
};

export const isTokenExpiringSoon = (token, threshold = 300) => {
  // 默认5分钟
  const payload = parseJwt(token);
  if (!payload) return true;
  return payload.exp - Date.now() / 1000 < threshold;
};
