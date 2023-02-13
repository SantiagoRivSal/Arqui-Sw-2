import { PropertyItems } from "../Main/PropertyItems";
import React, { useEffect, useState } from "react";
import Libertador from "../../images/logo.jpeg";
import {Link} from "react-router-dom";
import swal from "sweetalert2";
import '../Main/main.css';
import Cookies from "universal-cookie";
const Cookie = new Cookies();
/*async function getPropertiesBySearch(field, query){
  return fetch( "http://localhost:8000/search=" + field + "_" + query, {
    method: "GET",
    header: "Content-Type: application/json"
  }).then(response=>response.json())
}

async function getPropertiesBySearchAll(query){
  return fetch("http://localhost:8000/searchAll=" + query, {
    method: "GET",
    header: "Content-Type: application/json"
  }).then(response=>response.json())
}

async function getProperties(){
  return await fetch("http://localhost:8000/search=*_*", {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}*/
function LogOut(){
  //var resultado = window.confirm('Estas seguro?');
  swal.fire({
      text: "Estas seguro?",
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Yes'
    }).then((result) => {
      if (result.isConfirmed) {
          Cookie.set("user", "undefined,undefined", { path: "/" });
          window.location.replace("/");
      }})
}

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
  //const [failedSearch, setFailedSearch] = useState(false)
  //const [querying, setQuerying] = useState(false)
  //const [query, setQuery] = useState("")
  
  /*async function searchQueryAll(query){
    if (query == "") {
      query = "*"
    }
    await getPropertiesBySearchAll(query).then(response=>{
      if (response != null) {
        if (response.length > 0) {
          setProperties(response)
          //setFailedSearch(false)
        } else {
          setProperties([])
          //setFailedSearch(true)
        }
      } else {
        //setFailedSearch(false)
        getProperties().then(response => setProperties(response))
      }
    })
  }

  async function searchQuery(field, query){
    if(query === ""){
      query = "*"
    }
    await getPropertiesBySearch(field, query).then(response=>{
      if(response != null){
        if(response.length > 0){
          setProperties(response)
         // setFailedSearch(false)
        }else{
          setProperties([])
         // setFailedSearch(true)
        }
      }
      else{
        //setFailedSearch(false)
        getProperties().then(response=>setProperties(response))
      }
    })
  }

  function searchAllDelete(e){
    searchQueryAll(e.target.value);
    //setQuerying(false);
  }

  function searchDelete(field, query){
    searchQuery(field, query);
   // setQuerying(false);
  }
  //ver lo de opciones para adaptarlo

  if(query == "" && properties.length <= 0){
    searchQuery("*","*")
  }*/


    
  const [SearchCountries, setSearchContry] = useState([]);
  const [SearchCities, setSearchCity] = useState([]);
  const [SearchServices, setSearchService] = useState([]);
  
  const [valueCity, setValueCity] = useState("");
  const [valueCountry, setValueCountry] = useState("");
  const [valueService, setValueService] = useState("");
  
  
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
    const search = await fetch("http://localhost:8983/solr/property/select?defType=lucene&fq=city%3A%22" + valueCity + "%22&fq=country%3A%22" + valueCountry + "%22&fq=service%3A%22" + valueService + "%22&indent=true&q.op=OR&q=*%3A*")
      .then((res) => res.json())
      if(search.response.numFound==0){
        swal.fire({
          icon: 'error',
          text: "NO HAY PROPIEDADES CON LOS FILTROS SELECCIONADOS",
        })
        .then((result) => {
          if (result.isConfirmed) {
              window.location.reload();
          }})
    }else{
    setProperties(search.response.docs)
    console.log(search.response.docs); 
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
      <div class="logo-div">
        <img class="logo" src={Libertador} />
      </div>
      <div>
      {id_user!="undefined"?
                <button className="btn" onClick={()=>LogOut()}>
                Log Out
            </button>:
               <Link to="/" className="botones_menu">
                </Link> 
             }
      </div>
      <div class="buscar">
        BUSCA TU HOGAR IDEAL
      </div>

      <div className="containerInput" >
        <input
          class="form-control inputBuscar"
          //value={busqueda}
          placeholder="Buscador de Productos"
          //onChange={handleChange}
         
        />
        <input class="search-button"
        value = "Buscar"
         type = "button"
        //onClick = {handleSubmit}
        />
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
        {properties.length == 0 || (properties.length > 0 && properties.every((property) => id_user == property.userid)) ? (
      <p>Solo hay propiedades a tu nombre</p>
    ) : null}
        </div>
    </header>
  )
}