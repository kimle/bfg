import React from 'react';
import Button from 'material-ui/Button';
import Grid from 'material-ui/Grid';
import Restaurant from 'material-ui-icons/Restaurant';


class Generator extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      hasGenerated: false,
      meal: []
    };
    this.getMeal = this.getMeal.bind(this);
    this.url = this.props.url;
  }

  getMeal() {
    return fetch(this.url + '/add', {
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

  render() {
    const hasGenerated = this.state.hasGenerated;

    return (
      <div>
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
      </div>
    );
  }
}

export default Generator;