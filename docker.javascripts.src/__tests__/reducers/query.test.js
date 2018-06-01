import { createStore, applyMiddleware } from 'redux';

import masterReducer from '../../src/reducers';
import * as constants from '../../src/constants'

describe('when handling the SET_QUERY event', () => {
  let store = createStore(masterReducer)

  test('it sets the query attribute', () => {
    let query = 'select * from rows'
    store.dispatch({ type: constants.SET_QUERY, query: query });

    let state = store.getState().queryState;
    expect(state.query).toEqual(query);
  });
});
