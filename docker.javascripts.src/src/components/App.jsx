import React from 'react';
import Tables from './Tables.jsx';
import Rows from './Rows.jsx';
import CredentialsModal from './CredentialsModal.jsx';

class App extends React.Component {
  render() {
    return (
      <main className="app container">
        <div className="row component-container">
          <Tables />
          <Rows />
        </div>
        <CredentialsModal />
      </main>
    )
  }
}

export default App;
