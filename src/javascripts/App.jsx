import React from 'react';
import Tables from './Tables.jsx';
import Rows from './Rows.jsx';

class App extends React.Component {
  render() {
    return (
      <main className="app container-fluid">
        <div className="row">
          <Tables />
          <Rows />
        </div>
      </main>
    )
  }
}

export default App;
