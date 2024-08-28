import { Outlet } from "react-router-dom";
import Navbar from "../../components/navbar"

const Template = () => {
  return (
    <div className="flex h-screen w-screen overflow-x-hidden">
      <Navbar />
      <div className="flex-1 flex p-2 flex-col overflow-y-auto">
        <Outlet />
      </div>
    </div>
  );
};

export default Template;
