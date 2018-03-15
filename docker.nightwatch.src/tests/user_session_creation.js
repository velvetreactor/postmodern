module.exports = {
  'A User can create a PG session via the modal': browser => {
    let pgConnStr = 'postgres://postgres@postgres:5432/postgres?sslmode=disable'
    browser
      .url(browser.launch_url)
      .waitForElementVisible('#credentials-modal', 1000)
      .setValue('input[name="connectionString"]', pgConnStr)
      .click('.btn-success')
      .waitForElementNotVisible('#credentials-modal', 1000);

    browser.waitForElementVisible('.tables li', 1000);
  },
  'A User should not see the modal again after authenticating': browser => {
    let pgConnStr = 'postgres://postgres@postgres:5432/postgres?sslmode=disable'
    browser
      .url(browser.launch_url)
      .waitForElementVisible('#credentials-modal', 1000)
      .setValue('input[name="connectionString"]', pgConnStr)
      .click('.btn-success');
    browser
      .url(browser.launch_url)
      .expect.element('#credentials-modal')
      .to.not.be.visible;
  }
}
