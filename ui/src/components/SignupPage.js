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
    
        this.setState({
          waiting: false,
        });
      };
    
      onSubmit = (e) => {
        e.preventDefault();
    
        this.setState({
          waiting: true,
        });
    
        const userData = {
          email: this.state.newemail,
          password: this.state.newpassword,
        };
        axios
        .post("/api/signup", userData)
        .then((res) => {
            console.log("Signup success with res.data =", res.data);
            history.push("/");
            })
        .catch((err) => {
        console.log("Signup failed, Errors:");
        console.log(err);   
        });
      };

  render() {
    return (
      <Container>
        <Center>
          <Row>
            <Form>
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
          <Link to="/">
            <Button className="display-btn" variant="primary" type="submit">
                Create New Account
            </Button>
        </Link>
          </Row>
        </Center>
      </Container>
    );
  }
}

export default SignupPage;
