import React from 'react';
import { connect } from 'react-redux';

import QueryBox from './QueryBox.jsx';

class Rows extends React.Component {
  constructor(props) {
    super(props);
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

function mapStateToProps(state) {
  return {
    rows: state.rowState.rows
  }
}

function mapDispatchToProps(dispatch) {
  return {
    setRows: rows => {
      dispatch({ type: 'SET_ROWS', rows: rows });
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Rows);
