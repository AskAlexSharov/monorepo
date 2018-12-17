import React, {useState} from 'react';
import Room from './Room';

function App() {
  const [name, setName] = useState("");
  const [channel, setChannel] = useState("default");

  return <div>
    name: <br />
    <input value={name} onChange={(e) => setName(e.target.value)} /> <br />

    channel: <br />
    <input value={channel} onChange={(e) => setChannel(e.target.value)} /> <br />

    <Room channel={channel} name={name} />
  </div>
}


export default App;
