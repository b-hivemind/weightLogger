import React, { useState, useEffect } from 'react'
import axios from 'axios'

import WeightViz from './components/WeightViz';
import LastWeight from './components/LastWeight';
import { WeightSender } from './components/WeightSender';
import { Login } from './components/Login';

import './App.css';

function App() {
  const [weightData, setWeightData] = useState([])

  let fetchWeightData = async () => {
    const baseURL = "http://10.0.0.184:8081/entries"
    const intervals = [2, 7, 30]
    let tempData = []
    for(let interval of intervals) {
      await axios.get(baseURL + "/" + interval)
      .then(function(response) {
        tempData.push(response.data.reverse())
      })
      .catch(function(error) {
        console.log(error)
      })
    }
    setWeightData([...tempData])
  }

  useEffect(() => {
    fetchWeightData()
  }, [])

  return(
    <div className='App'>
      <Login />
      {/*
      <div className='appContainer'>
        <div className='numbers'>
          <div className='weightSenderContainer'>
            <WeightSender stateHandler={fetchWeightData}/>
          </div>
          <div className="lastWeightContainer">
            {weightData.length > 0 && weightData[0] !== null && <LastWeight weightArr={weightData[0]}/>}
          </div>
        </div>
        <div className="deltas">
          {weightData.length > 1 && weightData[1] !== null && weightData[1].length > 2 && <WeightViz weightArr={weightData[1]}/>}
          {weightData.length > 2 && weightData[2] !== null && weightData[2].length > 7 && <WeightViz weightArr={weightData[2]}/>}
        </div>
      </div>
      */}
    </div>

  )
}

export default App;
