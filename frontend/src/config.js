export const config = {
  auth0: {
    domain: process.env.REACT_APP_AUTH_DOMAIN,
    clientId: process.env.REACT_APP_AUTH_CLIENT_ID,
  },
  backend: {
    rootUrl: process.env.REACT_APP_BACKEND_ROOT_URL
  }
};
