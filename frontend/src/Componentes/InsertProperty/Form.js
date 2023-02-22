import React from 'react'
import './insert.css'

function imagen(id) {
    var x = document.getElementById("imagen").value;
    if(x == ""){
    document.getElementById("div-imagen-form").style.display = "none";
    }
    else{
        document.getElementById("div-imagen-form").style.display = "block";
        document.getElementById(id).src = x;
    }
 
}

const Form = ({ onChange, onSubmit, form }) => (
    <div className="container" id="container">
        <h1 className='publicar'>
            Publicar propiedad
        </h1>
        <form className='form-publicar'
            onSubmit={onSubmit}
        >
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Tittle"
                    name="tittle"
                    onChange={onChange}
                    value={form.tittle}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Description"
                    name="description"
                    onChange={onChange}
                    value={form.description}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Image"
                    name="image"
                    id="imagen"
                    onChange={onChange}
                    onBlur={() => imagen("imagen-2")}
                    value={form.image}
                />
            </div>
            <div className='div-imagen-form' id="div-imagen-form">
            <img className="imagen" id="imagen-2" src='' />
            </div>
            
            <div className="form-group">
                <input
                    type="text"
                    min="1"
                    className="form-control, inputpb"
                    placeholder="Size"
                    name="size"
                    onChange={onChange}
                    value={form.size}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputpb"
                    placeholder="Rooms"
                    name="rooms"
                    onChange={onChange}
                    value={form.rooms}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    min="1"
                    className="form-control, inputpb"
                    placeholder="Bathrooms"
                    name="bathrooms"
                    onChange={onChange}
                    value={form.bathrooms}
                />
            </div>
            <div className="form-group">
                <select
                    type="text"
                    className="form-control, inputpc"
                    name="service"
                    onChange={onChange}
                    value={form.service}>
                    <option value="" selected disabled>Service</option>
                    <option value="Alquiler">Alquiler</option>
                    <option value="Venta">Venta</option>
                </select>
            </div>
            <div className="form-group">
                <input
                    type="text"
                    min="1"
                    className="form-control, inputpb"
                    placeholder="Price"
                    name="price"
                    onChange={onChange}
                    value={form.price}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Street"
                    name="street"
                    onChange={onChange}
                    value={form.street}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="City"
                    name="city"
                    onChange={onChange}
                    value={form.city}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="State"
                    name="state"
                    onChange={onChange}
                    value={form.state}
                />
            </div>
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Country"
                    name="country"
                    onChange={onChange}
                    value={form.country}
                />
            </div>
            <div>
                <button
                    type="submit"
                    className="btn-summit"
                >
                    Publicar
                </button>
            </div>
        </form>
    </div>
)



export default Form