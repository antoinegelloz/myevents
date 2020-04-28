import * as React from "react"; 
import * as ReactDOM from "react-dom"; 
import {EventListContainer} from "./components/event_list_container"; 
import {EventBookingFormContainer} from "./components/event_booking_form_container"; 
import {HashRouter as Router, Route} from "react-router-dom"; 
import {Navigation} from "./components/navigation";

class App extends React.Component<{}, {}> { 
  render() { 
    const eventList = () => <EventListContainer eventListURL="https://events.gelloz.org/events" />; 
    const eventBooking = ({match}: any) => <EventBookingFormContainer eventID={match.params.id} 
        eventServiceURL="https://events.gelloz.org" 
        bookingServiceURL="https://bookings.gelloz.org" />; 
   
    return <Router> 
      <Navigation brandName="My Events"/>
      <div className="container"> 
        <Route exact path="/" component={eventList} /> 
        <Route path="/events/bookings/:id" component={eventBooking} /> 
      </div> 
    </Router> 
  }

}

ReactDOM.render(<App/>, document.getElementById("myevents-app"));