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
    let tableItems = this.state.tableNames.map(name => {
      return <li>{name}</li>
    })

    return (
      <section className='tables col-lg-2'>
        <ul>
          {tableItems}
        </ul>
      </section>
    )
  }
}

export default Tables;
