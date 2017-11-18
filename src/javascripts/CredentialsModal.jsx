import React from 'react';

class CredentialsModal extends React.Component {
  constructor(props) {
    super(props);
  }

  postSession() {
    console.log('Posted Session Securely');
  }

  render() {
    return (
      <div id="credentials-modal" className="modal fade" role="dialog" aria-hidden="true">
        <div className="modal-dialog modal-sm">
          <div className="modal-content">
            <div className="modal-header">
              <h4 className="modal-title">New Connection</h4>
            </div>
            <div className="modal-body">
              <form>
                <div className="form-group">
                  <label className="form-control-label">Postgres URL</label>
                  <input type="text" className="form-control" />
                </div>
                <div className="form-group">
                  <label className="form-control-label">Username</label>
                  <input type="text" className="form-control" />
                </div>
                <div className="form-group">
                  <label className="form-control-label">Password</label>
                  <input type="password" className="form-control" />
                </div>
              </form>
            </div>
            <div className="modal-footer">
              <button onClick={this.postSession} type="button" className="btn btn-success">Connect</button>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default CredentialsModal;
