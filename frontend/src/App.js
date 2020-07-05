import React from 'react';
import logo from './logo.svg';
import './App.css';
import { config } from './config';
import { LoginButton, LogoutButton, Profile } from './auth';


class App extends React.Component {

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
