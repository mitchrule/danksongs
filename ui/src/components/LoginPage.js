import axios from "axios";
import React, { Component } from "react";
import { Container, Form, Button, Row, Col } from "react-bootstrap";
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
        <Container fluid="md">
          <Row>
            <Col>
            <Center>
            <h1>Dank Songs</h1>
            </Center>
            </Col>
          </Row>
          <Row>
           <Col md={{ span: 6, offset: 3 }}>
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
            </Col>
          </Row>
          <Row>
            <Col>
            <Center>
              <Button className="display-btn" variant="success" type="submit">
                Login
              </Button>
            </Center>
            </Col>
          </Row>
          <p />
          <Row>
            
            <Col>
            <Center>
            <Link to="/signup">
              <Button variant="primary">Go To Sign Up Page</Button>
            </Link>
            </Center>
            </Col>
           
          </Row>
          <Row>
            
            <Col>
            <Center>
            <Link to="/playlist">
              <Button variant="primary">View Playlist</Button>
            </Link>
            </Center>
            </Col>
            
          </Row>
        </Container>

    );
  }
}

export default LoginPage;
