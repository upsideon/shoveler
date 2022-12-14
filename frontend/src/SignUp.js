import React from 'react';
import { Navigate } from 'react-router-dom';
import PropTypes from 'prop-types';
import { createSlice, configureStore } from '@reduxjs/toolkit';

import 'regenerator-runtime/runtime';
import axios from 'axios';
import {StatusCodes as HttpStatus} from 'http-status-codes';

import PopUpModal from './PopUpModal';

function Email() {
  return (
    <div>
      <label htmlFor="email">
        <b>Email</b>
      </label>
      <input type="text" name="email" required />
    </div>
  );
}

function Password() {
  return (
    <div>
      <label htmlFor="password">
        <b>Password</b>
      </label>
      <input type="password" name="password" required />
    </div>
  );
}

function PasswordConfirm() {
  return (
    <div>
      <label htmlFor="password_confirm">
        <b>Confirm Password</b>
      </label>
      <input type="password" name="password_confirm" required />
    </div>
  );
}

function SubmitAndCancel(props) {
  const { closeOnClick } = props;
  return (
    <div className="button-grid">
      <button className="cancel-button" type="button" onClick={closeOnClick}>Cancel</button>
      <button className="submit-button" type="submit">Submit</button>
    </div>
  );
}

SubmitAndCancel.propTypes = {
  closeOnClick: PropTypes.func.isRequired,
};

const jwtSlice = createSlice({
  name: 'jwt',
  initialState: {
    token: '',
  },
  reducers: {
    added: (state, token) => {
      state.token = token
    },
  }
});

const { added } = jwtSlice.actions;

const jwtStore = configureStore({
  reducer: jwtSlice.reducer,
});

function getHeaders() {
  const token = jwtStore.getState().token.payload;
  const bearer = `Bearer ${token}`;
  return { headers: { 'Authorization': bearer } };
}

class SignIn extends React.Component {
  constructor(props) {
    super(props);
    this.props = props;
    this.state = { isLoggedIn: false };
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  async handleSubmit(event) {
    event.preventDefault();
    const email = event.target.email.value;
    const password = event.target.password.value;
    const response = await axios.post(
      'http://localhost:8080/login',
      { email, password },
    );

    if (response.status === HttpStatus.OK) {
      this.setState({ isLoggedIn: true });
      jwtStore.dispatch(added(response.data.token));
    }
  }

  render() {
    if (this.state.isLoggedIn) {
      return <Navigate to='/dashboard' replace />;
    }

    const { closeOnClick } = this.props;

    return (
      <PopUpModal id="signInModal" closeOnClick={closeOnClick}>
        <form className="modal-content" onSubmit={this.handleSubmit}>
          <div className="container">
            <div>
              <h1>Sign In</h1>
              <p> Your shoveling buddy is waiting for you! </p>
              <hr />
              <Email />
              <Password />
              <SubmitAndCancel closeOnClick={closeOnClick} />
            </div>
          </div>
        </form>
      </PopUpModal>
    );
  }
}

SignIn.propTypes = {
  closeOnClick: PropTypes.func.isRequired,
};

function SignUp(props) {
  const { closeOnClick, onSubmit } = props;
  return (
    <PopUpModal id="signUpModal" closeOnClick={closeOnClick}>
      <form className="modal-content" onSubmit={onSubmit}>
        <div className="container">
          <div>
            <h1>Sign Up</h1>

            <p> Find your shoveling buddy today. </p>
            <hr />
            <Email />
            <Password />
            <PasswordConfirm />
            <SubmitAndCancel closeOnClick={closeOnClick} />
          </div>
        </div>
      </form>
    </PopUpModal>
  );
}

SignUp.propTypes = {
  closeOnClick: PropTypes.func.isRequired,
  onSubmit: PropTypes.func.isRequired,
};

async function handleSignUp(event) {
  event.preventDefault();
  const email = event.target.email.value;
  const password = event.target.password.value;
  const password_confirm = event.target.password_confirm.value;
  const response = await axios.post(
    `http://localhost:8080/accounts`,
    { email, password, password_confirm },
  );
}

export {
  SignIn, SignUp, handleSignUp, getHeaders,
};
