import './App.css'
import CreateCharacter from './components/CreateCharacter.jsx';
import ViewCharacters from "./components/ViewCharacters.jsx"
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {

  return (
    <>
      <h1>App</h1>
      <ViewCharacters/>
      <CreateCharacter/>
    </>
  )
}

export default App
