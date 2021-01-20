"use strict";


// Fetches records summary
function recordsSummary(tableID) {
    fetchHK("v1/recordsSummary")
        .then(d => insertSummaryTable(tableID, d))
        .catch(e => console.log(`${JSON.stringify(e)}`));
}

// Fetches records summary
function recordsTypes(tableID) {
    fetchHK("v1/recordsTypes")
        .then(d => insertSummaryTable(tableID, d))
        .catch(e => console.log(`${JSON.stringify(e)}`));
}

// insertSummaryTable inserts Key/Value pairs in table with ID
function insertSummaryTable(tid, d) {
    let t = document.getElementById(tid); // Table with ID
    addHeadingRow(t, ["Type", "Count"]);
    addKeyCountDataRows(t, d);
}

// addHeadingRow
function addHeadingRow(table, headings) {
    let th = table.createTHead();
    let row = th.insertRow();
    headings.forEach(key => {
        let th = document.createElement("th");
        let text = document.createTextNode(key);
        th.appendChild(text);
        row.appendChild(th);
    });
}

// addSummaryDataRows adds "key" and "count" data
function addKeyCountDataRows(table, data) {
    for (const d of data) {
        let row = table.insertRow();
        let k = row.insertCell();
        let t = document.createTextNode(d.key);
        k.appendChild(t)

        let v = row.insertCell();
        const count = numberWithComma(d.count);
        t = document.createTextNode(count);
        v.appendChild(t)
    }
}

// numberWithComma to convert an int to "," separated
function numberWithComma(n) {
    return parseInt(n, 10).toLocaleString('en-US', {maximumFractionDigits: 0})
}