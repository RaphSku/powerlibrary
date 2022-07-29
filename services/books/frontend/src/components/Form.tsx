import React, { useState } from 'react';
import { Book, postBook } from '../apis/api';

import "./Form.css";

function Form() {

    const [submitValues, setSubmitValues] = useState<Book>({
        id: 0,
        title: "",
        subtitle: "",
        author: "",
        isbn: "",
        edition: 0,
        year: 0
    });

    function clear(event: any) {
        (document.getElementById("title") as HTMLInputElement).value    = "";
        (document.getElementById("subtitle") as HTMLInputElement).value = "";
        (document.getElementById("author") as HTMLInputElement).value   = "";
        (document.getElementById("isbn") as HTMLInputElement).value     = "";
        (document.getElementById("edition") as HTMLInputElement).value  = "";
        (document.getElementById("year") as HTMLInputElement).value     = "";

        setSubmitValues({
            id: 0,
            title: "",
            subtitle: "",
            author: "",
            isbn: "",
            edition: 0,
            year: 0
        });
    }

    async function handleSubmit(event: any) {
        event.preventDefault();
        
        const data: Book = {
            id: 0,
            title: submitValues.title, 
            subtitle: submitValues.subtitle, 
            author: submitValues.author,
            isbn: submitValues.isbn,
            edition: Number(submitValues.edition), 
            year: Number(submitValues.year)
        };

        postBook(data)
    }

    function handleChange(event: any) {
        const target = event.target;
        const value  = target.value;
        const attr   = target.name;

        setSubmitValues((previous) => ({...previous, [attr]: value}));
    }

    return (
        <section>
            <div className="add">
                <form className="form-books">
                    <label>
                        Title<br></br>
                        <input id="title" type="text" name="title" onChange={handleChange} />
                    </label>
                    <label>
                        Subtitle<br></br>
                        <input id="subtitle" type="text" name="subtitle" onChange={handleChange} />
                    </label>
                    <label>
                        Author<br></br>
                        <input id="author" type="text" name="author" onChange={handleChange} />
                    </label>
                    <label>
                        ISBN<br></br>
                        <input id="isbn" type="text" name="isbn" onChange={handleChange} />
                    </label>
                    <label>
                        Edition<br></br>
                        <input id="edition" type="number" name="edition" onChange={handleChange} />
                    </label>
                    <label>
                        Year<br></br>
                        <input id="year" type="number" name="year" onChange={handleChange} />
                    </label>
                </form>
                <div className="btns">
                    <button className="btn" onClick={handleSubmit}>Add Book!</button>
                    <button className="btn" onClick={clear}>Clear!</button>
                </div>
            </div>
        </section>
    )
}

export default Form;