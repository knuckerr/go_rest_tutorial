import React, {useState } from 'react';
import {Form, Button, Container} from 'react-bootstrap';

function App() {
  const [Username, SetUsername] = useState("username");
  const [Password, SetPassword] = useState("password")

  return (
    <Container>
    <p> {Username} </p>
    <p> {Password} </p>
      <Form>
        <Form.Group controlId="formBasicEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control
           type="email"
           placeholder="Enter email"
           onChange = {(e => SetUsername(e.target.value))}
           />
          <Form.Text className="text-muted">
            We'll never share your email with anyone else.
          </Form.Text>
        </Form.Group>

        <Form.Group controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
          type="password"
          placeholder="Password"
          onChange = {(e => SetPassword(e.target.value))}
          />
        </Form.Group>
        <Button variant="primary" type="submit">
          Submit
        </Button>
      </Form>
    </Container>
  );
}

export default App;
