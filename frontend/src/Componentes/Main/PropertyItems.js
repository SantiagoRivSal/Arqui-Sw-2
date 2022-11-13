import React from "react";

export const PropertyItems = (
    { id,
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
    }) => {

    return (
            <div class="Property">
                <div>
                    <div class="Property_img">
                        <img class="image" src={image} alt="" />
                    </div>
                </div>
                <div class="Property_footer">
                    <h1>{tittle}</h1>
                    <h2>{description}</h2>
                    <p class="price">U$S {price}</p>
                    <p class="Datos">{size} m2</p>
                    <p class="Datos">{rooms} Ambientes</p>
                    <p class="Datos">{bathrooms} Baños</p>
                    <p class="Datos">{service}</p>
                    <p class="Datos-2">Direccion: {street},{city},{state}</p>
                    <p class="Datos-2">{country}</p>
                </div>

            </div>

      



        /*
                <div className="Property">
        
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
                */
    )

}