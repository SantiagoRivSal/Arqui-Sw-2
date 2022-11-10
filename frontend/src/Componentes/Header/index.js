import React from 'react'
import Libertador from "../../images/logo.jpeg";
import PropertyItems from "../Main/"
import './header.css';

export const Header = () => {
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
              <option>Argentina</option>
            </select>
          </div>
          <div class="lista">
            <div class="select">
              <select list="ciudad-PT" name="ciudad-PT" id="ciudad-PT">
                <option value="" selected disabled>Ciudad</option>
                <option>La Rioja</option>
              </select>
            </div>
          </div>
          <div class="lista">
            <div class="select">
              <select list="ciudad-PT" name="ciudad-PT" id="ciudad-PT">
                <option value="" selected disabled>Service</option>
                <option>Alquiler</option>
                <option>Venta</option>
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