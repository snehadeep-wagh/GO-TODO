import React, { Component } from "react"
import axios from 'axios'
import { Card, Header, Form, Input, Icon, Button } from "semantic-ui-react"

let endPoint = "http://localhost:9000";

class ToDOList extends Component {
    constructor(props) {
        super(props);

        this.state = {
            task: "",
            items: []
        }
    }

    onSubmit = (e) => {
        e.preventDefault()
        let { task } = this.state;
        console.log("onSubmit called!: " + task)
        if (task) {
            axios.post(endPoint + "/api/createTask",
                { "Task": task },
                {
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded",
                    }
                },
            ).then((res) => {
                this.getTask();
                this.setState({
                    task: "",
                });
                console.log(res);
            }).catch((err) => {
                console.log("Error while creating task: " + err);
            });
        }
    };

    getTask = () => {
        axios.get(endPoint + "/api/task")
            .then((res) => {
                // res.data is not null update the value
                if (res.data) {
                    this.setState({
                        items: res.data.map((item) => {
                            let color = "yellow";
                            let style = {
                                wordWrap: "break-word",
                            };

                            if (item.status) {
                                color = "green";
                                style["textDecorationLine"] = "line-through";
                            }
                            else {
                                style["textDecorationLine"] = "none";
                            }

                            return (
                                <Card key={item._id} color={color} className="rough">
                                    <Card.Content>
                                        <Card.Header textAlign="left">
                                            <div style={style}>{item.task}</div>
                                        </Card.Header>

                                        <Card.Meta textAlign="right">
                                            <span style={{ paddingRight: 10 }} onClick={() => this.updateTask(item._id)}>
                                                <Icon name="check circle" color="blue" />
                                                Done</span>
                                            <span style={{ paddingRight: 10 }} onClick={() => this.undoTask(item._id)} >
                                                <Icon name="undo" color="blue" />
                                                Undo</span>
                                            <span style={{ paddingRight: 10 }} onClick={() => this.deleteTask(item._id)}>
                                                <Icon name="delete" color="red" />
                                                Delete</span>
                                        </Card.Meta>
                                    </Card.Content>
                                </Card>
                            );
                        }),
                    });
                }
                else {
                    this.setState({
                        items: [],
                    });
                }
            }).catch((err) => {
                console.log("Error while getting all tasks: " + err)
            });
    };

    undoTask = (id) => {
        axios.put(endPoint + "/api/undoTask/" + id, {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
        }).then((res) => {
            console.log(res);
            this.getTask();
        });
    };

    deleteTask = (id) => {
        axios.delete(endPoint + "/api/deleteTask/" + id, {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
        }).then((res) => {
            console.log(res);
            this.getTask();
        });
    };

    updateTask = (id) => {
        axios.put(endPoint + "/api/task/" + id, {
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
        }).then((res) => {
            console.log(res);
            this.getTask();
        });
    };

    componentDidMount() {
        this.getTask();
    }

    taskChangeHandler = (event) => {
        this.setState({
            [event.target.name]: event.target.value
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

                        {/* <Button value="submit" type="submit">Add Task</Button> */}
                    </Form>
                </div>

                <div className="row">
                    <Card.Group>{this.state.items}</Card.Group>
                </div>

            </div>
        );
    }
}

export default ToDOList