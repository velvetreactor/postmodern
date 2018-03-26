import { combineReducers } from 'redux';

function databaseState(state = {}, action) {
  switch (action.type) {
    case 'TABLES_FETCH_SUCCEEDED':
      return Object.assign({}, state, { tables: action.tables });
    default:
      return state.tables ? { tables: state.tables } : { tables: [] };
  }
}

function tableState(state = {}, action) {
  switch (action.type) {
    case 'CHANGE_TABLE':
      return Object.assign({}, state, { tableName: action.tableName });
    default: // happens on ANY state change
      return state.tableName ? { tableName: state.tableName } : { tableName: '' };
  }
}

function queryState(state = {}, action) {
  switch (action.type) {
    case 'SET_QUERY':
      return Object.assign({}, state, { query: action.query });
    default:
      return state.query ? { query: state.query } : { query: '' };
  }
}

function rowState(state = {}, action) {
  switch (action.type) {
    case 'QUERY_EXECUTION_SUCCEEDED':
      return Object.assign({}, state, { rows: action.tableRows });
    case 'TABLE_ROWS_FETCH_SUCCEEDED':
      return Object.assign({}, state, { rows: action.tableRows });
    default:
      return state.rows ? { rows: state.rows } : { rows: [] };
  }
}

function credentialsErrorState(state = {}, action) {
  switch (action.type) {
    case 'DB_CONNECTION_FAILED':
      let errorMsg = action.error.response.body.error
      return Object.assign({}, state, { visible: true, message: errorMsg });
    default:
      if (state.credentialsError) {
        return state.credentialsError;
      } else {
        return { visible: false, message: '' };
      }
  }
}

function sqlErrorState(state = {}, action) {
  switch (action.type) {
    case 'QUERY_EXECUTION_FAILED':
      let errorMsg = action.error;
      return Object.assign({}, state, { visible: true, message: errorMsg });
    default:
      if (state.sqlError) {
        return state.sqlError;
      } else {
        return { visible: false, message: '' };
      }
  }
}

export default combineReducers({
  sqlErrorState,
  credentialsErrorState,
  databaseState,
  tableState,
  queryState,
  rowState
});
