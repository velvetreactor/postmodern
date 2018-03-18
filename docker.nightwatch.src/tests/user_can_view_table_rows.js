const TestHelper = require('../test_helper.js');

module.exports = {
  'A User can click on table and view it\'s rows': browser => {
    TestHelper.login(browser);
    browser.refresh();
    browser
      .url(browser.launch_url)
      .waitForElementVisible('.tables li', 1000)
      .click('li.tables-items a');
    browser
      .waitForElementVisible('.table-rows tbody tr', 1000);
    browser.end();
  }
}
