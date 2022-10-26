import React, { useState, useEffect } from 'react'

import axios from 'axios'

import WeightViz from './WeightViz';
import LastWeight from './LastWeight';
import { WeightSender } from './WeightSender';

const Dashboard = ({ user, setUser }) => {
    const [weightData, setWeightData] = useState([])

    let fetchWeightData = async () => {
      const baseURL = "http://10.0.0.184:8081/entries"
      const intervals = [2, 7, 30]
      let tempData = []
      for(let interval of intervals) {
        await axios.get(baseURL + "/" + interval, {
          headers: {
            bearer: user["token"],
          },
        })
        .then(function(response) {
          tempData.push(response.data)
        })
        .catch(function(error) {
          console.log(error)
          window.localStorage.removeItem("token")
          window.localStorage.removeItem("username")
        })
      }
      setWeightData([...tempData])
    }

    useEffect(() => {
      fetchWeightData()
    }, [])
  
    return (
        <div className='appContainer'>
            <div className="greeting">
              <h2 className="title">Hi, {user["username"]}!</h2>
            </div>
            <div className='numbers'>
                <div className='weightSenderContainer'>
                    <WeightSender token={user["token"]} stateHandler={fetchWeightData}/>
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
  )
}

export default Dashboard