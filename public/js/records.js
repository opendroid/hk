"use strict";


// Fetches records summary
function recordsSources(tableID, errElementID) {
    fetchHK("/v1/recordsSources")
        .then(d => insertKeyValueTable(tableID, d))
        .catch(e => displayError(errElementID, e));
}

// Fetches records summary
function recordsTypes(tableID, errElementID) {
    fetchHK("/v1/recordsTypes")
        .then(d => insertKeyValueTable(tableID, d))
        .catch(e => displayError(errElementID, e));
}

// Fetches records summary
function recordsAll(tableID, errElementID) {
    fetchHK("/v1/recordsAll")
        .then(d => insertAllKeyValueTable(tableID, d))
        .catch(e => displayError(errElementID, e));
}
