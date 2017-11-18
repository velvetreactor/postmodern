import React from 'react';
import Tables from './Tables.jsx';
import Rows from './Rows.jsx';

class App extends React.Component {
  render() {
    return (
      <main className="app container">
        <div className="row component-container">
          <Tables />
          <Rows />
        </div>
      </main>
    )
  }
}

export default App;
