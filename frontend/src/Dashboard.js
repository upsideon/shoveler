import React from 'react';

import Beacons from './Beacons';

import './Dashboard.css';

function Dashboard() {
  return (
    <div className="dashboard">
      <div className="dashboard-header">
        <h1>Dashboard</h1>
        <hr/>
      </div>
      <Beacons />
    </div>
  );
}

export default Dashboard;
