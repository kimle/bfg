import React, { Component } from 'react';
import Button from 'material-ui/Button';
import Grid from 'material-ui/Grid';
import Restaurant from 'material-ui-icons/Restaurant';

import './App.css';
import MeekCard from './MeekCard';

class App extends Component {

  constructor() {
    super();
    this.state = {
      hasGenerated: false,
      meal: [],
      days: null
    }
    this.getMeal = this.getMeal.bind(this);
  }

  componentWillMount() {
    this.getMeekDays();
  }

  getMeal() {
    return fetch(window.location.origin + '/add', {
      method: 'GET',
      mode: 'cors'
    })
    .then(response => response.json())
    .then(data => {
      if (data) {
        this.setState({
          hasGenerated: true,
          meal: data.ingredients
        })
      }
    })
    .catch(error => console.log(error));
  }

  getMeekDays() {
    return fetch(window.location.origin + '/freemeek', {
      method: 'GET',
      mode: 'cors',
    })
    .then(response => response.json())
    .then(data => {
      if (data) {
        this.setState({ days: data.days });
      }
    })
    .catch(error => console.log(error));
  }

  render() {
    const hasGenerated = this.state.hasGenerated;
    
    return (
      <div className="App">
        <h2>BFG</h2>
        <p>Click the generate button for a yummy dish!</p>
        <div className="App">
          <Grid item lg={12}>
            { hasGenerated ?
            <div className="App">
              <h3>Your meal: </h3>
              {this.state.meal.map((value, index) => <strong key={index}><p>{value}</p></strong>)}
            </div> : null }
          </Grid>
          <Grid item lg={12}>
            <div className="App">
              <Button className="button" variant="raised" color="secondary" onClick={this.getMeal}>
                Generate
                <Restaurant className="rightIcon"></Restaurant>
              </Button>
            </div>
          </Grid>
        </div>
        <MeekCard days={this.state.days}/>
      </div>
    );
  }
}

export default App;
