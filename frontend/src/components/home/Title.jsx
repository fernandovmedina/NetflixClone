import React, {useState, useEffect} from "react"
import Play from '../../img/play.png'
import More from '../../img/more.png'
import Expand from '../../img/expand.png'
import Chapter from './Chapter'
import '../../styles/home/Title.css'

export const Title = ({serie}) => {
    // Variable que guarda la url del wallper de la serie
    let serieWallper = `http://127.0.0.1:8080/api/series/${serie}/img`
    // Variable que guarda la informacion de algun capitulo
    const [capitulo, setCapitulo] = useState([])
    // Obtenemos la informacion de cada capitulo
    const dataCapitulo = (numeroTemporada, numeroCapitulo) => {
        useEffect(() => {
            fetch(`http://127.0.0.1:8080/api/series/${serie}/${numeroTemporada}/${numeroCapitulo}`)
                .then(response => response.json())
                .then(data => setCapitulo(data))
                .catch(err => console.error(err))
        }, [])
        return capitulo
    }
    // Variable que guarda el numero de capitulos de alguna temporada
    const [numeroCapitulos, setNumeroCapitulos] = useState(0)
    // Obtenemos el numero de capitulos
    const dataNumeroCapitulos = (numeroTemporada) => {
        useEffect(() => {
            fetch(`http://127.0.0.1:8080/api/series/capitulos/${serie}/temporadas/${numeroTemporada}`)
                .then(response => response.json())
                .then(data => setNumeroCapitulos(data))
                .catch(err => console.error(err))
        }, [])
        return numeroCapitulos
    }
    // Variable que guarda el wallper del capitulo
    const [chapterWallper, setChapterWallper] = useState('')
    // Obtenemos la ruta de la img
    const dataChapterWallper = (numeroTemporada, numeroCapitulo) => {
        useEffect(() => {
            fetch(`http://127.0.0.1:8080/api/series/${serie}/${numeroTemporada}/${numeroCapitulo}/img`)
                .then(data => setChapterWallper(data))
                .catch(err => console.error(err))
        }, [])
        return chapterWallper
    }
    // Arreglo que contendra los chapter box
    const divs = []
    // For para renderizar los capitulos de cierta temporada
    for(let i = 1; i <= dataNumeroCapitulos(1); i++) {
        divs.push(
            <div>
                <Chapter 
                    imgSrc={dataChapterWallper(1, i)}
                    number={i}
                    name={dataCapitulo(1, i).data.nombre_capitulo}
                    duracion={dataCapitulo(1, i).data.duracion}
                />
            </div>
        )
    }

    console.log(divs)

    return (
        <>
        <div className="title">
            <img src={serieWallper} className="active" />
            <div className="hover-container inactive">
                <img src={serieWallper} />
                <div className="hover-info" >
                    <div className="hover-icons-info">
                        <div className="hover-icons-info-left">
                            <img src={Play} />
                            <img src={More} />                            
                        </div>
                        <div className="hover-icons-info-right">
                            <img src={Expand} id="expand" />
                        </div>
                    </div>
                    <div className="hover-info-info">
                        <p>98% para ti</p>
                        <p>1h 29 min</p>
                    </div>
                </div>
            </div>
        </div>

        <div className="title-window hidden" id="window">
            <div className="title-window-img">
                <img src={serieWallper} />
            </div>
            <div className="title-window-chapters">
                {divs}
            </div>
        </div>
        </>
    );
}
