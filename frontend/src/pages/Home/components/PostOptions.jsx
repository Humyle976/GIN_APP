import { useMutation, useQueryClient } from "@tanstack/react-query";
import { deletePost } from "../../../services/deletePost";
import { GlobalLoader } from "../../../globals/GlobalLoader";
export function PostOptions({ IsOwner, handleHidePost, Fullname, ID }) {
  const queryClient = useQueryClient();
  const { mutate, isPending } = useMutation({
    mutationFn: deletePost,
    onSuccess: () => {
      queryClient.invalidateQueries(["posts"]);
    },
  });
  return (
    <>
      {IsOwner ? (
        <>
          {isPending && <GlobalLoader />}
          <button
            onClick={handleHidePost}
            className="p-2 w-full border-y border-x text-start cursor-pointer hover:bg-gray-200"
          >
            Hide Post
          </button>
          <button
            className="p-2 w-full border-x text-start cursor-pointer hover:bg-gray-200"
            onClick={() => mutate(ID)}
          >
            Delete Post
          </button>
          <button className="p-2 w-full border-y border-x text-start cursor-pointer hover:bg-gray-200">
            Edit Post
          </button>
        </>
      ) : (
        <>
          {isPending && <GlobalLoader />}
          <button
            onClick={handleHidePost}
            className="p-2 w-full border-y border-x text-start cursor-pointer hover:bg-gray-200"
          >
            Hide Post
          </button>
          <button className="p-2 w-full border-x text-start cursor-pointer hover:bg-gray-200">
            Report
          </button>
          <button className="p-2 w-full border-y border-x text-start cursor-pointer hover:bg-gray-200">
            Block {Fullname}
          </button>
        </>
      )}
    </>
  );
}
