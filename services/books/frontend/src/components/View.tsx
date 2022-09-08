import React, { useState } from 'react';
import { Book, getBooks } from '../apis/api';

import "./View.css";

function View() {
    const [readValue, setReadValue] = useState<Book[]>();

    async function handleView() {
        var books = await getBooks();
        setReadValue(books);
    }

    function hideView() {
        var books:Book[] = [];
        setReadValue(books);
    }

    return (
        <section>
            <div className="view">
                {readValue?.map((book, key) => 
                    <div className="view__container" key={key}>
                        <ul>
                            <li><span>ID</span><br></br>{book.id}</li>
                            <li><span>Title</span><br></br>{book.title}</li>
                            <li><span>Subtitle</span><br></br>{book.subtitle}</li>
                            <li><span>Author</span><br></br>{book.author}</li>
                            <li><span>ISBN</span><br></br>{book.isbn}</li>
                            <li><span>Edition</span><br></br>{book.edition}</li>
                            <li><span>Year</span><br></br>{book.year}</li>
                            <li><span>Shelf Name</span><br></br>{book.shelf_name}</li>
                            <li><span>Shelf Level</span><br></br>{book.shelf_level}</li>
                        </ul>
                    </div>
                )}
            </div>
            <div className="btns">
                <button className="get btn" onClick={handleView}>View Books!</button>
                <button className="hide btn" onClick={hideView}>Hide Books!</button>
            </div>
        </section>
    )
}

export default View;