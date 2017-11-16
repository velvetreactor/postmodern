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
      <div className="row">
        <textarea name="" id="" cols="30" rows="10" onChange={(evt) => {
          this.setState({ query: evt.target.value });
        }}></textarea>
        <button onClick={() => {
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
