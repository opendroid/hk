"use strict";

// summaryGTable: Fetches records summary
function summaryWalkSteps(wDivID, dDivID, errElementID) {
    fetchHK("/v1/walks")
        .then(d => readyWalkStepCountChart(wDivID, d))
        .then(d => readyWalkDistanceChart(dDivID, d))
        .catch(e => displayError(errElementID, e));
}

// Draws energy graphs
function readyWalkStepCountChart(divID, d) {
    google.charts.setOnLoadCallback(() => drawWalkingDistanceStepCountGraph(divID, d.step_count, "Step Count"));
    return d;
}

// Draws energy graphs
function readyWalkDistanceChart(divID, d) {
    google.charts.setOnLoadCallback(() => drawWalkingDistanceStepCountGraph(divID, d.distance_walking_running, "Walking Distance (Miles)"));
    return d;
}

// drawWalkingDistanceStepCountGraph Plots the  actual graph onto a DIV
function drawWalkingDistanceStepCountGraph(divID, stepCount, msg) {
    if (stepCount && Array.isArray(stepCount) && stepCount.length) {
        let value = 0;
        const d = stepCount.map(d => {
            value += d.value;
            return [new Date(d.yyyymmdd),
                {v: d.value, f: `${numberWithComma(d.value)} ${d.unit}, ${d.source} (trips#${d.count})`}];
        });
        const labelText = `${msg}`
        const dataHeader = [{label: "Date", type: "date"}, {label: labelText, type: "number"}];
        drawChart(divID, [dataHeader, ...d], msg)
    }
}

