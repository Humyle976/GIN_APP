import { Outlet } from "react-router-dom";
import Sidebar from "./components/Sidebar";

function AppLayout() {
  return (
    <div className="flex bg-black">
      <Sidebar />
      <Outlet />
    </div>
  );
}

export default AppLayout;
