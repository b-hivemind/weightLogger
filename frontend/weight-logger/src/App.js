import React, { useState, useEffect } from 'react'
import axios from 'axios'

import WeightViz from './components/WeightViz';
import LastWeight from './components/LastWeight';
import { WeightSender } from './components/WeightSender';

import './App.css';

function App() {
  const [data, setData] = useState(null)

  const getData = async () => {
    const queryURL = "http://10.0.0.220:10000/dashboard"
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
            {data !== null && <LastWeight weightArr={data[0]}/>}
          </div>
        </div>
        <div className="deltas">
          {data !== null && data.slice(1, data.length).map((weightArr) => <WeightViz key={data.indexOf(weightArr)} weightArr={weightArr}/>)}
        </div>
      </div>
    </div>
  );
}

export default App;
