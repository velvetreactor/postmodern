import request from 'superagent';

class API {
  static async Authenticate(connectionString) {
    let authRes = await request.post('/sessions').send({ connectionString });
    if (authRes.statusCode === 200) {
      $('#credentials-modal').modal('hide');
      return authRes;
    } else {
      throw new Error(authRes.body);
    }
  }

  static async FetchTables() {
    let tablesRes = await request.get('/tables');
    if (tablesRes.statusCode === 200) {
      return tablesRes;
    } else {
      throw new Error(tablesRes.body);
    }
  }
}

export default API
