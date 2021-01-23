"use strict";

// drawExposureGraphs Fetches records summary
function drawExposureGraphs(eDivID, errElementID) {
    fetchHK("/v1/exposure")
        .then(d => readyExposureChart(eDivID, d.exposure))
        .catch(e => displayError(errElementID, e));
}

// readyExposureChart
function readyExposureChart(divID, d) {
    google.charts.setOnLoadCallback(() => drawExposureGraph(divID, d));
    return d;
}

// drawMassDataGraph Plots the  actual graph onto a DIV
function drawExposureGraph(divID, exposure) {
    if (exposure && Array.isArray(exposure) && exposure.length) {
        console.log(`Exposure data: ${exposure.length}`);
        const d = exposure.map(d => [new Date(d.creation_timestamp_sec * 1000), d.value]);
        const dataHeader = ["Date", "Exposure"];
        const data = google.visualization.arrayToDataTable([dataHeader, ...d]);
        // https://developers.google.com/chart/interactive/docs/gallery/areachart#Configuration_Options
        const options = {
            title: 'Audio Exposure', titlePosition: "none",
            titleTextStyle: {fontSize: 32},
            curveType: 'function', legend: {position: 'in'}
        };
        const chart = new google.visualization.LineChart(document.getElementById(divID));
        chart.draw(data, options);
    }
}
