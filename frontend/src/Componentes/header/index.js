import React, { useEffect, useState } from "react";
import Libertador from "../../images/logo.jpeg";
import {Link} from "react-router-dom";
import swal from "sweetalert2";
import '../Main/main.css';
import Cookies from "universal-cookie";

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
export const Header = ()=>{
    
    return(
        <header>
        {id_user!="undefined"?
        <ul>
            <li>
               <Link to="/home" className="botones_menu">
               INICIO
               </Link> 
            </li>
            <li>
               <Link to="/properties" className="botones_menu">
               CARGAR PRODUCTOS
               </Link> 
            </li>
            <li>
               <Link to="/message" className="botones_menu"> 
                MENSAJES
                </Link> 
            </li>
            <button className="btn" onClick={()=>LogOut()}>
                SALIR
            </button>
        </ul>
            :
               <Link to="/" className="botones_menu">
                </Link> 
             }
        </header>
    )
}