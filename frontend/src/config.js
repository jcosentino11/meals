export const config = {
  auth0: {
    domain: process.env.REACT_APP_AUTH_DOMAIN,
    clientId: process.env.REACT_APP_AUTH_CLIENT_ID,
    audience: process.env.REACT_APP_AUTH_AUDIENCE,
    scope: process.env.REACT_APP_AUTH_SCOPE,
  },
  backend: {
    rootUrl: process.env.REACT_APP_BACKEND_ROOT_URL
  }
};
