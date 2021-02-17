import axios from "axios";
import React, { Component } from "react";
import {
  Container,
  Form,
  Row,
} from "react-bootstrap";
import Center from "react-center";
import { Link } from "react-router-dom";
import Button from 'react-bootstrap/Button';

class SignupPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      newusername: "",
      newpassword: "",
    };
  }

  onChange = (e) => {
    this.setState({
      [e.target.id]: e.target.value,
    });

    console.log(this.state);
  };

  onSubmit = (e) => {
    console.log("Submitting");

    e.preventDefault();

    this.setState({
      waiting: true,
    });

    const userData = {
      newusername: this.state.newusername,
      newpassword: this.state.newpassword,
    };

    console.log("Submitting Data...");
    axios
      .post("/api/signup", userData)
      .then((res) => {
        console.log("Signup success with data =", res);
        this.props.history.push("/");
      })

      .catch((err) => {
        console.log("Signup failed, Errors:");
        console.log(err);
      });
  };

  render() {

    {{console.log(this.state)}}
    return (
      <Container>
        <Center>
          <Row>
            <Form onSubmit={this.onSubmit}>
              <Form.Group controlId="newusername">
                <Form.Label>New Username</Form.Label>
                <Form.Control
                  type="text"
                  placeholder="Username"
                  onChange={this.onChange}
                  defaultValue={this.state.newusername}
                />
              </Form.Group>
              <Form.Group controlId="newpassword">
                <Form.Label>New Password</Form.Label>
                <Form.Control
                  type="password"
                  placeholder="Password"
                  onChange={this.onChange}
                  defaultValue={this.state.newpassword}
                />
              </Form.Group>
            </Form>
          </Row>
          <Row>
            <Button className="display-btn" variant="primary" onclick="submit">
              Create New Account
            </Button>
          </Row>
        </Center>
      </Container>
    );
  }
}

export default SignupPage;
