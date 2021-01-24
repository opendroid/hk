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
        google.charts.setOnLoadCallback(() => addSummaryTable(tableID, d));
        return d;
    } catch (err) {
        console.error(`readyGTable: ${err}`);
    }
}

// addSummaryTable Plots the  actual graph onto a DIV
function addSummaryTable(tableID, summary) {
    try {
        if (summary && Array.isArray(summary) && summary.length) {
            const data = new google.visualization.DataTable();
            // Add table headers
            data.addColumn('date', 'Date');
            data.addColumn('number', 'Energy Burned');
            data.addColumn('number', 'Goal');
            data.addColumn('number', 'Exercise');
            data.addColumn('number', 'Goal');
            data.addColumn('number', 'Stand Hours');
            data.addColumn('number', 'Goal');
            // Add table data
            const items = summary.map((d) => {
                return [new Date(d.yyyymmdd),
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
            table.draw(data, options);
        }
    } catch (err) {
        console.error(`Table error: ${err}`);
    }
}

// energyGraph Fetches records summary
function drawSummaryGraphs(eDivID, exDivID, sDivID, errElementID) {
    fetchHK("/v1/activitySummary")
        .then(d => filterBadDatesData(d))
        .then(d => readyEnergyCharts(eDivID, d))
        .then(d => readyExerciseCharts(exDivID, d))
        .then(d => readyStandCharts(sDivID, d))
        .catch(e => displayError(errElementID, e));
}

// filterBadDatesData removes dates before 2000, Apple watch has some bad data.
function filterBadDatesData(data) {
    if (data && Array.isArray(data) && data.length) {
        return data.filter(d => {
            try {
                const yyyymmdd = d.yyyymmdd.split("-");
                return parseInt(yyyymmdd[0]) > 2000;
            } catch (e) {
                return false
            }
        })
    }
    return [];
}

// Draws energy graphs
function readyEnergyCharts(divID, d) {
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
        const d = summary.map(d => {
            return [new Date(d.yyyymmdd),
                {
                    v: parseFloat(d.energy_burned),
                    f: `${numberWithComma(parseFloat(d.energy_burned))} ${d.energy_burned_unit}`
                },
                {
                    v: parseFloat(d.energy_burned_goal),
                    f: `${numberWithComma(parseFloat(d.energy_burned_goal))} ${d.energy_burned_unit}`
                }]
        });
        const dataHeader = [{label: "Date", type: "date"}, "Energy Burned (Cal)", "Goal"];
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
        const d = summary.map(d => {
            return [new Date(d.yyyymmdd),
                {v: parseFloat(d.exercise_time), f: `${parseFloat(d.exercise_time)} Minutes`},
                {v: parseFloat(d.exercise_time_goal), f: `${parseFloat(d.exercise_time_goal)} Minutes`}];
        });
        const dataHeader = [{label: "Date", type: "date"}, "Exercise (Minutes)", "Goal"];
        const data = google.visualization.arrayToDataTable([dataHeader, ...d]);
        // https://developers.google.com/chart/interactive/docs/gallery/areachart#Configuration_Options
        const options = {
            title: 'Exercise Minutes', titlePosition: "none",
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
        const d = summary.map(d => {
            return [new Date(d.yyyymmdd),
                {v: parseFloat(d.stand_hours), f: `${parseFloat(d.stand_hours)} Hours`},
                {v: parseFloat(d.stand_hours_goal), f: `${parseFloat(d.stand_hours_goal)} Hours`}];
        });
        const dataHeader = [{label: "Date", type: "date"}, "Stand (Hours)", "Goal"];
        const data = google.visualization.arrayToDataTable([dataHeader, ...d]);
        // https://developers.google.com/chart/interactive/docs/gallery/areachart#Configuration_Options
        const options = {
            title: 'Stand Hours', titlePosition: "none",
            titleTextStyle: {fontSize: 32},
            curveType: 'function', legend: {position: 'in'}
        };
        const chart = new google.visualization.LineChart(document.getElementById(divID));
        chart.draw(data, options);
    }
}