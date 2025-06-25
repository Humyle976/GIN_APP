import { useQuery, useQueryClient } from "@tanstack/react-query";
import { getFeed } from "../../../services/apiFeed";
import { useState } from "react";
import Typography from "@mui/material/Typography";
import CircularProgress from "@mui/material/CircularProgress";
import PostCard from "./PostCard";
import Box from "@mui/material/Box";
import ErrorPage from "../../../globals/ErrorPage";

function Posts() {
  const queryClient = useQueryClient();
  const [refetching, setRefetching] = useState(false);

  const { data, isLoading, isError, error } = useQuery({
    queryKey: ["posts"],
    queryFn: getFeed,
    staleTime: 1000 * 60 * 60 * 60 * 60 * 60,
    retry: false,
  });

  async function refetchPosts() {
    setRefetching(true);
    await queryClient.invalidateQueries(["posts"]);
    setRefetching(false);
  }
  if (isLoading || refetching)
    return (
      <Box className="flex items-center justify-center h-screen w-full">
        <CircularProgress className="text-4xl" />
      </Box>
    );
  if (isError) return <ErrorPage errorMessage={error.message} />;

  if (!data)
    return (
      <Box className="flex flex-col gap-3 items-center justify-center h-screen w-full">
        <Typography variant="h4">Couldn't find posts to show</Typography>
        <button
          className="p-3 bg-blue-700 rounded-lg text-white cursor-pointer font-bold"
          onClick={refetchPosts}
        >
          Reload
        </button>
      </Box>
    );
  return (
    <Box className="flex flex-col">
      {data.map((post) => (
        <PostCard key={post.PostID} post={post} />
      ))}
    </Box>
  );
}

export default Posts;
