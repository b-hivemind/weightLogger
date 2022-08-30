import React, { useState, useEffect } from 'react'
import axios from 'axios'

import WeightViz from './components/WeightViz';
import LastWeight from './components/LastWeight';
import { WeightSender } from './components/WeightSender';

import './App.css';

function App() {
  const [data, setData] = useState(null)

  const getData = async () => {
    const queryURL = "http://10.0.0.228:10000/dashboard"
    const results = await axios.get(queryURL)
    setData(results.data)
  }

  useEffect(() => {
      getData()
    }, [])

  return (
    <div className='App'>
      <div className='appContainer'>
        <div className='numbers'>
          <div className='weightSenderContainer'>
            <WeightSender stateHandler={getData}/>
          </div>
          <div className="lastWeightContainer">
            {data !== null && data[0] !== null && <LastWeight weightArr={data[0]}/>}
          </div>
        </div>
        <div className="deltas">
          {data !== null && data[1] !== null && data[1].length > 1 && <WeightViz weightArr={data[1]}/>}
          {data !== null && data[2] !== null && data[2].length > 7 && <WeightViz weightArr={data[2]}/>}
        </div>
      </div>
    </div>
  );
}

export default App;
