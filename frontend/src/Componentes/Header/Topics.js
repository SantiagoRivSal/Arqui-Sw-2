import React from "react";

export const SearchCountry =(
    {id,
    country,
})=>{
    return(
        <div className="Country">
            <select><option value="">{country}</option></select>
            
        </div>
    )
}

export const SearchCity =(
    {id,
    city,
})=>{
    return(
        <div className="City">
            <option value="">{city}</option>
        </div>
    )
}

export const SearchService =(
    {id,
    service,
})=>{
    return(
        <div className="Service">
            <option value="">{service}</option>
        </div>
    )
}