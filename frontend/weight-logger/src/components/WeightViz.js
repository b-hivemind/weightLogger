import WeightGraph from "./WeightGraph"
import { FaCaretUp, FaCaretDown } from 'react-icons/fa'
import React from 'react'


const WeightViz = ({weightArr}) => {    

  const delta = () => {
    return (weightArr[0].weight - weightArr[weightArr.length - 1].weight).toFixed(2)
  }

  return (
  <div className='weightViz'>
    <div className='weightDelta'>
      <div className='deltaText'><h1>{weightArr.length}-day </h1></div>
      <div className='deltaNum'>
        <h1 className={`${delta() > 0 ? 'gain' : 'loss'}`}> 
          {Math.abs(delta())}lb&nbsp;
          {delta() < 0 && <FaCaretDown/>}
          {delta() > 0 && <FaCaretUp/>}
        </h1>
      </div>
    </div>
    <div className='weightGraph'>
      {weightArr.length > 2 && <WeightGraph delta={delta()} weightGraphData={weightArr}/>}    
    </div>
  </div>
  )   
}

export default WeightViz