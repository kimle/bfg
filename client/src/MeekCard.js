import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Card, { CardContent, CardMedia } from 'material-ui/Card';
import Typography from 'material-ui/Typography';

import meek from './images/meek.jpeg';

const styles = {
  card: {
    maxWidth: 450,
    margin: '10px auto',
  },
  media: {
    height: 200,
  },
};

function MeekCard(props) {
  const { classes } = props;
  return (
    <div className="App-meek">
      <Card className={classes.card}>
        <CardMedia
          className={classes.media}
          image={meek}
          title="FreeMeekMill"
        />
        <CardContent>
          <Typography variant="headline" component="h2">
            #JusticeForMeek
          </Typography>
          <Typography component="p">
            <strong>PSA</strong> - Meek Mill has been wrongfully jailed for <strong>{props.days}</strong> days.
          </Typography>
          <Typography component="p">
            <a href="https://www.change.org/p/justice-for-meek-mill-is-justice-for-all-join-the-campaign">
              Please sign the following petition and show your support for him.
            </a>
          </Typography>
        </CardContent>
      </Card>
    </div>
  );
}

MeekCard.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(MeekCard);
