import React, { Component } from "react";
import "./App.css"
import { connect, sendMsg } from "./api";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";

class App extends Component {

    constructor(props) {
        super(props);
        this.state = {
            chatHistory: [] // define empty chatHistory starting state
        }
    }

    //life-cycle method
    componentDidMount() {
        connect( (msg) => {
            console.log("New Message")
            this.setState(prevState => ({
                chatHistory: [...this.state.chatHistory, msg] // update chatHistory state
            })) // Add new message to chatHistory Array immutably.
            console.log(this.state);
        });
    }

    send(event) {
        if(event.keyCode === 13) { // Enter key  
            sendMsg(event.target.value);
            event.target.value = "";
        }

        // console.log("hello front end");
        // sendMsg("Hello backend");
    }

    render() {
        return (
            <div className="App">
                <Header />
                <ChatHistory chatHistory={this.state.chatHistory} /> {/*display ChatHistory component*/}
                <ChatInput send={this.send} />
                {/* <button onClick={this.send}>Hit</button> sends hello message */}
            </div>
        );
    }

}

export default App;
