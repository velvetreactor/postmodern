import React from 'react'
import { connect } from 'react-redux';

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
        <button className="btn btn-primary execute-query-btn" onClick={() => {
          let sqlQuery = this.state.query;
          this.props.changeQuery(sqlQuery);
        }}>Query</button>
      </div>
    )
  }
}

function setQueryAction(query) {
  return {
    type: 'SET_QUERY',
    query: query
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
      dispatch(setQueryAction(query))
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(QueryBox);
