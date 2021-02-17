/*import axios from "axios";
import React, { Component } from "react";
import {
  Container,
  Form,
  Button,
  Row,
} from "react-bootstrap";
import Center from "react-center";
import { Link } from "react-router-dom";

class LoginPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
    };
  }

  onChange = (e) => {
    this.setState({
      [e.target.id]: e.target.value,
    });
  };

  onSubmit = (e) => {
    e.preventDefault();

    this.setState({
      waiting: true,
    });

    const userData = {
      username: this.state.username,
      password: this.state.password,
    };
    axios
      .post("/api/login", userData)
      .then((res) => {
        console.log("Signup success with data =", res);
        this.props.history.push("/vote");
      })
      .catch((err) => {
        console.log("Signup failed, Errors:");
        console.log(err);
      });
  };

  render() {
    //{{console.log(this.state)}}
    return (
      <Container>
        <Center>
          <Row>
            <Form onSubmit={this.onSubmit}>
              <Form.Group controlId="username">
                <Form.Label>Username</Form.Label>
                <Form.Control
                  type="text"
                  placeholder="Username"
                  onChange={this.onChange}
                />
              </Form.Group>
              <Form.Group controlId="password">
                <Form.Label>Password</Form.Label>
                <Form.Control
                  type="password"
                  placeholder="Password"
                  onChange={this.onChange}
                />
              </Form.Group>
            </Form>
            <Button className="display-btn" variant="primary" type="submit">
              Login
            </Button>
          </Row>
          <Row>
            <Link to="/signup">
              <Button>Go To Sign Up Page</Button>
            </Link>
          </Row>
          <Link to="/playlist">
            <Row>
              <Button>View Playlist</Button>
            </Row>
          </Link>
        </Center>
      </Container>
    );
  }
}

export default LoginPage;
*/

import React, { Component } from "react";
import { Button } from "react-bootstrap";

class LoginPage extends Component {
  render() {
    return (<Button variant="primary">Primary</Button>);
  }
}

export default LoginPage;

