module.exports = {
  'A User can create a PG session via the modal': browser => {
    let pgConnStr = 'postgres://postgres@postgres:5432/postgres?sslmode=disable'
    browser
      .url(browser.launch_url)
      .waitForElementVisible('#credentials-modal', 1000)
      .setValue('input[name="postgres-conn-string"]', pgConnStr);
    browser
      .click('.btn-success');

    browser.expect.element('.tables li').to.be.present;
  }
}
