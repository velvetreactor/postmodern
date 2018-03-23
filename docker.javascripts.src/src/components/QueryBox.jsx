import React from 'react'
import { connect } from 'react-redux';

import SqlErrorAlert from './SqlErrorAlert.jsx';

class QueryBox extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      query: ''
    };
  }

  render() {
    return (
      <div className="row form-group query-box">
        <textarea className="form-control" onChange={(evt) => {
          this.setState({ query: evt.target.value });
        }}></textarea>
        <div className="row query-feedback">
          <SqlErrorAlert />
          <button className="btn btn-default execute-query-btn" onClick={() => {
            this.props.changeQuery(this.state.query);
          }}>Query</button>
        </div>
      </div>
    )
  }
}

function mapStateToProps(state) {
  return {
    query: state.queryState.query,
    tableName: state.tableState.tableName
  };
}

function mapDispatchToProps(dispatch) {
  return {
    changeQuery: query => {
      dispatch({ type: 'QUERY_EXECUTION_REQUESTED', query })
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(QueryBox);
