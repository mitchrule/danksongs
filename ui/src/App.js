/** 
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
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

export default App;
*/

import React, { Component } from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import SimpleReactLightbox from "simple-react-lightbox";
import LoginPage from "./components/LoginPage";

class App extends Component {
  render() {
    return (
        <Router>
          <div className="app">
            <SimpleReactLightbox>
              {" "}
              {/* Don't touch this wrapper, needed for gallery */}
              <Switch>
                <Route exact path="/" component={LoginPage} />
              </Switch>
            </SimpleReactLightbox>
          </div>
        </Router>
    );
  }
}

export default App;
