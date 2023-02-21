import React from 'react'
import '../InsertProperty/insert.css'


const Message = ({ onChange, onSubmit, form }) => (
    <div className="container">
        <h1 className='publicar'>
            Enviar Mensaje
        </h1>
        <form className='form-publicar'
            onSubmit={onSubmit}
        >
            <div className="form-group">
                <input
                    type="text"
                    className="form-control, inputp"
                    placeholder="Escriba un Mensaje para el Dueño de la Propiedad"
                    name="body"
                    onChange={onChange}
                    value={form.body}
                />
            </div>
            <div>
                <button
                    type="submit"
                    className="btn-summit"
                >
                    Enviar
                </button>
            </div>
        </form>
    </div>
)



export default Message