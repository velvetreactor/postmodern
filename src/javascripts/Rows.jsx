import React from 'react';
import request from 'superagent';
import { connect } from 'react-redux';

import QueryBox from './QueryBox.jsx';

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
    if (prevProps == this.props) {
      return false;
    }

    if (prevProps.tableName != this.props.tableName) {
      let url = `/rows?table=${this.props.tableName}`
      return request.get(url).end((err, res) => {
        this.setState({ rows: res.body });
      });
    } else if (prevProps.query != this.props.query) {
      let url = '/query';
      let postBody = { query: this.props.query };
      return request.post(url).send(postBody).end((err, res) => {
        this.setState({ rows: res.body });
      });
    }

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
        <QueryBox />
        <div className="row">
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
        </div>
      </section>
    )
  }
}

function mapStateToProps(state) {
  return {
    tableName: state.tableState.tableName,
    query: state.queryState.query
  }
}

export default connect(mapStateToProps)(Rows);
