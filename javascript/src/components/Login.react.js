import React from "react";
import TextBox from "./TextBox.react";

class Login extends React.Component{
	render(){
		return (
			<div>
				<TextBox DefaultText="" Name="userid" FieldName="UserID"/>
				<TextBox DefaultText="" Name="password" FieldName="Password"/>
			</div>
		);
	}
}

//React.render(<Login />,document.getEleementById("root"));
export default Login;