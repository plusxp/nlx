// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

hljs.initHighlightingOnLoad();
mermaid.initialize({
    startOnLoad:true
});

if (!window._env) {
  console.info('#540 The links in the navbar won\'t work during development because we depend on env.js which is generated as part of starting the Nginx container.')
  return
}

const navbarLogoLink = document.getElementById('navbar-logo-link');
navbarLogoLink.href = window._env.NAVBAR_HOME_PAGE_URL;

const headerAboutLink = document.getElementById('navbar-about-link');
headerAboutLink.href = window._env.NAVBAR_ABOUT_PAGE_URL;

const headerDirectoryLink = document.getElementById('navbar-directory-link');
headerDirectoryLink.href = window._env.NAVBAR_DIRECTORY_URL;
