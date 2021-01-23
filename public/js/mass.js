"use strict";

// energyGraph Fetches records summary
function drawMassGraphs(mDivID, lmDivID, bmiDivID, fatPercentDivID, errElementID) {
    fetchHK("/v1/bodyMass")
        .then(d => readyMassCharts(mDivID, d))
        .then(d => readyLeanMassCharts(lmDivID, d))
        .then(d => readyBMICharts(bmiDivID, d))
        .then(d => readyFatPercentCharts(fatPercentDivID, d))
        .catch(e => displayError(errElementID, e));
}

function readyMassCharts (divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.mass));
    return d;
}

function readyLeanMassCharts (divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.lean_body_mass));
    return d;
}


function readyBMICharts (divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.bmi));
    return d;
}

function readyFatPercentCharts (divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.fat_percent));
    return d;
}


// drawMassDataGraph Plots the  actual graph onto a DIV
function drawMassDataGraph(divID, mass) {
    if (mass && Array.isArray(mass) && mass.length) {
        console.log(`Mass data: ${mass.length}`);
        const d = mass.map(d => [new Date(d.creation_timestamp_sec * 1000), d.value]);
        const dataHeader = ["Date", "Mass"];
        const data = google.visualization.arrayToDataTable([dataHeader, ...d]);
        // https://developers.google.com/chart/interactive/docs/gallery/areachart#Configuration_Options
        const options = {
            title: 'Body Mass', titlePosition: "none",
            titleTextStyle: {fontSize: 32},
            curveType: 'function', legend: {position: 'in'}
        };
        const chart = new google.visualization.LineChart(document.getElementById(divID));
        chart.draw(data, options);
    }
}
