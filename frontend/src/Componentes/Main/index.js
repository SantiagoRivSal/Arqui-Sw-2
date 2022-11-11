import React, { useContext, useEffect, useState} from "react";
import { PropertyItems } from "./PropertyItems";

import './main.css';



export const Main = () => {
  const [properties,setProperties] = useState([]);

  const fetchApi = async()=>{
    const response = await fetch('http://localhost:8090/properties/all')
    .then((response) => response.json());
    setProperties(response);
    };
    useEffect(()=>{
    fetchApi();
    },[])
  return (
    <main>
      <div class="baner">
        <p>
          BIENVENIDOS A LIBERTADOR INMOBILIARIOS
        </p>
      </div>
      <div className="Property">
      {
                properties.map(property =>(
                  <PropertyItems key={property.id}
                  id={property.id}
                  tittle ={property.tittle}
                  size={property.size}
                  bathrooms={property.bathrooms}
                  service={property.service}
                  city={property.address.city}
                  state={property.address.state}
                  country={property.address.country}
                  street={property.address.street}
                  price={property.price}
                  rooms={property.rooms}
                  image={property.image}
                  description={property.description}
                  /> 
                ))
            }

        </div>
    </main>
  )
}