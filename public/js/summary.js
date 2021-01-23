"use strict";

// summaryGTable: Fetches records summary
function summaryGTable(tableID, errElementID) {
    fetchHK("/v1/activitySummary")
        .then(d => readyGTable(tableID, d))
        .catch(e => displayError(errElementID, e));
}

// Draws energy graphs
function readyGTable(tableID, d) {
    try {
        console.log(`Try table ${tableID}. Len: ${d.length}`)
        google.charts.load('current', {'packages': ['table']});
        google.charts.setOnLoadCallback(() => drawSummaryTable(tableID, d));
        return d;
    } catch (err) {
        console.error(`readyGTable: ${err}`);
    }
}

// drawSummaryTable Plots the  actual graph onto a DIV
function drawSummaryTable(tableID, summary) {
    try {
        if (summary && Array.isArray(summary) && summary.length) {
            const data = new google.visualization.DataTable();
            // Add table headers
            data.addColumn('string', 'Date');
            data.addColumn('number', 'Energy Burned');
            data.addColumn('number', 'Goal');
            data.addColumn('number', 'Exercise');
            data.addColumn('number', 'Goal');
            data.addColumn('number', 'Stand Hours');
            data.addColumn('number', 'Goal');
            // Add table data
            const items = summary.map((d) => {
                return [d.yyyymmdd,
                    {v: parseFloat(d.energy_burned), f: `Cal ${numberWithComma(d.energy_burned)}`},
                    {v: parseFloat(d.energy_burned_goal), f: `${numberWithComma(d.energy_burned_goal)}`},
                    {v: parseFloat(d.exercise_time), f: `Min ${numberWithComma(d.exercise_time)}`},
                    {v: parseFloat(d.exercise_time_goal), f: `${numberWithComma(d.exercise_time_goal)}`},
                    {v: parseFloat(d.stand_hours), f: `Hr ${numberWithComma(d.stand_hours)}`},
                    {v: parseFloat(d.stand_hours_goal), f: `${numberWithComma(d.stand_hours_goal)}`}];
            });
            data.addRows([...items]);
            const options = {showRowNumber: true};
            const table = new google.visualization.Table(document.getElementById(tableID));
            const date_formatter = new google.visualization.DateFormat({formatType: 'medium', timeZone: -8});
            date_formatter.format(data,2);
            table.draw(data, options);
        }
    } catch (err) {
        console.error(`Table error: ${err}`);
    }
}

// energyGraph Fetches records summary
function drawGraphs(eDivID, exDivID, sDivID, errElementID) {
    fetchHK("/v1/activitySummary")
        .then(d => readyEnenrgyCharts(eDivID, d))
        .then(d => readyExerciseCharts(exDivID, d))
        .then(d => readyStandCharts(sDivID, d))
        .catch(e => displayError(errElementID, e));
}

// Draws energy graphs
function readyEnenrgyCharts(divID, d) {
    google.charts.setOnLoadCallback(() => drawEnergyGraph(divID, d));
    return d;
}

// Draws exercise graphs
function readyExerciseCharts(divID, d) {
    google.charts.setOnLoadCallback(() => drawExerciseGraph(divID, d));
    return d;
}

// Draws stand graphs
function readyStandCharts(divID, d) {
    google.charts.setOnLoadCallback(() => drawStandHoursGraph(divID, d));
    return d;
}

// drawEnergyGraph Plots the  actual graph onto a DIV
function drawEnergyGraph(divID, summary) {
    if (summary && Array.isArray(summary) && summary.length) {
        const d = summary.map(d => [new Date(d.yyyymmdd), parseFloat(d.energy_burned), parseFloat(d.energy_burned_goal)]);
        const dataHeader = ["Date", "Energy Burned", "Goal"];
        const data = google.visualization.arrayToDataTable([dataHeader, ...d]);
        // https://developers.google.com/chart/interactive/docs/gallery/areachart#Configuration_Options
        const options = {
            title: 'Energy Burned', titlePosition: "none",
            titleTextStyle: {fontSize: 32},
            curveType: 'function', legend: {position: 'in'}
        };
        const chart = new google.visualization.LineChart(document.getElementById(divID));
        chart.draw(data, options);
    }
}

// drawExerciseGraph Plots the  actual graph onto a DIV
function drawExerciseGraph(divID, summary) {
    if (summary && Array.isArray(summary) && summary.length) {
        const d = summary.map(d => [new Date(d.yyyymmdd), parseFloat(d.exercise_time), parseFloat(d.exercise_time_goal)]);
        const dataHeader = ["Date", "Exercise", "Goal"];
        const data = google.visualization.arrayToDataTable([dataHeader, ...d]);
        // https://developers.google.com/chart/interactive/docs/gallery/areachart#Configuration_Options
        const options = {
            title: 'Energy Burned', titlePosition: "none",
            titleTextStyle: {fontSize: 32},
            curveType: 'function', legend: {position: 'in'}
        };
        const chart = new google.visualization.LineChart(document.getElementById(divID));
        chart.draw(data, options);
    }
}

// drawStandHoursGraph Plots the  actual graph onto a DIV
function drawStandHoursGraph(divID, summary) {
    if (summary && Array.isArray(summary) && summary.length) {
        const d = summary.map(d => [new Date(d.yyyymmdd), parseFloat(d.stand_hours), parseFloat(d.stand_hours_goal)]);
        const dataHeader = ["Date", "Stand", "Goal"];
        const data = google.visualization.arrayToDataTable([dataHeader, ...d]);
        // https://developers.google.com/chart/interactive/docs/gallery/areachart#Configuration_Options
        const options = {
            title: 'Energy Burned', titlePosition: "none",
            titleTextStyle: {fontSize: 32},
            curveType: 'function', legend: {position: 'in'}
        };
        const chart = new google.visualization.LineChart(document.getElementById(divID));
        chart.draw(data, options);
    }
}