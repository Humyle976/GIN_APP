import { useRef, useState } from "react";
import { MenuItem, Select } from "@mui/material";
import { IoMdAttach } from "react-icons/io";
import { getMediaType } from "../../../services/getMediaType";
import { IoCloseSharp } from "react-icons/io5";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createPost } from "../../../services/createPost";
import { Toaster } from "react-hot-toast";
import toast from "react-hot-toast";

import Avatar from "@mui/material/Avatar";
import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import { GlobalLoader } from "../../../globals/GlobalLoader";

function CreatePost() {
  const [postText, setPostText] = useState("");
  const [postFile, setPostFile] = useState(null);
  const [filePreview, setFilePreview] = useState(null);
  const [visibility, setVisibility] = useState("public");
  const inputRef = useRef();

  const queryClient = useQueryClient();

  const { mutate, isPending } = useMutation({
    mutationFn: createPost,
    onSuccess: (data) => {
      queryClient.setQueryData(["posts"], (oldPost = []) => {
        if (!oldPost) return [data.post];
        else return [data.post, ...oldPost];
      });
      setPostText("");
      setPostFile(null);
      setFilePreview(null);
      setVisibility("public");
    },
    onError: (error) => {
      console.log(error);
    },
  });
  function onChooseImg() {
    inputRef.current.click();
  }

  function handleFileChange(e) {
    const file = e.target.files[0];

    if (file) {
      const type = getMediaType(file?.name);
      if (type === "unknown") {
        toast.error("Invalid File Type");
        return;
      }
      setPostFile(file);
      setFilePreview(URL.createObjectURL(file));
    }
  }

  function closePreview() {
    setPostFile(null);
    setFilePreview(null);
  }

  return (
    <Box className="w-full flex gap-3 p-2 border-b bg-black">
      <Toaster position="top-center" reverseOrder={false} />
      <Avatar className="cursor-pointer" />
      <Box className="flex flex-col gap-2 w-full bg-black">
        <form
          onSubmit={(e) => {
            e.preventDefault();
            mutate({ visibility, postText, postFile });
          }}
        >
          <Box>
            <TextField
              id="multiline-flexible"
              multiline
              maxRows={4}
              value={postText}
              className="w-full bg-white rounded-md"
              onChange={(e) => setPostText(e.target.value)}
              placeholder="What's on your mind?"
            />
          </Box>

          {filePreview && (
            <Box className="mt-2">
              {getMediaType(postFile?.name || "") === "image" ? (
                <Box className="relative w-full max-w-xs">
                  <img
                    src={filePreview}
                    alt="preview"
                    className="w-full max-h-60 rounded-lg"
                  />
                  <button
                    type="button"
                    onClick={closePreview}
                    className="absolute top-2 right-2 bg-white rounded-full p-1 text-gray-800 shadow-md cursor-pointer"
                  >
                    <IoCloseSharp />
                  </button>
                </Box>
              ) : (
                <Box className="relative w-full max-w-md">
                  <video
                    src={filePreview}
                    controls
                    className="w-full max-h-60 rounded-lg"
                  />
                  <button
                    type="button"
                    onClick={closePreview}
                    className="absolute top-2 right-2 bg-white rounded-full p-1 text-gray-800 shadow-md cursor-pointer"
                  >
                    <IoCloseSharp />
                  </button>
                </Box>
              )}
            </Box>
          )}

          <Box className="flex justify-between items-center mt-5 mb-2">
            <Box className="flex gap-2 items-center">
              <Select
                name="visibility"
                className="w-30 h-10 !text-white border-1 border-white"
                value={visibility}
                onChange={(e) => setVisibility(e.target.value)}
              >
                <MenuItem value="public">Public</MenuItem>
                <MenuItem value="private">Friend's Only</MenuItem>
              </Select>

              <input
                name="file"
                type="file"
                className="hidden"
                accept="image/*, video/*"
                ref={inputRef}
                onChange={handleFileChange}
              />
              <IoMdAttach
                className="text-xl cursor-pointer"
                onClick={onChooseImg}
              />
            </Box>

            <button
              type="submit"
              disabled={postText.length === 0}
              className={`${
                postText.length > 0
                  ? "bg-red-500 cursor-pointer"
                  : "bg-gray-500"
              } p-2 rounded-md text-white`}
            >
              Post
            </button>
          </Box>
        </form>
      </Box>
      {isPending && <GlobalLoader />}
    </Box>
  );
}

export default CreatePost;
