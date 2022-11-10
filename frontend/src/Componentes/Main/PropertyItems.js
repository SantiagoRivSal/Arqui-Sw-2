import React from "react";

export const PropertyItems =(
    {Id,
    Tittle,
    Size,
    Bathrooms,
    Service,
    City,
    State,
    Country,
    Street,
    Price,
    Rooms,
    Image
})=>{


    return(
        <div className="Property">
        <a href="#">
        <div className="Property_img">
            <img className="image" src={Image} alt=""/>
        </div>
        </a>
        <div className="Property_footer">
            <h1>{Tittle}</h1>
            <p className="size">{Size} metros cuadrados</p>
            <p className="price">U$S {Price}</p>
            <p className="Datos">Ambientes: {Rooms}</p>
            <p className="Datos">Baños: {Bathrooms}</p>
            <p className="Datos">Servicio: {Service}</p>
            <p>{description}</p>
        </div>
        </div>
    )
}