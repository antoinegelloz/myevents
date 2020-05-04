import React, { useState, useEffect } from 'react';
import { Container, Grid, LinearProgress, Card, CardActions, CardContent } from '@material-ui/core';
import Event from './models/event';
import { Service } from './Service';
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles({
  root: {
    flexGrow: 1,
  },
  card: {
      width: 300,
  },
  cardText: {
      paddingBottom: 0,
  },
  title: {
    fontSize: 14,
  },
  gradientbutton: {
    background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
    border: 0,
    borderRadius: 3,
    boxShadow: '0 3px 5px 2px rgba(255, 105, 135, .3)',
    color: 'white',
    height: 40,
    padding: '0 30px',
  },
});

interface EventListProps {
    eventserviceURL: string,
}

export default function EventList(props: EventListProps) {
    const classes = useStyles();
    const [result, setResult] = useState<Service<Event[]>>({
        status: 'loading'
      });
    
      useEffect(() => {
        fetch(props.eventserviceURL + '/events')
          .then(response => response.json())
          .then(response => setResult({ status: 'ready', payload: response }))
          .catch(error => setResult({ status: 'error', error }));
      }, [props.eventserviceURL]);

    return (
    <Container>
        {result.status === 'loading' && <LinearProgress />}
        {result.status === 'error' && <Typography variant="h3" align="center" gutterBottom>Connection error :(</Typography>}
        {result.status === 'ready' &&
        (<Grid container className={classes.root} spacing={2}>
            <Grid item xs={12}>
                <Grid container justify="center" spacing={2}>
                    {result.payload != null && result.payload.map((event) => (
                    <Grid item>
                        <Card className={classes.card}>
                            <CardContent className={classes.cardText}>
                                <Typography className={classes.title} color="textPrimary" gutterBottom>
                                    {event.name}
                                </Typography>
                                <Typography className={classes.title} color="textSecondary" gutterBottom>
                                    {event.location.name}
                                </Typography>
                                <Typography className={classes.title} color="textSecondary" gutterBottom>
                                    {new Date(event.start_date).toDateString()}
                                </Typography>
                            </CardContent>
                            <CardActions>
                                <Button className={classes.gradientbutton} href={`#events/bookings/${event.ID}`}>
                                    Book
                                </Button>
                            </CardActions>
                        </Card>
                    </Grid>
                ))}
                </Grid>
            </Grid>
        </Grid>)}
    </Container>
    );
}