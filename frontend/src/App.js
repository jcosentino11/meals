import React from 'react';
import logo from './logo.svg';
import './App.css';
import * as firebase from "firebase/app";
import "firebase/auth";
import { FirebaseAuthProvider, IfFirebaseAuthed, } from "@react-firebase/auth";
import { config } from "./config";


class App extends React.Component {

  state = {
    msg: ""
  }

  helloWorldEndpoint() {
    return config.backend.rootUrl + "/"
  }

  getUserToken = async () => {
    return await firebase
      .auth()
      .currentUser
      .getIdToken(true)
      .catch(console.log);
  }

  getMessage = async () => {
    const token = await this.getUserToken();
    const headers = {
      headers: new Headers({
        'Authorization': 'Bearer ' + token
      })
    };

    return await fetch(this.helloWorldEndpoint(), headers)
      .then(res => res.text());
  }

  getAndSetMessage = async () => {
    try {
      const data = await this.getMessage();
      this.setState({ msg: data });
    } catch (err) {
      console.log("Failed to get message: " + err);
    }
  }

  clearState() {
    this.setState({ msg: "" })
  }

  render() {
    return (
      <FirebaseAuthProvider firebase={firebase} {...config.firebase}>
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <div>
              <button
                onClick={() => {
                  const googleAuthProvider = new firebase.auth.GoogleAuthProvider();
                  firebase.auth().signInWithPopup(googleAuthProvider).then(result => {
                    this.getAndSetMessage();
                  });
                }}
              >
                Sign In with Google
              </button>
              <button
                onClick={() => {
                  firebase.auth().signOut().finally(() => {
                    this.clearState()
                  });
                }}
              >
                Sign Out
              </button>
              <div>
                <IfFirebaseAuthed>
                  {() => {
                    return (
                      <div>
                        <div>You are authenticated</div>
                        <div>Mesage from api: {this.state.msg}</div>
                      </div>
                    );
                  }}
                </IfFirebaseAuthed>
              </div>
            </div>
          </header>
        </div>
      </FirebaseAuthProvider>
    );
  }
}

export default App;
