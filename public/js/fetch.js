"use strict";

// fetches a API response.
function fetchHK(url) {
    return new Promise((resolve, reject) => {
        try {
            let req = new XMLHttpRequest();
            req.open("GET", url);
            req.setRequestHeader("Content-Type", "application/json");
            req.send();
            req.onload = () => {
                if (req.status >= 200 && req.status < 300) {
                    try {
                        const res = JSON.parse(req.responseText);
                        resolve(res)
                    } catch (e) {
                        reject({status: req.status, err: `${e}`, text: req.responseText});
                    }
                }
            };
            req.onerror =  () => {
                reject({
                    status: req.status,
                    text: req.statusText,
                    err: "response error"
                });
            };

        } catch (e) {
            reject({err: `${e}`});
        }
    });
}