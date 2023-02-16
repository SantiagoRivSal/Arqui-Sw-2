import React, { useContext, useEffect, useState} from "react";
import NewProperty from './newProperty'
import Cookies from "universal-cookie";
import swal from "sweetalert2";

export const InsertProperty = () => {
    const cookies = new Cookies();
    const id_user = cookies.get("user").split(",")[0];
    const [form, setForm] = useState({
      'tittle': "",
      'description': "",
      'size':"",
      'rooms':"",
      'bathrooms':"" ,
      'service': "",
      'price':"",
      'image': "",
      'userid': id_user,
      'street': "",
      'city': "",
      'state': "",
      'country': "",
    });
  
    const handleChange = (event) => {
      const { name, value } = event.target;
      setForm({
        ...form,
        [name]: value,
      });
    };

    const handleSubmit = async event => {
      event.preventDefault();
      const requestOptions = {
        method: "POST",
        headers: { 'Content-type': 'application/json'},
        mode: 'no-cors',
        body: JSON.stringify(form),
      };
      await fetch('http://localhost:8090/properties/load', requestOptions)
        .then((response) => {
          if (response.status !== 201) {
            alert("No se pudo agregar la propiedad");
          } else {
            alert("Propiedad agregada con éxito");
            response.json()
          }
        })
        .catch((error) => {
          console.error("Error:", error);
          alert("No se pudo agregar la propiedad");
        });
    };

    return <NewProperty 
    form={form}
    onChange={handleChange}
    onSubmit={handleSubmit}
/>


};
