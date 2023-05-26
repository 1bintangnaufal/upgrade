const cardData = [
    {
    image: "Assets/Images/minnie.jpg",
    commentary: "Keren 😳",
    rating: 4,
    author: "Minnie"
    },

    {
    image: "Assets/Images/yuqi.jpg",
    commentary: "Apa ini? Siapa suruh begini ha?",
    rating: 2,
    author: "Yuqi"
    },

    {
    image: "Assets/Images/shuhua.jpg",
    commentary: "Sopan santunnya lebih diterapkan lagi ya kak",
    rating: 1,
    author: "Shuhua"
    },

    {
    image: "Assets/Images/soyeon.jpg",
    commentary: "Kalo dikasih kerjaan tuh kerjain yang bener",
    rating: 2,
    author: "Soyeon"
    },

    {
    image: "Assets/Images/soojin.jpg",
    commentary: "Yang sabar ya kak",
    rating: 3,
    author: "Soojin"
    },

    {
    image: "Assets/Images/miyeon.jpg",
    commentary: "Hebat!",
    rating: 4,
    author: "Miyeon"
    },
];

function showAll() {
    let markUp = "";

    cardData.forEach(function (item) {
        markUp += `
        <div class="card">
            <img src="${item.image}">
            <p>"${item.commentary}"</p>
            <p class="author">${item.rating} <i class="fa-solid fa-star fa-fw"></i> from ${item.author}</p>
        </div>
        `;
    });

    document.getElementById("card-wrap").innerHTML = markUp;
};

showAll();

function sort(rating) {
    let markUp = "";

    const sorted = cardData.filter(function (item) {
        return item.rating === rating;
    });

    if (sorted.length === 0) {
        markUp += `
        <h2>Empty</h2>
        `;
    } else {
        sorted.forEach(function (item) {
            markUp += `
            <div class="card">
                <img src="${item.image}">
                <p>"${item.commentary}"</p>
                <p class="author">${item.rating} <i class="fa-solid fa-star fa-fw"></i> from ${item.author}</p>
            </div>
            `;
        });
    };

    document.getElementById("card-wrap").innerHTML = markUp;
};