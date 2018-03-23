const pgConnStr = 'postgres://postgres@postgres:5432/postgres?sslmode=disable';
module.exports = {
  pgConnStr: pgConnStr,
  login: browser => {
    browser.maximizeWindow();
    browser
      .url(browser.launch_url)
      .waitForElementVisible('#credentials-modal', 1000)
      .setValue('input[name="connectionString"]', pgConnStr)
      .click('.btn-success')
      .waitForElementNotVisible('#credentials-modal', 1000);
  }
}
