import React from 'react';
import request from 'superagent';
import { combineReducers } from 'redux';
import { connect } from 'react-redux';

class Tables extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      tableNames: []
    }
  }

  componentDidMount() {
    // request.get('/tables').end((err, res) => {
    //   this.setState({ tableNames: res.body });
    // });
  }

  render() {
    let tableItems = this.state.tableNames.map(name => {
      return <li><a href="#" onClick={() => {
        this.props.changeCurrentTable(name);
      }}>{name}</a></li>
    })

    return (
      <section className='tables col-lg-2'>
        <ul>
          {tableItems}
        </ul>
      </section>
    )
  }
}

function mapDispatchToProps(dispatch) {
  return {
    changeCurrentTable: tableName => {
      dispatch(changeTableAction(tableName));
    }
  }
}

function mapStateToProps(state) {
  return {
    tableName: state.tableState.tableName
  }
}

function changeTableAction(tableName) { // Action creator
  return {
    type: 'CHANGE_TABLE',
    tableName: tableName
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Tables);
