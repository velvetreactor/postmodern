import { takeLatest, call, put } from 'redux-saga/effects';
import API from './api.js';

function* connectToDatabase(action) {
  try {
    const authRes = yield call(API.Authenticate, action.payload.connectionString)
    yield put({ type: 'DB_CONNECTION_SUCCEEDED' });
    yield put({ type: 'TABLES_FETCH_REQUESTED' });
  } catch(err) {
    yield put({ type: 'DB_CONNECTION_FAILED', error: err })
  }
}

function* fetchTables() {
  try {
    const res = yield call(API.FetchTables);
    yield put({ type: 'TABLES_FETCH_SUCCEEDED', tables: res.body.tables })
  } catch(err) {
    console.log(err);
    yield put({ type: 'TABLES_FETCH_FAILED', error: err })
  }
}

function* fetchTableRows(action) {
  try {
    const res = yield call(API.FetchTableRows, action.tableName)
    yield put({ type: 'TABLE_ROWS_FETCH_SUCCEEDED', tableRows: res.body.rows })
  } catch(err) {
    console.log(err)
    yield put({ type: 'TABLE_ROWS_FETCH_FAILED' })
  }
}

function* checkSession() {
  try {
    const sessionRes = yield call(API.CheckSession);
    $('#credentials-modal').modal('hide');
    yield put({ type: 'SESSION_CHECK_SUCCEEDED' })
    yield put({ type: 'TABLES_FETCH_REQUESTED' })
  } catch(err) {
    console.log(err);
    $('#credentials-modal').modal('show');
    yield put({ type: 'SESSION_CHECK_FAILED', error: err })
  }
}

function* mySaga() {
  yield takeLatest('DB_CONNECTION_REQUESTED', connectToDatabase);
  yield takeLatest('TABLES_FETCH_REQUESTED', fetchTables);
  yield takeLatest('SESSION_CHECK_REQUESTED', checkSession);
  yield takeLatest('TABLE_ROWS_FETCH_REQUESTED', fetchTableRows);
}

export default mySaga;
