const TestHelper = require('../support/test_helper.js');

module.exports = {
  'A User can execute SQL queries': browser => {
    TestHelper.login(browser);
    let sqlQry = "SELECT * FROM items WHERE name = 'Pencil'"
    browser
      .setValue('.query-box textarea', sqlQry)
      .click('.execute-query-btn')
      .waitForElementVisible('.tables li', 1000);
    browser.end();
  }
}
