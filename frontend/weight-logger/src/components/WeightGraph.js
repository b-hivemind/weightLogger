import React from 'react'
import { Chart as ChartJS } from 'chart.js/auto'
import { Chart, Line } from 'react-chartjs-2'



const WeightGraph = ({ delta, weightGraphData }) => {
  const options = {
    responsive: true,
    plugins: {
      legend: {
        display: false
      },
      title: {
        display: false,
        text: weightGraphData.length + ' day weight',
      },
    },
    scales: {
      x: {
        grid: {
          display: false,
        },
        ticks: {
          color: '#a5adcb'
        },
      },
      y: {
        grid: {
          display: false,
        },
        ticks: {
          color: '#a5adcb'
        },
      },
    },
  }

  const data = {
    labels: (weightGraphData.map((item) => item.date)).reverse(),
    datasets: [{
      label: 'weight',
      data: (weightGraphData.map((item) => item.weight)).reverse(),
      backgroundColor: '#494d64',
      borderColor: delta > 0 ? '#ed8796' : '#a6da95',
      color: delta > 0 ? '#ed8796' : '#a6da95',
      pointRadius: 1,
    }]
  }
  
  return (
    <div className={`weightGraphLineChart`}>
      <Line
        options={options}
        data={data}
      />
    </div>
  )
}

export default WeightGraph
