const TestHelper = require('../test_helper.js');

module.exports = {
  'A User can see tables after navigating to the page for a 2nd time': browser => {
    TestHelper.login(browser);
    browser.refresh();
    browser
      .url(browser.launch_url)
      .waitForElementVisible('.tables li', 1000);
    browser.end();
  }
}
