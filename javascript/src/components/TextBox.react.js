import React from "react";

class TextBox extends React.Component{
	Constructor(props){
		Super(props);
		this.setState({
			text: props.DefaultText,
			boxName: props.Name,
			FieldName: props.FieldName
		});
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