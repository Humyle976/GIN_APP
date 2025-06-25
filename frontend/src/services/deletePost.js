import axios from "axios";

export async function deletePost(id) {
  return await axios.delete(`http://localhost:8000/posts/${id}`, {
    withCredentials: true,
  });
}
