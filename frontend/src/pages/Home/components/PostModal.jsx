import Modal from "@mui/material/Modal";
import Backdrop from "@mui/material/Backdrop";
import Fade from "@mui/material/Fade";
import Box from "@mui/material/Box";

export function PostModal({ url, showModal, setShowModal }) {
  function handleClose() {
    setShowModal(false);
  }

  return (
    <Modal
      open={showModal}
      onClose={handleClose}
      closeAfterTransition
      slots={{ backdrop: Backdrop }}
      slotProps={{ backdrop: { timeout: 500 } }}
    >
      <Fade in={showModal}>
        <Box
          className="fixed fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 outline-none "
          component="div"
          sx={{
            bgcolor: "background.paper",
            boxShadow: 24,
            p: 1,
            borderRadius: 1,
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <img
            src={`http://localhost:8000${url}`}
            alt="Full view"
            className="max-w-full max-h-[90vh] object-contain rounded"
          />
        </Box>
      </Fade>
    </Modal>
  );
}
