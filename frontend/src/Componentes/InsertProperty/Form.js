import React from 'react'

const Form = ({ onChange, onSubmit, form }) => (
    <div className="container">
        <form 
            onSubmit={onSubmit}
        >
            <div className="form-group">
                <input 
                    type="text" 
                    className="form-control" 
                    placeholder="Tittle" 
                    name="tittle"
                    onChange={onChange}
                    value={form.tittle}
                />
            </div>
            <div className="form-group">
                <input 
                    type="text" 
                    className="form-control" 
                    placeholder="Description" 
                    name="description"
                    onChange={onChange}
                    value={form.description}
                />
            </div>
            <div className="form-group">
                <input 
                    type="text" 
                    className="form-control" 
                    placeholder="Image" 
                    name="image"
                    onChange={onChange}
                    value={form.image}
                />
            </div>
            <div className="form-group">
                <input 
                    type="number"
                     min="1"
                    className="form-control" 
                    placeholder="Size" 
                    name="size"
                    onChange={onChange}
                    value={form.size}
                />
            </div>
            <div className="form-group">
                <input 
                    type="number"
                     min="1"
                    className="form-control" 
                    placeholder="Rooms" 
                    name="rooms"
                    onChange={onChange}
                    value={form.rooms}
                />
            </div>
            <div className="form-group">
                <input 
                    type="number"
                     min="1"
                    className="form-control" 
                    placeholder="Bathrooms" 
                    name="bathrooms"
                    onChange={onChange}
                    value={form.bathrooms}
                />
            </div>
            <div className="form-group">
                <input 
                    type="text"
                    className="form-control" 
                    placeholder="Service (Alquiler o Venta)" 
                    name="service"
                    onChange={onChange}
                    value={form.service}
                />
            </div>
            <div className="form-group">
                <input 
                    type="number"
                     min="1"
                    className="form-control" 
                    placeholder="Price" 
                    name="price"
                    onChange={onChange}
                    value={form.price}
                />
            </div>
            <div className="form-group">
                <input 
                    type="text"
                    className="form-control" 
                    placeholder="Street" 
                    name="street"
                    onChange={onChange}
                    value={form.street}
                />
            </div>
            <div className="form-group">
                <input 
                    type="text"
                    className="form-control" 
                    placeholder="City" 
                    name="city"
                    onChange={onChange}
                    value={form.city}
                />
            </div>
            <div className="form-group">
                <input 
                    type="text"
                    className="form-control" 
                    placeholder="State" 
                    name="state"
                    onChange={onChange}
                    value={form.state}
                />
            </div>
            <div className="form-group">
                <input 
                    type="text"
                    className="form-control" 
                    placeholder="Country" 
                    name="country"
                    onChange={onChange}
                    value={form.country}
                />
            </div>
            <div className="form-group">
                <input 
                    type="number"
                    min="1"
                    className="form-control" 
                    placeholder="User" 
                    name="userid"
                    onChange={onChange}
                    value={form.userid} disabled
                />
            </div>
            {form.tittle!='' && (form.service === 'Alquiler' || form.service === 'Venta' )&& form.price>=1 && form.bathrooms>=1 && form.size>=1 && form.rooms>=1? (
      // Aquí se renderiza el resto del formulario solo si el valor de "Service" es "Alquiler" o "Venta"
      <div>
            <button 
                type="submit" 
                className="btn btn-primary float-right"
            >
                Submit
            </button>
            </div>
    ) : null}
        </form>
    </div>
)

export default Form