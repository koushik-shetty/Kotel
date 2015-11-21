import React from "react";

class TextBox extends React.Component{
	Constructor(props){
		Super(props);
		console.log(this.state);
	}

	render(){
		return(
			<div className="input-box">
				{this.props.FieldName + ": "}	<input type="text" boxName={this.props.Name}/>
			</div>
		);
	}
}

export default TextBox;