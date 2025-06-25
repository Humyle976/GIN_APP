import axios from "axios";

export async function getFeed() {
  try {
    const res = await axios.get("http://localhost:8000/feed", {
      withCredentials: true,
    });
    return res.data.data;
  } catch (err) {
    throw new Error(err);
  }
}
