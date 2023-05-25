class testimonials {
    #image = "";
    #commentary = "";
    #author = "";

    constructor(image, commentary, author) {
        this.#image = image;
        this.#commentary = commentary;
        this.#author = author;
    };

    get image() {
        return this.#image;
    };

    get commentary() {
        return this.#commentary;
    };

    get author() {
        return this.#author;
    };

    get markUp() {
        return `
        <div class="card">
            <img src="${this.image}">
            <p>"${this.commentary}"</p>
            <p class="author">- ${this.author}</p>
        </div>
        `
    };
};

const cardData_1 = new testimonials(
    "Assets/Images/minnie.jpg",
    "Keren 😳",
    "Minnie"
);

const cardData_2 = new testimonials(
    "Assets/Images/yuqi.jpg",
    "Apa ini? Siapa suruh begini ha?",
    "Yuqi"
);

const cardData_3 = new testimonials(
    "Assets/Images/shuhua.jpg",
    "Sopan santunnya lebih diterapkan lagi ya kak",
    "Shuhua"
);

let wholeCardData = [cardData_1, cardData_2, cardData_3];
let markUp = "";

for (let i = 0; i < wholeCardData.length; i++) {
    markUp += wholeCardData[i].markUp;
};

document.getElementById("card-wrap").innerHTML = markUp;