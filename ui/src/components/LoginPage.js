import React, { Component } from "react";
import { Container, Form, Card, Badge, Col, Button, Row } from "react-bootstrap";

class LoginPage extends Component {

    render() {
        return (
          <Container>
              <Button>
                  Login
              </Button>
              <Button>
                  View Playlist
              </Button>

          </Container>
        );
    }
}

export default LoginPage;