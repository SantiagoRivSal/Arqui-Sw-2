import { PropertyItems } from "../Main/PropertyItems";
import React, { useEffect, useState } from "react";
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
    
  const [SearchCountries, setSearchContry] = useState([]);
  const [SearchCities, setSearchCity] = useState([]);
  const [SearchServices, setSearchService] = useState([]);
  
  const [valueCity, setValueCity] = useState("");
  const [valueCountry, setValueCountry] = useState("");
  const [valueService, setValueService] = useState("");
  
    const [busqueda, setBusqueda]= useState("");
    const fetchBuscador = async()=>{
      
        const response = await fetch('http://localhost:8000/search/tittle_*'+ busqueda+"*")
        /*.then((response) => response.json());
        if (response.status == 400) {
          swal.fire({
            icon: 'error',
            text: "No se encontro el producto",
          }).then((result) => {
            if (result.isConfirmed) {
                window.location.reload();
            }})
       }else{
        setProperties(response)
        console.log(response);
       }*/
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

    const handleChange2=e=>{
     setBusqueda(e.target.value);
      };

      const handleSubmit2= (event)=>{
        event.preventDefault();
        fetchBuscador();
    };


  const fetchApi1 = async () => {
    const response = await fetch('http://localhost:8090/properties/country')
      .then((response) => response.json());
    setSearchContry(response);
  };
  useEffect(() => {
    fetchApi1();
  }, [])

  const fetchApi2 = async () => {
    const response = await fetch('http://localhost:8090/properties/city')
      .then((response) => response.json());
    setSearchCity(response);
  };
  useEffect(() => {
    fetchApi2();
  }, [])

  const fetchApi3 = async () => {
    const response = await fetch('http://localhost:8090/properties/service')
      .then((response) => response.json());
    setSearchService(response);
  };
  useEffect(() => {
    fetchApi3();
  }, [])

  const fetchApiProperty = async () => {
    const search = await fetch("http://localhost:8000/searchAll/*" +  valueService+ "_" + valueCountry + "_*" + valueCity+"*")
      .then((res) => res.json())
      if(search===null){
        swal.fire({
          icon: 'error',
          text: "NO HAY PROPIEDADES CON LOS FILTROS SELECCIONADOS",
        }) 
        .then((result) => {
          if (result.isConfirmed) {
              window.location.reload();
          }})
    }else{
    setProperties(search)
    console.log(search); 
      }
  };
  const handleChange = e => {
    setValueCity(e.target.value);
    setValueCountry(e.target.value);
    setValueService(e.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    fetchApiProperty();
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
            placeholder="Buscador de Productos"
            onChange={handleChange2}
          
          />
          </div>
          <div>
            <button class="btn-search" type="button" onClick={handleSubmit2}>
                Buscar
              </button>
          </div>
        </div>
      </div>

      <div class="search-padre">
        <div class="search-hijo">
          <div class="select">
            <select list="pais-PT" name="pais-PT" id="pais-PT" valueCountry={valueCountry} onChange={(event) => setValueCountry(event.target.value)}>
              <option value="" selected disabled>Pais</option>
              {
                SearchCountries.map(country => (
                  <option value={country} onChange={handleChange}>{country}
                  </option>
                ))
              }
            </select>
          </div>
          <div class="lista">
            <div class="select">
              <select list="ciudad-PT" name="ciudad-PT" id="ciudad-PT" valueCity={valueCity} onChange={(event) => setValueCity(event.target.value)}>
                <option value="" selected disabled>Ciudad</option>
                {
                  SearchCities.map(city => (
                    <option value={city} onChange={handleChange}>{city}
                    </option>
                  ))
                }
              </select>
            </div>
          </div>
          <div class="lista">
            <div class="select">
              <select list="ciudad-PT" name="ciudad-PT" id="ciudad-PT" valueService={valueService} onChange={(event) => setValueService(event.target.value)}>
                <option value="" selected disabled>Servicio</option>
                {
                  SearchServices.map(service => (
                    <option value={service} onChange={handleChange}>{service}
                    </option>
                  ))
                }
              </select>
            </div>
          </div>
          <div class="div-search">
            <button class="btn-search" type="button" onClick={handleSubmit}>
              Aplicar
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