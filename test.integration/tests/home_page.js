module.exports = {
  'It shows a textarea': browser => {
    browser
      .url(browser.launch_url)
      .waitForElementVisible('body', 1000, false, () => {
        browser.expect.element('textarea').to.be.visible;
        browser.end();
      });
  }
}
