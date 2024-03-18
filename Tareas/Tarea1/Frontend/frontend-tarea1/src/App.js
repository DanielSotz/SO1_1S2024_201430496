import React, { useState } from "react";
import "./App.css";

function App() {
  const [userData, setUserData] = useState(null);

  const fetchData = async () => {
    try {
      const response = await fetch("http://localhost:8080/data");
      const data = await response.json();
      setUserData(data);
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>React-Go App</h1>
        <button onClick={fetchData}>Fetch User Data</button>
        {userData && (
          <div>
            <p>Nombre: {userData.name}</p>
            <p>Carnet: {userData.carnet}</p>
            <p>Timestamp: {userData.timestamp}</p>
          </div>
        )}
      </header>
    </div>
  );
}

export default App;
