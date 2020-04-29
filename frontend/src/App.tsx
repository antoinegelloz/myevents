import React from 'react';
import { HashRouter as Router, Route } from 'react-router-dom'; 
import EventList from './EventList'; 
import BookingForm from './BookingForm';
import event from './event.jpg';

export default function App() {
    const eventList = () => <EventList eventserviceURL="http://localhost:8181" />
    const eventBooking = ({match}: any) => <BookingForm eventID={match.params.id} 
      eventserviceURL="http://localhost:8181"
      bookingserviceURL="http://localhost:8282"/>; 

    const img = {
      width: '100%',
      height: 'auto',
      maxWidth: '100%',
      marginBottom: 8,
    }

    return (
      <Router>
          <React.Fragment>
            <img style={img} src={event} alt="djset"/>
            <Route exact path="/" component={eventList} /> 
            <Route path="/events/bookings/:id" component={eventBooking} /> 
          </React.Fragment>
      </Router>
  );
}