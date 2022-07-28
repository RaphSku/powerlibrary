import axios from 'axios';
import { Book } from "../components/Header";

export async function getBooks() {
    const response = await axios.get("http://localhost:9090/api/v1/books/");
    if (response.status !== 200) {
        console.log(response);
    }
    var books:Book[] = response.data;

    return books
}

export async function postBook(payload: Book) {
    const response = await axios.post("http://localhost:9090/api/v1/book/", payload, {headers: {"Content-type": "application/json; charset=UTF-8"}});
    if (response.status !== 200) {
        console.log(response);
    }
}