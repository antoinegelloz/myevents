import React from 'react';
import { HashRouter as Router, Route } from 'react-router-dom'; 
import EventList from './EventList'; 
import BookingForm from './BookingForm';
import event from './event.jpg';

export default function App() {
    var ES_URL = "http://localhost:8181"
    if (process.env.REACT_APP_EVENTSERVICE_URL !== undefined) {
      ES_URL = process.env.REACT_APP_EVENTSERVICE_URL
    }
    var BS_URL = "http://localhost:8282"
    if (process.env.REACT_APP_BOOKINGSERVICE_URL !== undefined) {
      BS_URL = process.env.REACT_APP_BOOKINGSERVICE_URL
    }

    const eventList = () => <EventList eventserviceURL={ES_URL} />
    const eventBooking = ({match}: any) => <BookingForm eventID={match.params.id} 
      eventserviceURL={ES_URL}
      bookingserviceURL={BS_URL}/>;

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