import React from 'react';
import { connect } from 'react-redux';

const CredentialsErrorAlert = props => {
  if (props.credentialsErrorState.visible) {
    return (
      <div className="alert alert-danger credentials-alert">
        {props.credentialsErrorState.message}
      </div>
    );
  } else {
    return null;
  }
}

function mapStateToProps(state) {
  return {
    credentialsErrorState: state.credentialsErrorState
  };
}

function mapDispatchToProps(dispatch) {
  return {};
}

export default connect(mapStateToProps, mapDispatchToProps)(CredentialsErrorAlert);
