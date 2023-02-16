import React from 'react'
import Form from './Form'
/*import Card from '../components/Card'
<div className="col-sm ExerciseNew_Card_Space">
<Card 
    {...form}
/>
</div>*/
const NewProperty = ({form, onChange, onSubmit}) => (
    <div className="NewProperty_Lateral_Spaces row">

        <div className="col-sm NewProperty_Form_Space">
            <Form
                onChange={onChange}
                onSubmit={onSubmit}
                form={form}
            />            
        </div>
    </div>
)

export default NewProperty