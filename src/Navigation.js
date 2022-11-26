import { SignIn, SignUp, handleSignIn, handleSignUp } from './SignUp';

function AccountButton(props) {
	return (
		  <button className="sign-in-button" onClick={props.onClick}><b>{props.text}</b></button>
	);
}

function Navigation() {
	const openModalClick = function (modalId) {
		return () => {
			const modal = document.getElementById(modalId);
			modal.style.display = 'block';
		};
	};

	const closeModalClick = function (modalId) {
		return () => {
			const modal = document.getElementById(modalId);
			modal.style.display = 'none';
		};
	};

	return (
		<div className="navigation">
	    <img src="images/shoveler_logo.png" className="title" alt="Shoveler logo."/>
		  <AccountButton text="Sign In" onClick={openModalClick("signInModal")}/>
		  <AccountButton text="Create Account" onClick={openModalClick("signUpModal")} />
		  <SignIn closeOnClick={closeModalClick("signInModal")} onSubmit={handleSignIn}/>
		  <SignUp closeOnClick={closeModalClick("signUpModal")} onSubmit={handleSignUp}/>
		</div>
	);
}


export { Navigation };
