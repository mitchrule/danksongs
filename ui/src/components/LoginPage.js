import axios from "axios";
import React, { Component } from "react";
import {
  Container,
  Form,
  Card,
  Badge,
  Col,
  Button,
  Row,
} from "react-bootstrap";
import Center from 'react-center';
import { Link } from "react-router-dom";

class LoginPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
    };
  }

  render() {
    //{{console.log(this.state)}}
    return (
      <Container>
        <Center>
          <Row>
            <Form>
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
            <Link to="/vote">
            <Button className="display-btn" variant="primary" type="submit">
              Login
            </Button>
            </Link>
          </Row>
          <Row>
          <Link to="/signup">
            <Button>
                Go To Sign Up Page
            </Button>
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
