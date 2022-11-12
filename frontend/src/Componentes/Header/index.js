import React, { useContext, useEffect, useState} from "react";
import Libertador from "../../images/logo.jpeg";
//import SearchCountry from "./Topics";
//import SearchCity from "./Topics";
//import SearchService from "./Topics";
import './header.css';

export const Header = () => {
  const [SearchCountries,setSearchContry] = useState([]);
  const [SearchCities,setSearchCity] = useState([]);
  const [SearchServices,setSearchService] = useState([]);
  const fetchApi1 = async()=>{
    const response = await fetch('http://localhost:8090/properties/country')
    .then((response) => response.json());
    setSearchContry(response);
    };
    useEffect(()=>{
    fetchApi1();
    },[])

    const fetchApi2 = async()=>{
      const response = await fetch('http://localhost:8090/properties/city')
      .then((response) => response.json());
      setSearchCity(response);
      };
      useEffect(()=>{
      fetchApi2();
      },[])

      const fetchApi3 = async()=>{
        const response = await fetch('http://localhost:8090/properties/service')
        .then((response) => response.json());
        setSearchService(response);
        };
        useEffect(()=>{
        fetchApi3();
        },[])
  return (
    <header class="header">
      <div class="logo-div">
        <img class="logo" src={Libertador} />
      </div>
      <div class="buscar">
        BUSCA TU HOGAR IDEAL
      </div>
      <div class="search-padre">
        <div class="search-hijo">
          <div class="select">
            <select list="pais-PT" name="pais-PT" id="pais-PT">
            <option value="" selected disabled>Pais</option>
            {
                SearchCountries.map(country =>(
                  <option value="">{country}</option> 
                ))
            }
            </select>
          </div>
          <div class="lista">
            <div class="select">
              <select list="ciudad-PT" name="ciudad-PT" id="ciudad-PT">
              <option value="" selected disabled>Ciudad</option>
            {
                SearchCities.map(city =>(
                  <option value="">{city}</option> 
                ))
            }
              </select>
            </div>
          </div>
          <div class="lista">
            <div class="select">
              <select list="ciudad-PT" name="ciudad-PT" id="ciudad-PT">
              <option value="" selected disabled>Servicio</option>
            {
                SearchServices.map(service =>(
                  <option value="">{service}</option> 
                ))
            }
              </select>
            </div>
          </div>
          <div class="div-search">
            <button class="btn-search" type="submit">
              Buscar
            </button>
          </div>
        </div>
      </div>
    </header>
  )
}

