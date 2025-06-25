import Box from "@mui/material/Box";

export function HidePost({ setHidePost }) {
  return (
    <Box className="w-full p-5 border-b border-x !bg-black !text-white flex gap-2 items-center justify-center">
      <p>You’ve hidden this post.</p>
      <button
        className="bg-gray-700 p-2 place-self-end cursor-pointer"
        onClick={() => setHidePost(false)}
      >
        Undo
      </button>
    </Box>
  );
}
