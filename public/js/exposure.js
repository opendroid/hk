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
        let ed = [];
        exposure.forEach(d => {
            const duration = numberWithComma((d.end_date - d.creation_timestamp_sec) / 60);
            ed.push([new Date(d.creation_timestamp_sec * 1000),
                {v: d.value, f: `${numberWithComma(d.value)} ${d.unit} (${duration} Minutes), ${d.source}`}]);
        });
        const dataHeader = [{label: "Date", type: "datetime"}, {label: "Exposure", type: "number"}];
        const data = google.visualization.arrayToDataTable([dataHeader, ...ed]);
        // https://developers.google.com/chart/interactive/docs/gallery/areachart#Configuration_Options
        const options = {
            title: 'Audio Exposure', titlePosition: "none",
            titleTextStyle: {fontSize: 32},
            curveType: 'function', legend: {position: 'in'},
            explorer: {
                axis: 'horizontal',
                keepInBounds: true,
                maxZoomIn: .01,
                actions: ['dragToZoom', 'rightClickToReset']
            },
            crosshair: { trigger: "both", orientation: "both", color: "orange"}
        };
        const chart = new google.visualization.LineChart(document.getElementById(divID));
        chart.draw(data, options);
    }
}
