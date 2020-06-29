import React from 'react';
import logo from './logo.svg';
import './App.css';
import { config } from './config';
import { LoginButton, LogoutButton, Profile } from './auth';


class App extends React.Component {

  state = {
    msg: ""
  }

  helloWorldEndpoint() {
    return config.backend.rootUrl + "/"
  }

  // getUserToken = async () => {
  //   return await firebase
  //     .auth()
  //     .currentUser
  //     .getIdToken(true)
  //     .catch(console.log);
  // }

  // getMessage = async () => {
  //   const token = await this.getUserToken();
  //   const headers = {
  //     headers: new Headers({
  //       'Authorization': 'Bearer ' + token
  //     })
  //   };

  //   return await fetch(this.helloWorldEndpoint(), headers)
  //     .then(res => res.text());
  // }

  // getAndSetMessage = async () => {
  //   try {
  //     const data = await this.getMessage();
  //     this.setState({ msg: data });
  //   } catch (err) {
  //     console.log("Failed to get message: " + err);
  //   }
  // }

  clearState() {
    this.setState({ msg: "" })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <Profile />
          <div>
            <LoginButton />
            <LogoutButton />
          </div>
        </header>
      </div>
    );
  }
}

export default App;
