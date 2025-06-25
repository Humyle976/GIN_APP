import axios from "axios";
export async function createPost({ visibility, postText, postFile }) {
  const formData = new FormData();
  formData.append("visibility", visibility);
  formData.append("text", postText);
  if (postFile) {
    formData.append("file", postFile);
  }

  const res = await axios.post("http://localhost:8000/posts/", formData, {
    withCredentials: true,
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return res.data;
}
