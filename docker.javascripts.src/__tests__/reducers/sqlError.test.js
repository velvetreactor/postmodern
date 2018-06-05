import { createStore, applyMiddleware } from 'redux';

import masterReducer from '../../src/reducers';
import * as constants from '../../src/constants';

describe('when handling the QUERY_EXECUTION_FAILED event', () => {
  let store = createStore(masterReducer)

  test('it sets the visibility and message', () => {
    let error = 'My error msg'
    store.dispatch({ type: constants.QUERY_EXECUTION_FAILED, error: error });

    let state = store.getState().sqlErrorState;
    expect(state.message).toEqual(error);
    expect(state.visible).toEqual(true);
  });
});
