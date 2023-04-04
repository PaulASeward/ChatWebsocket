import React, { Component } from "react";
import "./ChatInput.scss"

class ChatInput extends Component {
    render() {
        return (
            // <div className="ChatInput">
            //     <input onKeyDown={this.props.send} />
            // </div>
            <div className="ChatInput">
            <input
              type="text"
              placeholder="Type a message... Hit enter to send"
              onKeyDown={this.props.send}
            />
          </div>
        );
    }
}

export default ChatInput;