import { useRef, useState } from "react";
import type { MessageProps } from "./message";
import Message from "./message";
import ConnectPopup from "./connectPopup";

type Msg = {
  Name: string;
  Message: string;
};

function App() {
  const [connected, setConnected] = useState(false);
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState<MessageProps[]>([]);

  const serverIpRef = useRef<string>("");
  const userNameRef = useRef<string>("");
  const chatWSRef = useRef<WebSocket | null>(null);

  const addMessage = (name: string, text: string) => {
    setMessages((prev) => [...prev, { name, text }]);
  };

  const sendMessage = () => {
    if (message.trim() === "") {
      return;
    }

    if (chatWSRef.current === null) {
      return;
    }

    const msg: Msg = {
      Name: userNameRef.current,
      Message: message,
    };

    chatWSRef.current.send(JSON.stringify(msg));

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

    chatWSRef.current = new WebSocket(`ws://${serverIpRef.current}/chat`);

    chatWSRef.current.onerror = () => {
      setConnected(false);
      return;
    };

    chatWSRef.current.onmessage = (event) => {
      const msg: Msg = JSON.parse(event.data);

      if (msg.Name === userNameRef.current) {
        return;
      }
      addMessage(msg.Name, msg.Message);
    };

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
