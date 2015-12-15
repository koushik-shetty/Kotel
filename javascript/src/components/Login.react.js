import React from "react";
import TextBox from "./TextBox.react";

class Login extends React.Component{
	render(){
		return (
			<div>
				<form>
					<TextBox DefaultText="" Name="userid" FieldName="UserID"/>
					<TextBox DefaultText="" Name="password" FieldName="Password" type="password"/>
					<input type="submit" formMethod="post" name="login" value="Login" formAction="http://localhost:5004/login" />
				</form>
			</div>
		);
	}
}

export default Login;