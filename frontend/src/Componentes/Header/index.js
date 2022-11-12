import React, { useContext, useEffect, useState} from "react";
import Libertador from "../../images/logo.jpeg";
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
    </header>
  )
}

