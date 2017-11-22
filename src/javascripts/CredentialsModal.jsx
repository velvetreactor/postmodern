import React from 'react';
import request from 'superagent';

class CredentialsModal extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '',
      password: '',
      host: '',
      port: '',
      db_name: ''
    }
  }

  postSession() {
    request
      .post('/sessions')
      .send(this.state)
      .end((err, res) => {
      });
  }

  handleInputChange(evt) {
    let field = evt.target.name;
    this.setState({ [field]: evt.target.value });
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
                  <label className="form-control-label">Username</label>
                  <input name="username" type="text" className="form-control" value={this.state.username} onChange={this.handleInputChange.bind(this)} />
                </div>
                <div className="form-group">
                  <label className="form-control-label">Password</label>
                  <input name="password" type="password" className="form-control" value={this.state.password} onChange={this.handleInputChange.bind(this)} />
                </div>
                <div className="form-group">
                  <label className="form-control-label">Host</label>
                  <input name="host" type="text" className="form-control" value={this.state.host} onChange={this.handleInputChange.bind(this)} />
                </div>
                <div className="form-group">
                  <label className="form-control-label">Port</label>
                  <input name="port" type="text" className="form-control" value={this.state.port} onChange={this.handleInputChange.bind(this)} />
                </div>
                <div className="form-group">
                  <label className="form-control-label">Database Name</label>
                  <input name="db_name" type="text" className="form-control" value={this.state.db_name} onChange={this.handleInputChange.bind(this)} />
                </div>
              </form>
            </div>
            <div className="modal-footer">
              <button onClick={this.postSession.bind(this)} type="button" className="btn btn-success">Connect</button>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default CredentialsModal;
