import { FaCubes, FaFile, FaHome } from "react-icons/fa";
import { FaGear, Fa1, Fa2, Fa3, Fa4 } from "react-icons/fa6";
import css from "./navbar.module.css";

const Navbar = () => {
  return (
    <nav className={css.navbar}>
      <a href="/"><FaHome /> Home</a>
      <a href="/colors"><Fa1 />Colors</a>
      <a href="/game-of-life"><Fa2 />Game of life</a>
      <a href="/time-date"><Fa3 />Time date</a>
      <a href="/images"><Fa4 />Images</a>
      <a href="/custom"><FaCubes /> Custom</a>
      <a href="/files"><FaFile /> Files</a>
      <div className="flex-1" />
      <a href="/configuration"><FaGear /> Configuration</a>
    </nav>
  );
}

export default Navbar
