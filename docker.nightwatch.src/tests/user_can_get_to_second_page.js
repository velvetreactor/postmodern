const TestHelper = require('../support/test_helper.js');

module.exports = {
  'A User can get to the 2nd page of the Tables show results': browser => {
    TestHelper.login(browser);
    browser.refresh();
    browser
      .url(browser.launch_url)
      .waitForElementVisible('.tables li', 1000)
      .click('li.tables-items a');
    browser
      .click('li.pages a[data-page-num=2]');
    browser
      .waitForElementVisible('.table-rows tbody tr', 1000);
    browser.end();
  }
}
