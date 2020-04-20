import * as React from "react";
import { Event } from "../models/event";
import { Link } from "react-router-dom";

export interface EventListItemProps {
    event: Event;
}

export class EventListItem extends React.Component<EventListItemProps, {}> {
    render() {
        const start = new Date(this.props.event.start_date)
        const end = new Date(this.props.event.end_date)

        return <tr>
            <td>{this.props.event.name}</td>
            <td>{this.props.event.location.name}</td>
            <td>{start.toDateString()} - {end.toDateString()}</td>
            <td>
                <Link to={`/events/bookings/${this.props.event.ID}`}>
                    Book now!
                </Link>
            </td>
        </tr>
    }
}

