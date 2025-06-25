import { useState } from "react";
import CreatePost from "./components/CreatePost";
import Posts from "./components/Posts";
import Box from "@mui/material/Box";
import { Outlet } from "react-router-dom";

function Home() {
  const [selected, setSelected] = useState(1);

  return (
    <Box className="w-full lg:w-1/2 text-white border-r">
      <Box className="w-full flex flex-col">
        <Box className="w-full flex justify-around border-y">
          <button
            className="p-3 w-full border-r cursor-pointer"
            onClick={() => setSelected(1)}
          >
            <span
              className={`${
                selected === 1 ? "border-b-2 border-blue-800 px-2" : ""
              }`}
            >
              For You
            </span>
          </button>
          <button
            className="p-3 w-full cursor-pointer"
            onClick={() => setSelected(2)}
          >
            <span
              className={`${
                selected === 2 ? "border-b-2 border-blue-800 px-2" : ""
              }`}
            >
              Following
            </span>
          </button>
        </Box>
        <CreatePost />
        <Posts />
      </Box>
      <Outlet />
    </Box>
  );
}
export default Home;
