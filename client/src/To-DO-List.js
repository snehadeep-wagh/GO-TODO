import React, { Component } from "react"
import axios from 'axios'
import { Card, Header, Form, Input, Icon } from "semantic-ui-react"

let endPoint = "http://localhost:9000";

class ToDOList extends Component {
    constructor(props) {
        super(props);

        this.state = {
            task: "",
            items: []
        }
    }

    getTask(){

    }

    componentDidMount() {
        this.getTask();
    }

    taskChangeHandler = (event)=>{
        this.setState({
            task: event.target.value
        })
    }

    render() {
        return (
            <div>
                <div className="row">
                    <Header className="header" color="yellow">To Do List App</Header>
                </div>
                <div className="row">
                    <Form onSubmit={this.onSubmit}>
                        <Input
                        type="text"
                        name="task"
                        value={this.state.task}
                        onChange={this.taskChangeHandler}
                        placeholder="create task"
                        fluid
                        />
                    </Form>
                </div>
            </div>
        );
    }
}

export default ToDOList