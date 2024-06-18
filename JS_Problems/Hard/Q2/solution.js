function filterAndTransform(people) {
    // Filter out people under 18 years old
    const adults = people.filter(person => person.age >= 18);

    // Transform the filtered objects into strings
    const transformed = adults.map(person => {
        return `${person.name} is ${person.age} years old and loves ${person.hobby}.`;
    });

    return transformed;
}

const people = [
    { name: "Kare", age: 22, hobby: "skiing" },
    { name: "Leo", age: 24, hobby: "equestrianism" }
];
console.log(filterAndTransform(people))
