import React, { Component } from 'react';
import Button from 'material-ui/Button';
import Grid from 'material-ui/Grid';
import './App.css';

class App extends Component {

  constructor() {
    super();
    this.state = {
      meal: []
    }
    this.getMeal = this.getMeal.bind(this);
  }

  getMeal() {
    fetch(window.location.origin + '/add', {mode: 'cors'})
      .then(res => res.json())
      .then(data => {
        if (data) {
          this.setState({
            meal: data.ingredients
          })
        }
      })
      .catch(error => console.log(error));
  }

  render() {
    return (
      <div className="App">
        <h2>BFG</h2>
        <p>Click the generate button for a yummy dish!</p>
        <div className="App">
          <Grid item lg={12}>
            <div className="App">
              <h3>Your meal: </h3>
              {this.state.meal.map((value, index) => <strong key={index}><p>{value}</p></strong>)}
            </div>
          </Grid>
          <Grid item lg={12}>
            <div className="App">
              <Button variant="raised" color="primary" onClick={this.getMeal}>Generate</Button>
            </div>
          </Grid>
        </div>
      </div>
    );
  }
}

export default App;
