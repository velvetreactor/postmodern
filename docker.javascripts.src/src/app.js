import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import 'babel-polyfill';
import createSagaMiddleware from 'redux-saga';

import App from './components/App.jsx';
import masterReducer from './reducers';
import mySaga from './reducers/sagas.js';

const sagaMiddleware = createSagaMiddleware();
const store = createStore(masterReducer, applyMiddleware(sagaMiddleware));

sagaMiddleware.run(mySaga);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app-entry')
);
