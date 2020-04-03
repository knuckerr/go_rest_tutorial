import React, {useState, useContext} from 'react';
import {Form, Button, Container} from 'react-bootstrap';
import UserContext from "../../contexts/UserContext";

const Login = () => {
  const [Username, SetUsername] = useState('');
  const [Password, SetPassword] = useState('');
  const [state, dispatch] = useContext(UserContext)

  const HandleSubmit = (e) => {
    e.preventDefault();
    const Json = {
      "password": Password,
      "username": Username,
    };
  };

  return (
    <Container>
      <Form>
        <Form.Group controlId="formBasicEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control
            type="email"
            placeholder="Enter email"
            onChange={e => SetUsername(e.target.value)}
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
            onChange={e => SetPassword(e.target.value)}
          />
        </Form.Group>
        <Button
         variant="primary"
         onClick={e => HandleSubmit(e)}
         type="submit">
          Submit
        </Button>
      </Form>
    </Container>
  );
};

export default Login;
