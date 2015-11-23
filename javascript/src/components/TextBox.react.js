import React from "react";

class TextBox extends React.Component{
	Constructor(props){
		Super(props);
		console.log(this.state);
	}

	render(){
		return(
			<div className="input-box">
				<form>
					<input type={this.props.type === undefined ? "text":this.props.type} onChange={function() {}} name={this.props.Name} placeholder={this.props.FieldName} required/>
				</form>
			</div>
		);
	}
}

export default TextBox;