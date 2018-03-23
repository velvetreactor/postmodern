module.exports = {
  'A User should continue to see the login modal if they click outside of it': browser => {
    browser
      .url(browser.launch_url)
      .waitForElementVisible('#credentials-modal', 1000);
    browser
      .moveTo('.modal-dialog', 0, 1000)
      .mouseButtonClick()
      .pause(500);
    browser.expect.element('#credentials-modal').to.be.visible;
    browser.end();
  }
}
