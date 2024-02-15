const apiUrl = '/api/v1/';

let packSizes = [];

function calculatePacks() {
    items = document.getElementById("items").value;
    document.getElementById("items").value = "";
    url = apiUrl+'packages/'+items;
    fetch(url)
        .then(response => {
            if (!response.ok) {
            throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(packs => {
            results = document.getElementById("items-results");
            results.innerHTML = "";
            for (let i = 0; i < packs.length; i++) {
                if (packs[i] > 0) {
                    results.innerHTML += "<tr><td>" + packs[i] + "</td><td>" + packSizes[i] + "</td></tr>";
                }
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function getPackSizes() {
    url = apiUrl+'sizes';
    fetch(url)
        .then(response => {
            if (!response.ok) {
            throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(items => {
            packSizes = items;
            results = document.getElementById("packs-results");
            results.innerHTML = "";
            for (let i = 0; i < items.length; i++) {
                results.innerHTML += "<tr><td>" + items[i] + "</td></tr>";
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function setPackSizes() {
    packs = document.getElementById("packs").value.split(",");
    document.getElementById("packs").value = "";
    packsInts = packs.map(str => {
        return parseInt(str, 10);
      });

    url = apiUrl+'sizes';
    fetch(url, {
        method: 'PUT',
        headers: {
            'Accept': 'application/json, text/plain, */*',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(packsInts)
    }).
    then(res => res.json())
    .then(res => console.log(res.message));

    getPackSizes();
}
