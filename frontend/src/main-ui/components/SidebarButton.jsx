import ListItemText from "@mui/material/ListItemText";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import { Link } from "react-router-dom";
function SidebarButton({ text, icon, path }) {
  return (
    <Link to={path}>
      <ListItem className="sticky top-0 lg:justify-center">
        <ListItemIcon>{icon}</ListItemIcon>
        <ListItemText primary={text} className="hidden lg:block" />
      </ListItem>
    </Link>
  );
}

export default SidebarButton;
