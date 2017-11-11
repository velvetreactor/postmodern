import React from 'react';
import Tables from './Tables.jsx';
import Rows from './Rows.jsx';

class App extends React.Component {
  render() {
    return (
      <main className="app">
        <Tables />
        <Rows />
      </main>
    )
  }
}

export default App;
