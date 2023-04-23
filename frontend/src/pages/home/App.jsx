import Navbar from '../../components/home/Navbar'
import Hero from '../../components/home/Hero'
import ScrollTitle from '../../components/home/ScrollTitle'
import './App.css'

function App() {
  return (
    <>
      <Navbar />
      <Hero />
      <ScrollTitle 
        title="Populares en netflix" 
        serie1="Breaking-Bad"
        serie2="How-to-sell-drugs-online"
        serie3="SEX-EDUCATION"
        serie4="Cobra-Kai"
        serie5="Guerra-de-Vecinos"
      />
      <ScrollTitle 
        title="Continuar viendo de Fernando Alejandro" 
        serie1="Better-Call-Saul"
        serie2="DAYBREAK"
        serie3="El-Juego-del-Calamar"
        serie4="NARCOS-MEXICO"
        serie5="Gambito-de-Dama"
      />
      <ScrollTitle 
        title="Volver a verlo" 
        serie1="Cobra-Kai"
        serie2="How-to-sell-drugs-online"
        serie3="NARCOS-MEXICO"
        serie4="Gambito-de-Dama"
        serie5="DAYBREAK"
      />
    </>
  )
}

export default App
