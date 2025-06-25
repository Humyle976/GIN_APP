import { createPortal } from "react-dom";
import Backdrop from "@mui/material/Backdrop";
import CircularProgress from "@mui/material/CircularProgress";

export function GlobalLoader() {
  return createPortal(
    <Backdrop
      open={true}
      sx={{
        position: "fixed",
        top: 0,
        left: 0,
        width: "100vw",
        height: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        zIndex: 9999,
      }}
    >
      <CircularProgress color="primary" />
    </Backdrop>,
    document.body
  );
}
