import React from "react";

export const PropertyItem =(
    {Id,
    Title,
    Description,
    Address,
    Size,
    Rooms,
    Bathrooms,
    Service,
    Price,
    Image
})=>{


    return(
        <div className="producto">
        <a href="#">
        <div className="producto_img">
            <img className="image" src={Image} alt=""/>
        </div>
        </a>
        <div className="producto_footer">
            <h1>{Title}</h1>
            <p>{Description}</p>
            <p>{Size} Metros Cuadrados</p>
            <p>{Rooms} Ambientes</p>
            <p>{Bathrooms} Baños</p>
            <p className="price">U$S {Price}</p>
        </div>
        </div>
    )
}