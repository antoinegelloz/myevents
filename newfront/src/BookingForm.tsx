import React, { useState, useEffect } from 'react';
import Event from "./models/event";
import { Grid, Typography, Button } from '@material-ui/core';
import Select from '@material-ui/core/Select';
import Divider from '@material-ui/core/Divider';
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

interface BookingFormProps {
    eventID: string,
}

export default function BookingForm(props: BookingFormProps) {
    const classes = useStyles();
    const [result, setResult] = useState<Service<Event>>({
        status: 'loading'
      });
    
    const [quantity, setQuantity] = useState('1');

    const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setQuantity(event.target.value as string);
    };

    const handleSubmit = () => { 
        const payload = {booking_quantity: parseInt(quantity)}; 
    
        setResult({ status: 'saving' })
    
        fetch('https://bookings.gelloz.org/events/bookings/' + props.eventID, {method: "POST", body: JSON.stringify(payload)})
        .then(function(response) {
            if(response.ok) {
                setResult({ status: 'done' })
            }
        }).catch(function(error) {
            setResult({ status: 'error', error })
        })
    }

    useEffect(() => {
        fetch('https://events.gelloz.org/events/id/' + props.eventID)
            .then(response => response.json())
            .then(response => setResult({ status: 'ready', payload: response }))
            .catch(error => setResult({ status: 'error', error }));
    }, [props.eventID]);

    return (
    <React.Fragment>
        {result.status === 'loading' && (<Typography variant="h3" align="center" gutterBottom>Loading...</Typography>)}
        {result.status === 'saving' && (<Typography variant="h3" align="center" gutterBottom>Saving...</Typography>)}
        {result.status === 'done' && (<Typography variant="h3" align="center" gutterBottom>Booking is done! Congratulations :)</Typography>)}
        {result.status === 'error' && (<Typography variant="h3" align="center" gutterBottom>Connection error :(</Typography>)}
        {result.status === 'ready' && (
        <Grid
        container
        spacing={0}
        direction="column"
        alignItems="center"
        justify="center"
        >
            <Typography variant="h3" align="center" gutterBottom>Book tickets for {result.payload.name}</Typography>
            <Divider />
            <br />
            <Select value={quantity} native onChange={handleChange} label="Quantity">
                <option value="1">1</option>
                <option value="2">2</option>
                <option value="3">3</option>
                <option value="4">4</option>
                <option value="5">5</option>
            </Select>
            <br />
            <br />
            <Button className={classes.root} onClick={handleSubmit}>
                Book
            </Button>
        </Grid>
        )}
    </React.Fragment>
    );
}