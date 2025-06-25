import { useState } from "react";
import { getMediaType } from "../../../services/getMediaType";
import { Popper } from "@mui/material";
import { PostModal } from "./postModal";
import { PostOptions } from "./PostOptions";
import Card from "@mui/material/Card";
import Box from "@mui/material/Box";
import Avatar from "@mui/material/Avatar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import CardContent from "@mui/material/CardContent";
import FavoriteIcon from "@mui/icons-material/Favorite";
import ModeCommentIcon from "@mui/icons-material/ModeComment";
import { HidePost } from "./HidePost";

export default function PostCard({ post }) {
  const [anchorEl, setAnchorEl] = useState(null);
  const [hidePost, setHidePost] = useState(false);
  const [showModal, setShowModal] = useState(false);

  const handleClick = (event) => {
    setAnchorEl(anchorEl ? null : event.currentTarget);
  };

  const openAnchor = Boolean(anchorEl);

  function handleHidePost() {
    setHidePost(true);
    setAnchorEl(false);
  }

  return hidePost ? (
    <HidePost setHidePost={setHidePost} />
  ) : (
    <Card
      className={`block w-full p-3 border-b border-x !bg-black !text-white`}
    >
      <Box className="w-full flex items-center justify-between">
        <Box className="flex gap-3">
          <Avatar aria-label="profile-picture" className="cursor-pointer" />
          <Box>
            <Typography className="cursor-pointer">
              {post.Fullname} {post.IsOwner ? " (You)" : ""}
            </Typography>
            <Typography className="text-sm text-gray-500">
              {getTime(post.CreatedAt)}
            </Typography>
          </Box>
        </Box>
        <IconButton className="!text-white" onClick={handleClick}>
          <MoreVertIcon />
        </IconButton>
        <Popper
          id="simple-popper"
          open={openAnchor}
          anchorEl={anchorEl}
          placement="bottom-end"
          className="bg-white flex flex-col items-start text-black w-50"
        >
          <PostOptions
            IsOwner={post.IsOwner}
            Fullname={post.Fullname}
            handleHidePost={handleHidePost}
            ID={post.PostID}
          />
        </Popper>
      </Box>

      <Box className="mt-1 border-1 border-blue-950/90 rounded-md text-white">
        <CardContent className="">
          <Typography className="mb-2">{post.Content}</Typography>

          {post.FileURL && (
            <Box component="div" className="cursor-pointer">
              {getMediaType(post.FileURL) === "image" ? (
                <img
                  src={`http://localhost:8000${post.FileURL}`}
                  alt="Post media"
                  className="max-h-125 w-full rounded-md"
                  onClick={() => setShowModal(true)}
                />
              ) : (
                <video
                  preload="auto"
                  controls
                  className="w-full max-h-100"
                  src={`http://localhost:8000${post.FileURL}`}
                />
              )}
            </Box>
          )}
        </CardContent>

        <Box className="flex gap-4 px-3 pb-2">
          <IconButton className="flex gap-1 items-center">
            <FavoriteIcon className="text-white" />
            <Typography className="!text-white">{post.Likes}</Typography>
          </IconButton>

          <IconButton className="flex gap-1 items-center">
            <ModeCommentIcon className="text-blue-500" />
            <Typography className="!text-white">{post.Comments}</Typography>
          </IconButton>
        </Box>
        <PostModal
          url={post.FileURL}
          showModal={showModal}
          setShowModal={setShowModal}
        />
      </Box>
    </Card>
  );
}

function getTime(date) {
  const currTime = new Date().getTime() / 1000;
  const postTime = new Date(date).getTime() / 1000;
  const diff = currTime - postTime;

  const oneMinute = 60;
  const oneHour = 3600;
  const oneDay = 86400;
  const oneWeek = 7 * oneDay;

  if (diff < oneMinute) {
    return `${Math.floor(diff)} secs ago`;
  } else if (diff < oneHour) {
    return `${Math.floor(diff / oneMinute)} mins ago`;
  } else if (diff < oneDay) {
    return `${Math.floor(diff / oneHour)} hrs ago`;
  } else if (diff < oneWeek) {
    return `${Math.floor(diff / oneDay)} days ago`;
  } else {
    const d = new Date(date);
    return d.toLocaleDateString("en-US", {
      day: "numeric",
      month: "short",
      year: "numeric",
    });
  }
}
