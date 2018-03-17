const TestHelper = require('../test_helper.js');

module.exports = {
  'A User should not see the modal again after authenticating': browser => {
    TestHelper.login(browser)
    browser
      .url(browser.launch_url)
      .expect.element('#credentials-modal')
      .to.not.be.visible;
    browser.end();
  },
  'A User can view the database tables after authenticating': browser => {
    TestHelper.login(browser)
    browser
      .waitForElementVisible('.tables li', 1000);
    browser.end();
  }
}
