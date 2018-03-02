import { combineReducers } from 'redux';

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
    case 'SET_ROWS':
      return Object.assign({}, state, { rows: action.rows });
    default:
      return state.rows ? { rows: state.rows } : { rows: [] };
  }
}

export default combineReducers({ tableState, queryState, rowState });
