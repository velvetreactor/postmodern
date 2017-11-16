import { combineReducers } from 'redux';

function tableState(state, action) {
  switch (action.type) {
    case 'CHANGE_TABLE':
      return Object.assign({}, state, { tableName: action.tableName });
    default:
      return { tableName: '' };
  }
}

function queryState(state, action) {
  switch (action.type) {
    case 'SET_QUERY':
      return Object.assign({}, state, { query: action.query });
    default:
      return { query: '' };
  }
}

export default combineReducers({ tableState, queryState });
