import React from 'react';
import logo from './logo.svg';
import './App.css';


class App extends React.Component {

  state = {
    msg: ""
  }

  componentDidMount() {
    fetch("/api/")
    .then(res => res.text())
    .then(data => {
      this.setState({msg: data})
    })
    .catch(console.log)
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Message from api: {this.state.msg}
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    );
  }
}

export default App;
