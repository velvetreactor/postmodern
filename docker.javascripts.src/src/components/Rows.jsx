import React from 'react';
import { connect } from 'react-redux';

import QueryBox from './QueryBox.jsx';

class Rows extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      rows: []
    }
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevProps == this.props) {
      return false;
    }

    if (prevProps.tableName != this.props.tableName) {
      let url = `/rows?table=${this.props.tableName}`
      return request.get(url).end((err, res) => {
        this.props.setRows(res.body);
      });
    } else if (prevProps.query != this.props.query) {
      let url = '/query';
      let postBody = { query: this.props.query };
      return request.post(url).send(postBody).end((err, res) => {
        this.props.setRows(res.body);
      });
    }
  }

  render() {
    let items;
    if (this.props.rows.length != 0) {
      items = this.props.rows.map((row, idx) => {
        return (
          <tr key={`row-${idx}`}>
            {Object.keys(this.props.rows[0]).map(name => <td key={`attr-${name}`}>{row[name]}</td>)}
          </tr>
        )
      });
    }
    let headers;
    if (this.props.rows.length != 0) {
      headers = Object.keys(this.props.rows[0]).map(name => <th key={`col-${name}`}>{name}</th>)
    }

    return (
      <section className='table-rows col-lg-10'>
        <QueryBox />
        <div className="row table-row">
          <table className="table table-striped table-bordered table-rows">
            <thead>
              <tr>
                {headers}
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

function setRowsAction(rows) {
  return {
    type: 'SET_ROWS',
    rows: rows
  }
}

function mapStateToProps(state) {
  return {
    tableName: state.tableState.tableName,
    query: state.queryState.query,
    rows: state.rowState.rows
  }
}

function mapDispatchToProps(dispatch) {
  return {
    setRows: rows => {
      dispatch(setRowsAction(rows));
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Rows);
