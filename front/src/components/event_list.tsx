import {Event} from "../models/event"; 
import {EventListItem} from "./event_list_item"; 
import * as React from "react"; 
 
export interface EventListProps { 
  events: Event[]; 
} 
 
export class EventList extends React.Component<EventListProps, {}> { 
  render() {
    var items:JSX.Element[] = null
    if (this.props.events) {
      items = this.props.events.map(e => <EventListItem event={e} />);
    }
 
    return <table className="table"> 
      <thead>
        <tr> 
          <th>Event</th> 
          <th>Where</th> 
          <th>When</th> 
          <th>Booking</th> 
        </tr> 
      </thead> 
      <tbody> 
        {items} 
      </tbody> 
    </table> 
  }   
}