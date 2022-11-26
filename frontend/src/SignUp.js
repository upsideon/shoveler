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
			<label htmlFor="password-confirm">
				<b>Confirm Password</b>
			</label>
			<input type="password" name="password-confirm" required />
		</div>
	);
}

function SubmitAndCancel(props) {
	return (
		<div className="button-grid">
			<button className="cancel-button" type="button" onClick={props.closeOnClick}>Cancel</button>
			<button className="submit-button" type="submit">Submit</button>
		</div>
	);
}

function PopUpModal(props) {
	return (
		<div className="modal" id={props.id}>
		  <span className="close" onClick={props.closeOnClick} title="Close">
		    &times;
		  </span>
		  {props.children}
		</div>
	);
}

function SignIn(props) {
	return (
		<PopUpModal id="signInModal" closeOnClick={props.closeOnClick}>
		  <form className="modal-content" onSubmit={props.onSubmit}>
		    <div className="container">
					<div>
						<h1>Sign In</h1>
						<p> Your shoveling buddy is waiting for you! </p>
						<hr/>
						<Email />
						<Password />
						<SubmitAndCancel closeOnClick={props.closeOnClick}/>
					</div>
		    </div>
		  </form>
		</PopUpModal>
	);
}

function handleSignIn(event) {
	event.preventDefault();
	alert("Sign In");
}

function SignUp(props) {
	return (
		<PopUpModal id="signUpModal" closeOnClick={props.closeOnClick}>
		  <form className="modal-content" onSubmit={props.onSubmit}>
		    <div className="container">
					<div>
						<h1>Sign Up</h1>

						<p> Find your shoveling buddy today. </p>
						<hr/>
						<Email />
						<Password />
						<PasswordConfirm />
						<SubmitAndCancel closeOnClick={props.closeOnClick}/>
					</div>
		    </div>
		  </form>
		</PopUpModal>
	);
}

function handleSignUp(event) {
	event.preventDefault();
	alert("Sign Up!");
}

export { SignIn, SignUp, handleSignIn, handleSignUp};
