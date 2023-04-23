import NetflixLogo from '../../img/netflix.png'
import ProfileLogo from '../../img/profile-logo.png'
import Search from '../../img/search.png'
import Notification from '../../img/notification.png'
import '../../styles/home/Navbar.css'

const Navbar = () => {
    return (
        <nav>
            <div className="left-navbar">
                <a href="/"><img src={NetflixLogo} /></a>
                <ul>
                    <li><a href="/">Inicio</a></li>
                    <li><a href="#">Series</a></li>
                    <li><a href="#">Películas</a></li>
                    <li><a href="#">Novedades populares</a></li>
                    <li><a href="#">Mi lista</a></li>
                    <li><a href="#">Explora por idiomas</a></li>
                </ul>
            </div>
            <div className="right-navbar">
                <img src={Search} className="icon" />
                <a href="#">Niños</a>
                <img src={Notification} className="icon" />
                <img src={ProfileLogo} />
            </div>
        </nav>
    );
}

export default Navbar
