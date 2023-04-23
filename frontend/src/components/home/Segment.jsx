import React from 'react'
import '../../styles/home/Segment.css'

const Segment = (props) => {
    return (
        <div className="segment-container">
            <h2>{props.title}</h2>
        </div>
    );
}

export default Segment
