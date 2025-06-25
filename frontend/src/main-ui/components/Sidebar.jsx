import { FaHome } from "react-icons/fa";
import { IoMdNotifications } from "react-icons/io";
import { FaMessage } from "react-icons/fa6";
import { FaUser } from "react-icons/fa";
import { IoMdSettings } from "react-icons/io";
import { FaUserFriends } from "react-icons/fa";
import SidebarButton from "./SidebarButton";
import List from "@mui/material/List";
import Box from "@mui/material/Box";
import LogoutButton from "./LogoutButton";

const sidebar = [
  {
    key: 1,
    icon: <FaHome className="text-3xl text-white" />,
    text: "Home",
    path: "/home",
  },
  {
    key: 2,
    icon: <FaUserFriends className="text-3xl text-white" />,
    text: "Friend Requests",
    path: "/friend-requests",
  },
  {
    key: 3,
    icon: <IoMdNotifications className="text-3xl text-white" />,
    text: "Notifications",
    path: "/notifications",
  },
  {
    key: 4,
    icon: <FaMessage className="text-3xl text-white" />,
    text: "Messages",
    path: "/messages",
  },
  {
    key: 5,
    icon: <FaUser className="text-3xl text-white" />,
    text: "Profile",
    path: "/profile",
  },
  {
    key: 6,
    icon: <IoMdSettings className="text-3xl text-white" />,
    text: "Settings",
    path: "/settings",
  },
];
function Sidebar() {
  return (
    <Box className="md:w-1/4 sticky top-0 h-screen flex flex-col items-center border-r border-white overflow-y-auto overflow-x-hidden">
      <List className="flex flex-col justify-center gap-5 flex-grow text-white">
        {sidebar.map((item) => (
          <SidebarButton
            key={item.key}
            icon={item.icon}
            text={item.text}
            path={item.path}
          />
        ))}
      </List>

      <LogoutButton />
    </Box>
  );
}

export default Sidebar;
