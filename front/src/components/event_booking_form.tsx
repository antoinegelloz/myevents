import * as React from "react";
import { Event } from "../models/event";
import { FormRow } from "./form_row";
import {ChangeEvent} from "react";

export interface EventBookingFormProps {
    event: Event;
    onSubmit: (quantity: number) => any
}

export interface EventBookingFormState {
    quantity: number;
}

export class EventBookingForm extends React.Component<EventBookingFormProps, EventBookingFormState> {
    constructor(p: EventBookingFormProps) {
        super(p);
        this.state = { quantity: 1 };
    }

    render() {
        return <div>
            <h2>Book tickets for {this.props.event.name}</h2>
            <form className="form-horizontal">
                <FormRow label="Event">
                    <p className="form-control-static">
                        {this.props.event.name}
                    </p>
                </FormRow>
                <FormRow label="Number of tickets">
                    <select className="form-control" value={this.state.quantity}
                        onChange={event => this.handleNewAmount(event)}>
                        <option value="1">1</option>
                        <option value="2">2</option>
                        <option value="3">3</option>
                        <option value="4">4</option>
                    </select>
                </FormRow>
                <FormRow>
                    <button className="btn btn-primary"
                        onClick={() => this.props.onSubmit(this.state.quantity)}>
                        Submit order
                    </button>
                </FormRow>
            </form>
        </div>
    }

    private handleNewAmount(event: ChangeEvent<HTMLSelectElement>) { 
        const state: EventBookingFormState = { 
          quantity: parseInt(event.target.value) 
        } 
        this.setState(state); 
      }
}