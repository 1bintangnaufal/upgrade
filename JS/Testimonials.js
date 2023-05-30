const promise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open("GET", "https://api.npoint.io/c23654399f5d9bc44e23", true);
    // console.log(xhr);
    xhr.onload = () => {
        if (xhr.status === 200) {
            resolve(JSON.parse(xhr.response));
        } else {
            reject("Error loading card data.");
        };
    };

    xhr.onerror = () => {
        reject("Network problem.");
    };

    xhr.send();
});

async function showAll() {
    const response = await promise;

    let markUp = "";

    response.forEach(function(item) {
        markUp += ` <div class="card">
                        <img src="${item.image}">
                        <p>"${item.commentary}"</p>
                        <p class="author">${item.rating} <i class="fa-solid fa-star fa-fw"></i> from ${item.author}</p>
                    </div>`;
    });

    document.getElementById("card-wrap").innerHTML = markUp;
};

showAll();

async function sort(rating) {
    const response = await promise;

    const sorted = response.filter((item) => {
        return item.rating === rating;
    });

    let markUp = "";

    if(sorted.length === 0) {
        markUp = `<h4>Empty</h4>`;
    } else {
        sorted.forEach((item) => {
            markUp += ` <div class="card">
                            <img src="${item.image}">
                            <p>"${item.commentary}"</p>
                            <p class="author">${item.rating} <i class="fa-solid fa-star fa-fw"></i> from ${item.author}</p>
                        </div>`;
        });
    };

    document.getElementById("card-wrap").innerHTML = markUp;
};