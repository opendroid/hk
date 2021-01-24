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
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.mass, "Mass"));
    return d;
}

function readyLeanMassCharts (divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.lean_body_mass, "Lean Body Mass"));
    return d;
}


function readyBMICharts (divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.bmi, "BMI"));
    return d;
}

function readyFatPercentCharts (divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.fat_percent, "Fat %"));
    return d;
}

// drawMassDataGraph Plots the  actual graph onto a DIV
function drawMassDataGraph(divID, mass, msg) {
    if (mass && Array.isArray(mass) && mass.length) {
        console.log(`Mass data: ${mass.length}`);
        const d = mass.map(d => {
            if (d.unit === "%") {
                return [new Date(d.creation_timestamp_sec * 1000),
                    {v: d.value * 100, f: `${numberWithComma(d.value * 100)} ${d.unit}, ${d.source}`}];
            } else {
                return [new Date(d.creation_timestamp_sec * 1000),
                    {v: d.value, f: `${d.value} ${d.unit}, ${d.source}`}];
            }
        });
        const dataHeader = [{label: "Date", type: "datetime"}, {label: msg, type: "number"}];
        const data = google.visualization.arrayToDataTable([dataHeader, ...d]);
        // https://developers.google.com/chart/interactive/docs/gallery/areachart#Configuration_Options

        const options = {
            title: `${msg}`, titlePosition: "none",
            titleTextStyle: {fontSize: 32},
            curveType: 'function', legend: {position: 'in'},
            explorer: { axis: 'horizontal', keepInBounds: true, maxZoomIn: .01, actions: ['dragToZoom', 'rightClickToReset']}
        };
        const chart = new google.visualization.LineChart(document.getElementById(divID));
        chart.draw(data, options);
    }
}
