import React from 'react';
import request from 'superagent';

class CredentialsModal extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      connectionString: ''
    }
  }

  componentWillMount() {
    request
      .get('/sessions')
      .end((err, res) => {
        if (res.statusCode != 200) {
          $('#credentials-modal').modal('show');
        }
      });
  }

  postSession() {
    request
      .post('/sessions')
      .send(this.state)
      .end((err, res) => {
        if (res.statusCode == 200) {
          $('#credentials-modal').modal('hide');
        }
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
                  <label className="form-control-label">PG Connection String</label>
                  <input
                    placeholder="postgres://postgres@localhost:5432/postgres?sslmode=verify"
                    name="connectionString"
                    type="text"
                    className="form-control"
                    value={this.state.connectionString}
                    onChange={this.handleInputChange.bind(this)}
                  />
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
