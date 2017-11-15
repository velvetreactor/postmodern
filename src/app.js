import React from 'react';
import ReactDOM from 'react-dom';
import App from './javascripts/App.jsx';
import { Provider } from 'react-redux';
import { createStore, combineReducers } from 'redux'
import { tableState } from './javascripts/Tables.jsx'

const masterReducer = combineReducers({ tableState });
const store = createStore(masterReducer);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app-entry')
);
