"use strict";

// displayError handler if a API returns error
function displayError(errElementID, e) {
    let el = document.getElementById(errElementID);
    let text = document.createTextNode(" Please upload data first.");
    el.appendChild(text);
    console.log(`Error: ${JSON.stringify(e)}`);
}

// insertKeyValueTable inserts Key/Value pairs in table with ID
function insertKeyValueTable(tid, d) {
    let t = document.getElementById(tid); // Table with ID
    addHeadingRow(t, ["Type", "Count"]);
    addKeyCountDataRows(t, d);
}

// insertSummaryTable inserts Key/Value pairs in table with ID
function insertAllKeyValueTable(tid, d) {
    let t = document.getElementById(tid); // Table with ID
    addHeadingRow(t, ["Source", "Type", "Metadata Key", "Count"]);
    addNameTypeKeyCountDataRows(t, d);
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

// addSummaryDataRows adds "key" and "count" data
function addNameTypeKeyCountDataRows(table, data) {
    for (const d of data) {
        let row = table.insertRow();
        let c, t;
        c = row.insertCell();
        t = document.createTextNode(d.name);
        c.appendChild(t)
        c = row.insertCell();
        t = document.createTextNode(d.type);
        c.appendChild(t)
        c = row.insertCell();
        t = document.createTextNode(d.data.key);
        c.appendChild(t)
        c = row.insertCell();
        const count = numberWithComma(d.data.count);
        t = document.createTextNode(count);
        c.appendChild(t)
    }
}

// numberWithComma to convert an int to "," separated
function numberWithComma(n) {
    return parseInt(n, 10).toLocaleString('en-US', {maximumFractionDigits: 0})
}