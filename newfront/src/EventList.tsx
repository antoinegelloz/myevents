import React, { useState, useEffect } from 'react';
import { Container, Typography, Paper, TableContainer, Table, TableHead, TableBody, TableRow, TableCell, Button } from '@material-ui/core';
import Event from './models/event';
import { Service } from './Service';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles({
  root: {
    background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
    border: 0,
    borderRadius: 3,
    boxShadow: '0 3px 5px 2px rgba(255, 105, 135, .3)',
    color: 'white',
    height: 40,
    padding: '0 30px',
  },
});

export default function EventList() {
    const classes = useStyles();
    const [result, setResult] = useState<Service<Event[]>>({
        status: 'loading'
      });
    
      useEffect(() => {
        fetch('https://events.gelloz.org/events')
          .then(response => response.json())
          .then(response => setResult({ status: 'ready', payload: response }))
          .catch(error => setResult({ status: 'error', error }));
      }, []);

    return (
    <Container>
        {result.status === 'loading' && (<Typography variant="h3" align="center" gutterBottom>Loading...</Typography>)}
        {result.status === 'error' && (<Typography variant="h3" align="center" gutterBottom>Connection error :(</Typography>)}
        {result.status === 'ready' &&
        (<TableContainer component={Paper}>
            <Table aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell>Event</TableCell>
                        <TableCell align="right">Where</TableCell>
                        <TableCell align="right">When</TableCell>
                        <TableCell align="right"></TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {result.payload.map((event) => (
                        <TableRow>
                            <TableCell component="th" scope="row">{event.name}</TableCell>
                            <TableCell align="right">{event.location.name}</TableCell>
                            <TableCell align="right">{new Date(event.start_date).toDateString()} - {new Date(event.end_date).toDateString()}</TableCell>
                            <TableCell align="right">
                                <Button className={classes.root} href={`#events/bookings/${event.ID}`}>
                                    Book
                                </Button>
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>)}
    </Container>
    );
}