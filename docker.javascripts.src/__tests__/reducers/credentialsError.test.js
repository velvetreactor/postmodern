import { createStore, applyMiddleware } from 'redux';

import masterReducer from '../../src/reducers';
import * as constants from '../../src/constants'

describe('when handling the DB_CONNECTION_FAILED event', () => {
  let store = createStore(masterReducer)

  test('it sets the visibility and message', () => {
    let error = { response: { body: { error: 'My error msg' } } };
    store.dispatch({ type: constants.DB_CONNECTION_FAILED, error: error });

    let state = store.getState().credentialsErrorState;
    let errorMsg = error.response.body.error
    expect(state.message).toEqual(errorMsg);
    expect(state.visible).toEqual(true);
  });
});
