import React, { useContext, useEffect, useState} from "react";
//import "bootstrap/dist/css/bootstrap.min.css";

import Cookies from "universal-cookie";
import swal from "sweetalert2";
const Cookie = new Cookies();
const id_user = setUser()

function setUser (){
    let cookieUser = Cookie.get("user")
  
    if(cookieUser!=undefined){
    let array = cookieUser.spult(",")
    return array[0]
    }else{
      return "undefined"
    }
  }

/*export const InsertProperty =()=>{  

    return <div>Pagina en Creacion</div>
}*/

export const InsertProperty = (props) => {
    const cookies = new Cookies();
    const id_user = cookies.get("user").split(",")[0];
    const [formData, setFormData] = useState({
      Id: "",
      Tittle: "",
      Description: "",
      Size: "",
      Rooms: "",
      Bathrooms: "",
      Service: "",
      Price: "",
      Image: "",
      UserId: id_user,
      Street: "",
      City: "",
      State: "",
      Country: "",
    });
  
    const handleChange = (event) => {
      const { name, value } = event.target;
      setFormData({
        ...formData,
        [name]: value,
      });
    };
  
    const handleSubmit = (event) => {
      event.preventDefault();
      const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      };
      fetch("http://localhost:8090/product", requestOptions)
        .then((response) => {
          if (response.status !== 201) {
            alert("No se pudo agregar el producto");
          } else {
            alert("Producto agregado con éxito");
            setFormData({
              Id: "",
              Tittle: "",
              Description: "",
              Size: "",
              Rooms: "",
              Bathrooms: "",
              Service: "",
              Price: "",
              Image: "",
              UserId: id_user,
              Street: "",
              City: "",
              State: "",
              Country: "",
            });
          }
        })
        .catch((error) => {
          console.error("Error:", error);
          alert("No se pudo agregar el producto");
        });
    };
  
};

