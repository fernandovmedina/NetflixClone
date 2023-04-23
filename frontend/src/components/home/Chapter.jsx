import React from "react"
import '../../styles/home/Chapter.css'

const Chapter = ({imgSrc, number, name, duracion}) => {
    return (
        <div className="chapter-container">
            <div className="chapter-img">
                <img src={imgSrc} />
            </div>
            <div className="chapter-info">
                <h1 className="chapter-number">{number}</h1>
                <h1 className="chapter-name">{name}</h1>
                <p>{duracion}</p>
            </div>
        </div>
    )
}

export default Chapter
