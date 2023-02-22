import { PropertyItems } from "../Main/PropertyItems";
import React, { useState } from "react";
import Libertador from "../../images/logo.jpeg";
import swal from "sweetalert2";
import '../Main/main.css';
import Cookies from "universal-cookie";



export const Main = () => {

const Cookie = new Cookies();
    let cookie = Cookie.get("user")
    let id_user;
    if(cookie!=undefined){
    let array = cookie.split(",")
     id_user = array[0]
    }
    else{
         id_user = "undefined"
    }
  const [properties, setProperties] = useState([]);
  
    const [busqueda, setBusqueda]= useState("");
    const fetchBuscador = async()=>{
      
        const response = await fetch('http://localhost:8000/search/tittle_*'+ busqueda+"*")
       .then((response) => response.json())
       if(response===null){
         swal.fire({
           icon: 'error',
           text: "NO HAY PROPIEDADES EN SU BUSQUEDA",
         }) 
         .then((result) => {
           if (result.isConfirmed) {
               window.location.reload();
           }})
     }else{
     setProperties(response)
     console.log(response); 
       }
        };

    const handleChange=e=>{
     setBusqueda(e.target.value);
      };

      const handleSubmit= (event)=>{
        event.preventDefault();
        fetchBuscador();
    };

  return (
    <header class="header">
      <div>
        <div class="logo-div">
        <img class="logo" src={Libertador} />
        </div>
        <div>
          
        </div>
      </div>
      <div class="buscar">
        BUSCA TU HOGAR IDEAL
      </div>

      <div class="search-padre">
        <div>
          <div>
          <input
            class="form-control inputBuscar"
            value={busqueda}
            placeholder="Buscar propiedades"
            onChange={handleChange}
          
          />
          </div>
          <div>
            <button class="btn-search" type="button" onClick={handleSubmit}>
                Buscar
              </button>
          </div>
        </div>
      </div>
        
      <div class="Property-padre">
            {properties
        .filter((property) => id_user != property.userid)
        .map((property) => (
                  <PropertyItems key={property.id}
                    id={property.id}
                    tittle={property.tittle}
                    size={property.size}
                    bathrooms={property.bathrooms}
                    service={property.service}
                    city={property.city}
                    state={property.state}
                    country={property.country}
                    street={property.street}
                    price={property.price}
                    rooms={property.rooms}
                    image={property.image}
                    description={property.description}
                    iduser={property.userid}
                  />
            ))
        }
        {(properties.length > 0 && properties.every((property) => id_user == property.userid)) ? (
      <p>Solo hay propiedades a tu nombre</p>
    ) : null}
        </div>
    </header>
  )
}