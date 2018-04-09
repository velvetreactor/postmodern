const TestHelper = require('../support/test_helper.js');

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
  },
  'A User can click to page 2 in a table rows result set': browser => {
    TestHelper.login(browser);
    browser
      .waitForElementVisible('.tables li', 1000)
      .click('li.tables-items a');
    browser
      .click('ul.pages a[data-results-page="2"]');
    browser.elements('css selector', '.table-rows tbody tr', qryObj => {
      let els = qryObj.value;
      browser.assert.equal(els.length, 8)
    });
    browser.end()
  }
}
