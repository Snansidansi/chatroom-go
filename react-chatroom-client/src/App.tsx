import { useState } from "react";

function App() {
  const [message, setMessage] = useState("");

  const sendMessage = () => {
    if (message.trim() === "") {
      return;
    }

    console.log(message);
    setMessage("");
  };

  const handleEnter = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      sendMessage();
    }
  };

  return (
    <div id="main-container">
      <h1>Chatroom</h1>
      <ul className="chat">
        <li>name: message</li>
      </ul>
      <div id="user-input-div">
        <input
          id="chat-input"
          type="text"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={handleEnter}
        />
        <button id="send-button" onClick={sendMessage}>
          Send
        </button>
      </div>
    </div>
  );
}

export default App;
