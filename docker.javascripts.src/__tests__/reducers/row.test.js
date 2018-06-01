import { createStore, applyMiddleware } from 'redux';

import masterReducer from '../../src/reducers';
import * as constants from '../../src/constants'

describe('when handling the QUERY_EXECUTION_SUCCEEDED event', () => {
  let store;
  beforeEach(() => {
    store = createStore(masterReducer);
  });

  test('it sets the rows array', () => {
    let rows = ['foo', 'bar'];
    store.dispatch({ type: constants.QUERY_EXECUTION_SUCCEEDED, tableRows: rows });

    let state = store.getState().rowState
    expect(state.rows).toEqual(rows);
  });
});

describe('when handling the TABLE_ROWS_FETCH_SUCCEEDED event', () => {
  let store;
  beforeEach(() => {
    store = createStore(masterReducer);
  });

  test('it sets the rows array', () => {
    let rows = ['foo', 'bar'];
    store.dispatch({ type: constants.TABLE_ROWS_FETCH_SUCCEEDED, tableRows: rows });

    let state = store.getState().rowState
    expect(state.rows).toEqual(rows);
  });
});
