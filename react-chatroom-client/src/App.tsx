import { useRef, useState } from "react";
import type { MessageProps } from "./message";
import Message from "./message";
import ConnectPopup from "./connectPopup";

function App() {
  const [connected, setConnected] = useState(false);
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState<MessageProps[]>([]);

  const serverIpRef = useRef<string>("");
  const userNameRef = useRef<string>("");

  const addMessage = (name: string, text: string) => {
    setMessages((prev) => [...prev, { name, text }]);
  };

  const sendMessage = () => {
    if (message.trim() === "") {
      return;
    }

    addMessage(userNameRef.current, message);
    setMessage("");
  };

  const handleEnter = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      sendMessage();
    }
  };

  const onConnect = (ip: string, userName: string) => {
    serverIpRef.current = ip;
    userNameRef.current = userName;
    setConnected(true);
  };

  return (
    <div id="main-container">
      <ConnectPopup show={!connected} onConnect={onConnect} />
      <h1>Chatroom</h1>
      <ul className="chat">
        {messages.map((msg, index) => (
          <li key={index}>
            <Message name={msg.name} text={msg.text} />
          </li>
        ))}
      </ul>
      <div id="user-input-div">
        <input
          type="text"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={handleEnter}
        />
        <button onClick={sendMessage}>Send</button>
      </div>
    </div>
  );
}

export default App;
