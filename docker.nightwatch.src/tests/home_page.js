module.exports = {
  'It shows a textarea': browser => {
    browser
      .url(browser.launch_url)
      .waitForElementVisible('body', 1000, false, () => {
        browser.expect.element('textarea').to.be.visible;
        browser.end();
      });
  },
  'It executes a SQL query': browser => {
    browser
      .url(browser.launch_url)
      .waitForElementVisible('body', 1000);
    let sqlString = 'SELECT * FROM items LIMIT 10;';
    browser.setValue('textarea', sqlString);
    browser.click('button');

    browser.expect.element('tr').to.be.visible;
  }
}
