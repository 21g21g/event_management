// Function to decode base64Url
function base64UrlDecode(str) {
    // Convert base64Url to base64
    str = str.replace(/-/g, '+').replace(/_/g, '/');
    // Decode base64
    const decoded = atob(str);
    // Convert decoded string to JSON
    console.log(JSON.parse(decoded))

    return JSON.parse(decoded);
  }
  
  // Function to decode JWT
  export function decodeToken(token) {
    try {
      // Split token into parts
      const [header, payload, signature] = token.split('.');
      // Decode payload
      
      return base64UrlDecode(payload);
    } catch (error) {
      console.error('Invalid token:', error);
      return null;
    }
  }
  
  // Function to get user ID from decoded token
  export function getUserIdFromToken(token) {
    console.log(token)
    const decodedToken = decodeToken(token);
    if (decodedToken && decodedToken['https://hasura.io/jwt/claims']) {
      return decodedToken['https://hasura.io/jwt/claims']['x-hasura-user-id'];
    }
    return null;
  }