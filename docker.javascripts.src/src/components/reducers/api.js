import request from 'superagent';

class API {
  static async Authenticate(connectionString) {
    let res = await request.post('/sessions').send({ connectionString });
    if (res.statusCode === 200) {
      $('#credentials-modal').modal('hide');
      return res;
    } else {
      throw new Error(res.body);
    }
  }

  static async FetchTables() {
    let res = await request.get('/tables');
    if (res.statusCode === 200) {
      return res;
    } else {
      throw new Error(res.body);
    }
  }

  static async CheckSession() {
    let res = await request.get('/sessions')
    if (res.statusCode === 200) {
      return res
    } else {
      throw new Error(res.body);
    }
  }

  static async FetchTableRows(tableName) {
    let res = await request.get(`/tables/${tableName}`);
    if (res.statusCode === 200) {
      return res;
    } else {
      throw new Error(res.body);
    }
  }

  static async ExecuteQuery(query) {
    let res = await request.post('/queries').send({ query });
    if (res.statusCode === 200) {
      return res;
    } else {
      throw new Error(res);
    }
  }
}

export default API
