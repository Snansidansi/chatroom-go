import { useState } from "react";

type ConnectPopupProps = {
  show: boolean;
  onConnect: (ip: string, name: string) => void;
};

function ConnectPopup(props: ConnectPopupProps) {
  const [ip, setIp] = useState("");
  const [userName, setUserName] = useState("");

  if (!props.show) {
    return null;
  }

  const handleConnect = () => {
    if (!ip.trim() || !userName.trim()) {
      alert("Bitte IP mit Port und Name eingeben!");
      return;
    }

    props.onConnect(ip, userName);
  };

  return (
    <div className="popup">
      <div className="popup-inner">
        <h1>Connect to Chatroom</h1>
        <div className="popup-input">
          <label htmlFor="ip">IP with Port:</label>
          <input
            type="text"
            id="ip"
            value={ip}
            onChange={(e) => setIp(e.target.value)}
          />
          <label htmlFor="name">Name:</label>
          <input
            type="text"
            id="name"
            value={userName}
            onChange={(e) => setUserName(e.target.value)}
          />
          <button onClick={handleConnect}>Connect</button>
        </div>
      </div>
    </div>
  );
}

export default ConnectPopup;
