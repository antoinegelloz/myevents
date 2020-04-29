import React, { useState, useEffect } from 'react';
import Event from "./models/event";
import { Grid, Typography, Button, LinearProgress } from '@material-ui/core';
import Divider from '@material-ui/core/Divider';
import { Service } from './Service';
import { createStyles, makeStyles } from '@material-ui/core/styles';
import { Formik, Form, Field } from 'formik';
import { TextField } from 'formik-material-ui';

interface Values {
  email: string;
}

const useStyles = makeStyles(() =>
  createStyles({
    root: {
        background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
        border: 0,
        borderRadius: 3,
        boxShadow: '0 3px 5px 2px rgba(255, 105, 135, .3)',
        color: 'white',
        height: 40,
        padding: '0 30px',
        margin: '20px',
    },
  }),
);

interface BookingFormProps {
    eventID: string,
    eventserviceURL: string,
    bookingserviceURL: string,
}

export default function BookingForm(props: BookingFormProps) {
    const classes = useStyles();
    const [result, setResult] = useState<Service<Event>>({
        status: 'loading'
      });
    
    const handleSubmit = (values: Values) => { 
        const payload = {booking_quantity: 1, user_email: values.email}; 
        setResult({ status: 'saving' })
        fetch(props.bookingserviceURL + '/events/bookings/' + props.eventID, {method: "POST", body: JSON.stringify(payload)})
        .then(function(response) {
            if(response.ok) {
                setResult({ status: 'done' })
            }
        }).catch(function(error) {
            setResult({ status: 'error', error })
        })
    }

    useEffect(() => {
        fetch(props.eventserviceURL + '/events/id/' + props.eventID)
            .then(response => response.json())
            .then(response => setResult({ status: 'ready', payload: response }))
            .catch(error => setResult({ status: 'error', error }));
    });

    return (
    <React.Fragment>
        {result.status === 'loading' 
            && <LinearProgress />}
        {result.status === 'saving' 
            && <LinearProgress />}
        {result.status === 'done' 
            && <Typography variant="h5" align="center" gutterBottom>Booking is done!<br /> Congratulations :)</Typography>}
        {result.status === 'error' 
            && <Typography variant="h5" align="center" gutterBottom>Connection error :(</Typography>}
        {result.status === 'ready' && (
        <Grid
        container
        spacing={0}
        direction="column"
        alignItems="center"
        justify="center"
        >
            <Typography variant="h5" align="center" gutterBottom>Book your ticket for<br />{result.payload.name}</Typography>
            <Divider />
            <br />
            <Formik
            initialValues={{ email: '' }}
            validate={values => {
                const errors: Partial<Values> = {};
                if (!values.email) {
                errors.email = 'Required';
                } else if (
                !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(values.email)
                ) {
                errors.email = 'Invalid email address';
                }
                return errors;
            }}
            onSubmit={(values, { setSubmitting }) => {
                setTimeout(() => {
                setSubmitting(false);
                handleSubmit(values);
                }, 500);
            }}
            >
            {({ submitForm, isSubmitting }) => (
                <Form style={{textAlign: 'center'}}>
                    <Field component={TextField} variant="outlined" name="email" type="email" label="Your email" />
                    {isSubmitting && <LinearProgress />}
                    <br />
                    <Button className={classes.root} disabled={isSubmitting} onClick={submitForm} >
                        Book
                    </Button>
                </Form>
            )}
            </Formik>
        </Grid>
        )}
    </React.Fragment>
    );
}