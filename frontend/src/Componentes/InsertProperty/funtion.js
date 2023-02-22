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
        [name]: name === "userid"||name === "size" || name === "rooms" || name === "bathrooms" || name === "price" ? Number(value) : value,
      });
    };


    const handleSubmit = async event => {
      event.preventDefault();
      if(Number(form.size)>0 && Number(form.rooms)>0 && Number(form.bathrooms)>0 && Number(form.price) > 0 && form.service != "" && form.tittle != ""  && form.tittle != ""){
        const requestOptions = {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
          'Access-Control-Allow-Origin': '*'
        },
        body: JSON.stringify({
          ...form,
          size: Number(form.size), // Convertir a número
          rooms: Number(form.rooms), // Convertir a número
          bathrooms: Number(form.bathrooms), // Convertir a número
          price: Number(form.price), // Convertir a número
          userid:Number(form.userid)
        }),
      };
      try {
        const response = await fetch('http://localhost:8090/properties/load', requestOptions);
        if (response.ok) {
          const data = await response.json();
          swal.fire({
            icon: 'success',
            text: "Propiedad Agregada",
        }).then((result) => {
          if (result.isConfirmed) {
            window.location.href = "http://localhost:3000/home"
          }});
          // Aquí podrías hacer algo con la respuesta, si es necesario
        } else {
          throw new Error("No se pudo agregar la propiedad");
        }
      } catch (error) {
        console.error("Error:", error);
        swal.fire({
          icon: 'error',
          text: "No se pudo agregar la propiedad",
        })
      }
      console.log(form)}
      else{
        swal.fire({
          icon: 'error',
          text: "No se pudo agregar la propiedad",
        }) 
      }
    };

    return <NewProperty 
    form={form}
    onChange={handleChange}
    onSubmit={handleSubmit}
/>


};
