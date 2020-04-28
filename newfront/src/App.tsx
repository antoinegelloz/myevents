import React from 'react';
import { HashRouter as Router, Route } from 'react-router-dom'; 
import Typography from '@material-ui/core/Typography'
import EventList from './EventList'; 
import BookingForm from './BookingForm'; 

export default function App() {
    const eventList = () => <EventList/>; 
    const eventBooking = ({match}: any) => <BookingForm eventID={match.params.id}/>; 

  return (
      <Router> 
          <React.Fragment>
            <Typography variant="h1" align="center" gutterBottom>My Events</Typography>
            <Route exact path="/" component={eventList} /> 
            <Route path="/events/bookings/:id" component={eventBooking} /> 
          </React.Fragment>
      </Router>
  );
}