import React from 'react';
import logo from './logo.svg';
import './App.css';
import * as firebase from "firebase/app";
import "firebase/auth";
import {
  FirebaseAuthProvider,
  FirebaseAuthConsumer,
  IfFirebaseAuthed,
  IfFirebaseAuthedAnd
} from "@react-firebase/auth";
import { config } from "./config";


class App extends React.Component {

  state = {
    msg: ""
  }

  componentDidMount() {
    fetch("/api/")
      .then(res => res.text())
      .then(data => {
        this.setState({ msg: data })
      })
      .catch(console.log)
  }

  render() {
    return (
      <FirebaseAuthProvider firebase={firebase} {...config}>
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <div>
              <button
                onClick={() => {
                  const googleAuthProvider = new firebase.auth.GoogleAuthProvider();
                  firebase.auth().signInWithPopup(googleAuthProvider);
                }}
              >
                Sign In with Google
              </button>
              <button
                data-testid="signin-anon"
                onClick={() => {
                  firebase.auth().signInAnonymously();
                }}
              >
                Sign In Anonymously
              </button>
              <button
                onClick={() => {
                  firebase.auth().signOut();
                }}
              >
                Sign Out
              </button>
              <FirebaseAuthConsumer>
                {({ isSignedIn, user, providerId }) => {
                  return (
                    <pre style={{ height: 300, overflow: "auto" }}>
                      {JSON.stringify({ isSignedIn, user, providerId }, null, 2)}
                    </pre>
                  );
                }}
              </FirebaseAuthConsumer>
              <div>
                <IfFirebaseAuthed>
                  {() => {
                    return <div>You are authenticated</div>;
                  }}
                </IfFirebaseAuthed>
                <IfFirebaseAuthedAnd
                  filter={({ providerId }) => providerId !== "anonymous"}
                >
                  {({ providerId }) => {
                    return <div>You are authenticated with {providerId}</div>;
                  }}
                </IfFirebaseAuthedAnd>
              </div>
            </div>
            <p>
              Message from api: {this.state.msg}
            </p>
          </header>
        </div>
      </FirebaseAuthProvider>
    );
  }
}

export default App;
