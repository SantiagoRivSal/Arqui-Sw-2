import React from "react";

export const PropertyItems =(
    {id,
    tittle,
    size,
    bathrooms,
    service,
    city,
    state,
    country,
    street,
    price,
    rooms,
    image,
    description
})=>{

    return(
        <div className="Property">
        <a href="#">
        <div className="Property_img">
            <img className="image" src={image} alt=""/>
        </div>
        </a>
        <div className="Property_footer">
            <h1>{tittle}</h1>
            <h2>{description}</h2>
            <p className="size">{size} metros cuadrados</p>
            <p className="price">U$S {price}</p>
            <p className="Datos">Ambientes: {rooms}</p>
            <p className="Datos">Baños: {bathrooms}</p>
            <p className="Datos">Servicio: {service}</p>
            <p className="Datos">Direccion: {street},{city},{state}</p>
            <p className="Datos">Pais: {country}</p>
        </div>
        </div>
    )
    
}