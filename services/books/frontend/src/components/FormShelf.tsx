import React, { useState } from 'react';
import { Shelf } from '../apis/api';

import "./FormShelf.css";

function FormShelf() {
    const [submitValues, setSubmitValues] = useState<Shelf>({
        id: 0,
        name: "",
        room: "",
        location: ""
    });

    function clear(event: any) {
        (document.getElementById("name") as HTMLInputElement).value     = "";
        (document.getElementById("room") as HTMLInputElement).value     = "";
        (document.getElementById("location") as HTMLInputElement).value = "";

        setSubmitValues({
            id: 0,
            name: "",
            room: "",
            location: ""
        });
    }

    async function handleSubmit(event: any) {
        event.preventDefault();

        fetch(`http://192.168.49.2/shelf?query=mutation+_{create(name:"${submitValues.name}",room:"${submitValues.room}",location:"${submitValues.location}"){id,name,room,location}}`)
            .then((response) => {
                console.log(response);
            })
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
                <form className="form-shelfs">
                    <label>
                        Name<br></br>
                        <input id="name" type="text" name="name" onChange={handleChange} />
                    </label>
                    <label>
                        Room<br></br>
                        <input id="room" type="text" name="room" onChange={handleChange} />
                    </label>
                    <label>
                        Location<br></br>
                        <input id="location" type="text" name="location" onChange={handleChange} />
                    </label>
                </form>
                <div className="btns">
                    <button className="btn" onClick={handleSubmit}>Add Shelf!</button>
                    <button className="btn" onClick={clear}>Clear!</button>
                </div>
            </div>
        </section>
    )
}

export default FormShelf;