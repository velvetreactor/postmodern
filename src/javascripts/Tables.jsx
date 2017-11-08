import React from 'react';
import request from 'superagent';

class Tables extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      tableNames: []
    }
  }

  componentDidMount() {
    request.get('/tables').end((err, res) => {
      this.setState({ tableNames: res.body });
    });
  }

  render() {
    console.log(this.state)
    let tableItems = this.state.tableNames.map(name => {
      return <li>{name}</li>
    })
    return (
      <div className='tables'>
        <ul>
          {tableItems}
        </ul>
      </div>
    )
  }
}

export default Tables;
