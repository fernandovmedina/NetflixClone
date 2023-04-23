import React from "react"
import HeroTitle from '../../img/hero-title.webp'
import Play from '../../img/play_black.png'
import Info from '../../img/info.png'
import '../../styles/home/Hero.css'

const Hero = () => {
    return (
        <div className="hero">
            <div className="hero-2">
                <div className="inner-hero">
                    <img src={HeroTitle} />
                    <div className="buttons">
                        <button><img src={Play} />Reproducir</button>
                        <button><img src={Info} />Más información</button>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Hero
