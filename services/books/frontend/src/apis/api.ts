import axios from 'axios';

export interface Book {
    id: number,
    title: string,
    subtitle: string,
    author: string,
    isbn: string,
    edition: number,
    year: number
}

export async function getBooks() {
    const response = await axios.get("http://localhost:8000/api/v1/books/");
    if (response.status !== 200) {
        console.log(response);
    }
    var books:Book[] = response.data;

    return books
}

export async function postBook(payload: Book) {
    const response = await axios.post("http://localhost:8000/api/v1/book/", payload, {headers: {"Content-type": "application/json; charset=UTF-8"}});
    if (response.status !== 200) {
        console.log(response);
    }
}