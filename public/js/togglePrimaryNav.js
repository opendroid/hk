"use strict";

/**
 * togglePrimaryNav: Method hides or shows the HTML navigation section id "nav-primary"
 * Uses classes: nav.hide, hk-top-bar-secondary-no-margin and hk-main-no-margin
 */
function togglePrimaryNav() {
  const classes = document.getElementById("nav-primary").classList;
  if (classes && classes.contains("hide")) { // Show navbar
    document.getElementById("nav-primary").setAttribute("class", "np");
    document.getElementById("hk-main").setAttribute("class", "hk-main");
    document.getElementById("nav-secondary").setAttribute("class", "ns secondary-light");
  } else { // Hide navbar
    document.getElementById("nav-primary").setAttribute("class", "np hide");
    document.getElementById("hk-main").setAttribute("class", "hk-main no-margin");
    document.getElementById("nav-secondary").setAttribute("class", "ns secondary-light no-margin");
  }
}
