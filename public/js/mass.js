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

function readyMassCharts(divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.mass, "Mass (lbs)"));
    return d;
}

function readyLeanMassCharts(divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.lean_body_mass, "Lean Body Mass (lbs)"));
    return d;
}

function readyBMICharts(divID, d) {
    google.charts.setOnLoadCallback(() => drawMassDataGraph(divID, d.bmi, "BMI"));
    return d;
}

function readyFatPercentCharts(divID, d) {
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
        drawChart(divID, [dataHeader, ...d], msg)
    }
}


/**
 * drawMassDashboards
 * @param divIDs
 */
function drawMassDashboards(divIDs) {
    fetchHK("/v1/bodyMass")
        .then(d => {
            readyMassDashboard(divIDs[0], prepareMassData(d.mass, "Mass"));
            return d;
        })
        .catch(e => displayError(divIDs[0].err, e))
        .then(d => {
            readyMassDashboard(divIDs[1], prepareMassData(d.lean_body_mass, "LBM"));
            return d;
        })
        .catch(e => displayError(divIDs[1].err, e))
        .then(d => {
            readyMassDashboard(divIDs[2], prepareMassData(d.bmi, "LBM"));
            return d;
        })
        .catch(e => displayError(divIDs[2].err, e))
        .then(d => {
            readyMassDashboard(divIDs[3], prepareMassData(d.fat_percent, "LBM"));
            return d;
        })
        .catch(e => displayError(divIDs[3].err, e))
}

// readyMassDashboard
function readyMassDashboard(divs, data) {
    google.charts.setOnLoadCallback(() => drawMassDashboardChart(divs, data));
    console.log(`readyMassDashboard text: ${data.length}, divs:${JSON.stringify(divs)}`);
}

function drawMassDashboardChart(divs, data) {
    let chart = prepareSummaryMainChart(divs.main);
    let control = prepareSummaryControl(divs.control);
    let dashboard = new google.visualization.Dashboard(document.getElementById(divs.dashboard));
    dashboard.bind(control, chart);
    dashboard.draw(data);
}

function prepareMassData(data, txt) {
    try {
        if (data && Array.isArray(data) && data.length) {
            console.log(`prepareMassData: ${data.length}`);
            const d = data.map(d => {
                if (d.unit === "%") {
                    return [new Date(d.creation_timestamp_sec * 1000),
                        {v: d.value * 100, f: `${numberWithComma(d.value * 100)} ${d.unit}, ${d.source}`}];
                } else {
                    return [new Date(d.creation_timestamp_sec * 1000),
                        {v: d.value, f: `${d.value} ${d.unit}, ${d.source}`}];
                }
            });
            const dataHeader = [{label: "Date", type: "datetime"}, {label: txt, type: "number"}];
            return [dataHeader, ...d]
        }
    } catch (e) {
        console.error(`prepareMassData; ${e}`);
    }
}

function prepareSummaryMainChart(mDivID) {
    // Option fields here https://developers.google.com/chart/interactive/docs/gallery/areachart
    let options = {
        chartType: 'LineChart',
        containerId: mDivID,
        options: {
            titleTextStyle: {fontSize: 32},
            curveType: 'function',
            legend: {position: 'in'},
            view: {columns: [0, 1]}, // Pick data from column 0 and 1
        }
    };
    return new google.visualization.ChartWrapper(options);
}

function prepareSummaryControl(cDivID) {
    return new google.visualization.ControlWrapper({
        controlType: 'ChartRangeFilter',
        containerId: cDivID,
        options: {
            filterColumnIndex: 0,
            ui: {
                chartType: 'LineChart',
                chartView: {
                    columns: [0, 1]
                },
                minRangeSize: 864000000
            }
        },
    });
}