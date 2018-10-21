const availableCardTypes = {
    amex: {
        prefix: ["37", "34"],
        digitCount: 15,
        cvvCount: 4
    },
    visa: {
        prefix: ["4"],
        digitCount: 16,
        cvvCount: 3
    },
    mastercard: {
        prefix: ["51", "52", "53", "54", "55"],
        digitCount: 16,
        cvvCount: 3
    }
};

function calculateLuhn(digits) {
    let currentSum = sum(digits, false);

    return (currentSum * 9) % 10;
}

/* Sum each digit from right to left, and double every second digit.
   If the double exceeds 9, then sum its digits (i.e., 654321 -> 358341 -> 24) */
function sum(digits, even) {
    const computed = [0, 2, 4, 6, 8, 1, 3, 5, 7, 9];
    let sum = 0,
        digit = 0,
        totalDigits = digits.length;

    while (totalDigits--) {
        digit = Number(digits[totalDigits]);
        sum += (even = !even) ? computed[digit] : digit;
    }

    return sum;
}

function getRandomPrefix(prefix) {
    return prefix[Math.floor(Math.random() * prefix.length)];
}

export function generatePAN(cardType) {
    if (availableCardTypes[cardType]) {
        let cardNumber = getRandomPrefix(availableCardTypes[cardType].prefix);
        while (cardNumber.length < availableCardTypes[cardType].digitCount - 1) {
            cardNumber = cardNumber + Math.floor(Math.random() * 10);
        }

        return cardNumber + calculateLuhn(cardNumber)
    } else {
        console.log(`Error, ${cardType} is not valid...`)
    }
}

export function generateExpiryDate() {
    let randomMonth = Math.floor(Math.random() * 12) + 1;
    if (randomMonth < 10) {
        randomMonth = "0" + randomMonth;
    }
    let yearNow = new Date(Date.now()).getFullYear();
    let randomYear = yearNow + (Math.floor(Math.random() * 4) + 2);

    return {
        month: randomMonth,
        year: `${randomYear.toString()[2]}` + `${randomYear.toString()[3]}`,
    };
}

export function generateCVV(cardType) {
    if (availableCardTypes[cardType]) {
        let cvv = "";
        while (cvv.length < availableCardTypes[cardType].cvvCount) {
            cvv = cvv + Math.floor(Math.random() * 10);
        }

        return cvv
    } else {
        console.log(`Error, ${cardType} is not valid...`)
    }
}

export function generateFormattedPAN(cardType) {
    switch (cardType) {
        case 'amex':
            return formatAmexPan(generatePAN(cardType));
        case 'visa':
        case 'mastercard':
            return formatStandardPan(generatePAN(cardType));
        default:
            console.log(`Card type  ${cardType}  is not recognised...`);
            return null
    }
}

function formatAmexPan(pan) {
    let formattedPan = "";
    for (let count = 0; count < pan.length; count++) {
        formattedPan = formattedPan + pan[count];
        if (count === 3 || count === 9) {
            formattedPan = formattedPan + " ";
        }
    }

    return formattedPan;
}

function formatStandardPan(pan) {
    let formattedPan = "";
    for (let count = 0; count < pan.length; count++) {
        formattedPan = formattedPan + pan[count];
        if ((count + 1) % 4 === 0) {
            formattedPan = formattedPan + " ";
        }
    }

    return formattedPan;
}