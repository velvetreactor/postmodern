import React from 'react';
import { connect } from 'react-redux';

const SqlErrorAlert = props => {
  if (props.sqlErrorState.visible) {
    return (
      <p className="col-lg-10 text-danger sql-error-message">{props.sqlErrorState.message}</p>
    )
  }
  return null;
}

function mapStateToProps(state) {
  return {
    sqlErrorState: state.sqlErrorState
  };
}

function mapDispatchToProps(dispatch) {
  return {};
}

export default connect(mapStateToProps, mapDispatchToProps)(SqlErrorAlert);
