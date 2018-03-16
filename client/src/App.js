import React, { Component } from 'react';

import Generator from './components/Generator';
import MeekCard from './components/MeekCard';
import './App.css';

class App extends Component {
  render() {
    const url = window.location.port === "3000" ? "http://localhost:8080" : window.location.origin;
    return (
      <div className="App">
        <Generator url={url}/>
        <MeekCard url={url}/>
      </div>
    );
  }
}

export default App;
