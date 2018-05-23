import { createStore, applyMiddleware } from 'redux';

import masterReducer from '../../src/reducers';
import * as constants from '../../src/constants'

describe('when handling the TABLES_FETCH_SUCCEEDED event', () => {
  let store = createStore(masterReducer)

  test('it sets the tables array', () => {
    let tables = ['foobar']
    store.dispatch({ type: constants.TABLES_FETCH_SUCCEEDED, tables: tables });

    let state = store.getState().databaseState;
    expect(state.tables).toEqual(tables);
  });
});
