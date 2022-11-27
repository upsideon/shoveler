import { React } from 'react';
import PropTypes from 'prop-types';

import {
  SignIn, SignUp, handleSignIn, handleSignUp,
} from './SignUp';

function AccountButton(props) {
  const { onClick, text } = props;
  return (
    <button
      className="sign-in-button"
      type="button"
      onClick={onClick}
    >
      <b>{text}</b>
    </button>
  );
}

AccountButton.propTypes = {
  onClick: PropTypes.func.isRequired,
  text: PropTypes.string.isRequired,
};

function openModalClick(modalId) {
  return () => {
    const modal = document.getElementById(modalId);
    modal.style.display = 'block';
  };
}

function closeModalClick(modalId) {
  return () => {
    const modal = document.getElementById(modalId);
    modal.style.display = 'none';
  };
}

function Navigation() {
  return (
    <div className="navigation">
      <img src="images/shoveler_logo.png" className="title" alt="Shoveler logo." />
      <AccountButton text="Sign In" onClick={openModalClick('signInModal')} />
      <AccountButton text="Create Account" onClick={openModalClick('signUpModal')} />
      <SignIn closeOnClick={closeModalClick('signInModal')} onSubmit={handleSignIn} />
      <SignUp closeOnClick={closeModalClick('signUpModal')} onSubmit={handleSignUp} />
    </div>
  );
}

export default Navigation;
