import { useState } from "react";
import './App.css';
import Button from 'react-bootstrap/Button';
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  const [number, setNumber] = useState(1);
  const [message, setMessage] = useState(null);

  const getMessage = async () => {
    const BACKEND_URL = process.env.REACT_APP_BACKEND_URL as string;
    
    await fetch(BACKEND_URL+"/fizzbuzz", {
      method: 'POST',
      body: JSON.stringify({
        count: number+1,
      }),
      headers: {
        'Content-type': 'application/json; charset=UTF-8',
      },
    })
    .then((response) => response.json())
    .then((data) => {
      setMessage(data.message);
    })
    .catch((err) => {
      console.log(err.message);
    });
  }

  const handleClick = async () => {
    await getMessage();
    setNumber(number+1);
  }

  return (
    <div className="App">
      <header className="App-header">
        <label>
          Your count
        </label>
        <p>
          {number}
        </p>
        <Button 
          className='shadow'
          variant="primary"
          size="lg"
          onClick={() => handleClick()}
        >
          Push me!
        </Button>

        <h1 style={{margin: 50}}>
          {message}
        </h1>
      </header>
    </div>
  );
}

export default App;
