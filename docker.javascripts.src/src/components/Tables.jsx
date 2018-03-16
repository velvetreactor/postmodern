import React from 'react';
import request from 'superagent';
import { connect } from 'react-redux';

class Tables extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    let tableItems = this.props.tables.map((name,idx) => {
      return <li key={`tables-${name}`}><a href="#" onClick={() => {
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
    tables: state.databaseState.tables,
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
