import React from 'react'
import moment from 'moment'
import { FaCaretUp, FaCaretDown } from 'react-icons/fa'

const LastWeight = ({ weightArr, units }) => {
    var delta = () => {
      if (weightArr.length < 1) {
        return 0
      }
      return (weightArr[0].weight - weightArr[weightArr.length - 1].weight).toFixed(2)
    }

    let lastWeight = weightArr[0].weight, lastDate = weightArr[0].date;

    return (
    <div className='lastWeightComponent'>
      <div className='lastWeight'>
        {weightArr != null && <h1 className='mainText deltaText'>{lastWeight}</h1>}
        {weightArr != null && <p className='deltaText'>{moment.unix(lastDate).format('MMM D \'YY HH:MM')}</p>}
      </div>
      {weightArr.length > 1 &&
        <div className='lastWeightDelta'>
          <h1 className={`${delta() > 0 ? 'gain' : 'loss'}`}>
              {Math.abs(delta())}{units.toLowerCase()}&nbsp;
              {delta() < 0 && <FaCaretDown/>}
              {delta() > 0 && <FaCaretUp/>}
          </h1>
          <p className='deltaText'>since {weightArr != null && moment.unix(weightArr[weightArr.length - 1].date).fromNow()}</p>
        </div>
      }
    </div>
   )
}

export default LastWeight