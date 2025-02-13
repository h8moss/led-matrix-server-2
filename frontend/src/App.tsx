import { createBrowserRouter, createRoutesFromElements, Route, RouterProvider } from 'react-router-dom'
import './App.css'
import Template from "./ui/pages/template"
import HomePage from './ui/pages/home'
import TimeDatePage from './ui/pages/timeDate'
import ImagesPage from './ui/pages/images'
import GameOfLifePage from './ui/pages/gameOfLife'
import FilesPage from './ui/pages/files'
import CustomPage from './ui/pages/custom'
import ConfigurationPage from './ui/pages/configuration'
import ColorsPage from './ui/pages/colors'

const router = createBrowserRouter(createRoutesFromElements(
  <Route path="/" element={<Template />} >
    <Route index element={<HomePage />}/>

    <Route path="colors" element={<ColorsPage />}/>
    <Route path="config" element={<ConfigurationPage />}/>
    <Route path="custom" element={<CustomPage />}/>
    <Route path="files" element={<FilesPage />}/>
    <Route path="game-of-life" element={<GameOfLifePage />}/>
    <Route path="images" element={<ImagesPage />}/>
    <Route path="time-date" element={<TimeDatePage />}/>
  </Route>
))

function App() {
  return (
    <RouterProvider router={router} />
  )
}

export default App
