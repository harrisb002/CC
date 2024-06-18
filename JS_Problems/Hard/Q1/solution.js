function mergeAndFormat(books, format) {
  // Aggregate books by title and author
  const aggregatedBooks = books.reduce((acc, { title, author, year }) => {
    const key = `${title} by ${author}`;
    if (!acc[key]) {
      acc[key] = {
        title,
        author,
        years: [year],
      };
    } else {
      acc[key].years.push(year);
    }
    return acc;
  }, {});

  // Format the aggregated books based on the specified format
  if (format === "string") {
    return Object.values(aggregatedBooks).map(
      (book) => `${book.title} by ${book.author} (${book.years.join(", ")})`
    );
  }

  return Object.values(aggregatedBooks).map(({ title, author, years }) => ({
    title,
    author,
    years,
  }));
}
