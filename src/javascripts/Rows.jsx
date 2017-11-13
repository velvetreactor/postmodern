import React from 'react';
import request from 'superagent';
import { connect } from 'react-redux';

class Rows extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      rows: []
    }
  }

  componentWillMount() {
    // First displayed table is Rings
    request.get('/rows?table=rings').end((err, res) => {
      this.setState({ rows: res.body });
    });
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevProps.tableName == this.props.tableName) {
      return false;
    }
    let url = `/rows?table=${this.props.tableName}`
    request.get(url).end((err, res) => {
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

function mapStateToProps(state) {
  return {
    tableName: state.tableState.tableName
  }
}

export default connect(mapStateToProps)(Rows);
