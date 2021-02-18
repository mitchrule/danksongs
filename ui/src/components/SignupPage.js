import axios from "axios";
import React, { Component } from "react";
import { Container, Form, Row, Col } from "react-bootstrap";
import Center from "react-center";
import { Link } from "react-router-dom";
import Button from "react-bootstrap/Button";

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
    e.preventDefault();
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
    {
      {
        console.log(this.state);
      }
    }
    return (
      <Container>
        <Row>
          <Col md={{ span: 6, offset: 3 }}>
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
              <Form.Group>
                <Center>
                  <Button variant="primary" type="submit">
                    Create New Account
                  </Button>
                </Center>
              </Form.Group>
            </Form>
          </Col>
        </Row>
        <p />
        <Row>
          <Col>
            <Center>
              <Link to="/">
                <Button variant="warning">Back to Login Page</Button>
              </Link>
            </Center>
          </Col>
        </Row>
      </Container>
    );
  }
}

export default SignupPage;
