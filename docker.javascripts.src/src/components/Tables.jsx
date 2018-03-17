import React from 'react';
import { connect } from 'react-redux';

class Tables extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    let tableItems = this.props.tables.map(name => {
      let uniqKey = `tables-${name}`
      return <li key={uniqKey} className={uniqKey}><a href="#" onClick={() => {
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
      dispatch({ type: 'TABLE_ROWS_FETCH_REQUESTED', tableName: tableName });
    }
  }
}

function mapStateToProps(state) {
  return {
    tables: state.databaseState.tables,
    tableName: state.tableState.tableName
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Tables);
