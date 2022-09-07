import React from 'react'
import moment from 'moment'
import { FaCaretUp, FaCaretDown } from 'react-icons/fa'

const LastWeight = ({ weightArr }) => {
    var delta = () => {
      if (weightArr.length < 1) {
        return 0
      }
      return (weightArr[0].weight - weightArr[weightArr.length - 1].weight).toFixed(2)
    }

    return (
    <div className='lastWeightComponent'>
      <div className='lastWeight'>
        {weightArr != null && <h1 className='mainText deltaText'>{weightArr[0].weight}</h1>}
        {weightArr != null && <p className='deltaText'>{weightArr[0].date}</p>}
      </div>
      <div className='lastWeightDelta'>
        <h1 className={`${delta() > 0 ? 'gain' : 'loss'}`}>
            {Math.abs(delta())}lb&nbsp;
            {delta() < 0 && <FaCaretDown/>}
            {delta() > 0 && <FaCaretUp/>}
        </h1>
        <p className='deltaText'>since {weightArr != null && moment(weightArr[weightArr.length - 1].date).fromNow()}</p>
      </div>

    </div>
   )
}

export default LastWeight