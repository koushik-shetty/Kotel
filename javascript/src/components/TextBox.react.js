import React from "react";

class TextBox extends React.Component{
	Constructor(props){
		Super(props);
		console.log(this.state);
	}

	render(){
		return(
			<div className="input-box">
					<input type={this.props.type === undefined ? "text":this.props.type} onChange={function() {}} name={this.props.Name} placeholder={this.props.FieldName} required/>
			</div>
		);
	}
}

export default TextBox;