import React, { useContext, useEffect, useState} from "react";

import './main.css';



export const Main = () => {
  /*const [property,setProperties] = useState([]);
  useEffect(()=>{
    fetch('http://localhost:8090/properties/all',
    {method:"GET",
    mode: 'no-cors'})
    .then((response) => response.json())
    },[])*/
  return (
    <main>
      <div class="baner">
        <p>
          BIENVENIDOS A LIBERTADOR INMOBILIARIOS
        </p>
      </div>
    </main>
  )
}