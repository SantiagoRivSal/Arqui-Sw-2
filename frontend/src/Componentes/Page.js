import React from "react";
import { Routes, Route } from "react-router-dom";
import { Main } from "./Main/index";
import Login from "./Login/login";


export const Page = () => {
    return (
        <section>
            <Routes>
                <Route path="/" exact element={<Login/>} />
                <Route path="/home" exact element={<Main/>} />
            </Routes>
        </section>
    )
}