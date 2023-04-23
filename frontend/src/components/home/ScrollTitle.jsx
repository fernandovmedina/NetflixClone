import React from "react"
import { Title } from './Title'
import Segment from "./Segment"
import '../../styles/home/ScrollTitle.css'

const ScrollTitle = ({title, serie1, serie2, serie3, serie4, serie5}) => {
    return (
        <>
        <Segment title={title} />
        <div className="titles-container">
            <div className="title-container">
                <Title serie={serie1} />
            </div>
            <div className="title-container">
                <Title serie={serie2} />
            </div>
            <div className="title-container">
                <Title serie={serie3} />
            </div>
            <div className="title-container">
                <Title serie={serie4} />
            </div>
            <div className="title-container">
                <Title serie={serie5} />
            </div>
        </div>
        </>
    );
}

export default ScrollTitle
