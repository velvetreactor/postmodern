import { createStore, applyMiddleware } from 'redux';

import masterReducer from '../../src/reducers';
import * as constants from '../../src/constants'

describe('when handling the CHANGE_TABLE event', () => {
  let store = createStore(masterReducer)

  test('it sets the tables array', () => {
    let expectedTable = 'foobar';
    store.dispatch({ type: constants.CHANGE_TABLE, tableName: expectedTable });

    let state = store.getState().tableState;
    expect(state.tableName).toEqual(expectedTable);
  });
});
