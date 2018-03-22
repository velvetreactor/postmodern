const TestHelper = require('../test_helper.js');

module.exports = {
  'A User should not see the modal again after authenticating': browser => {
    TestHelper.login(browser)
    browser
      .refresh()
      .expect.element('#credentials-modal')
      .to.not.be.visible;
    browser.end();
  },
  'A User can view the database tables after authenticating': browser => {
    TestHelper.login(browser)
    browser
      .waitForElementVisible('.tables li', 1000);
    browser.end();
  },
  'A User sees a danger alert box after inputting invalid credentials': browser => {
    let invalidCreds = 'qewery';
    browser
      .url(browser.launch_url)
      .waitForElementVisible('#credentials-modal', 1000)
      .setValue('input[name="connectionString"]', invalidCreds)
      .click('.btn-success')
      .waitForElementVisible('.alert-danger', 1000);
    browser.end();
  },
  'A User should continue to see the login modal if they click outside of it': browser => {
    browser
      .url(browser.launch_url)
      .waitForElementVisible('#credentials-modal', 1000);
    browser
      .moveTo('.modal-dialog', 0, 1000)
      .mouseButtonClick();
    browser.expect.element('#credentials-modal').to.be.visible;
    browser.end();
  }
}
