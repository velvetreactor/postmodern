import React from 'react';
import request from 'superagent';

class Rows extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      rows: []
    }
  }

  componentWillMount() {
    request.get('/rows?table=rings').end((err, res) => {
      this.setState({ rows: res.body });
    });
  }

  render() {
    let items = this.state.rows.map(row => {
      return (
        <tr>
          <td>{row.id}</td>
          <td>{row.name}</td>
          <td>{row.description}</td>
        </tr>
      )
    });
    return (
      <section className='table-rows col-lg-10'>
        <table>
          <thead>
            <tr>
              <th>id</th>
              <th>name</th>
              <th>description</th>
            </tr>
          </thead>
          <tbody>
            {items}
          </tbody>
        </table>
      </section>
    )
  }
}

export default Rows;
