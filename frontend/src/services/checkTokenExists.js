import axios from "axios";
export async function checkTokenExists() {
  try {
    const res = await axios.get("http://localhost:8000/auth/verify", {
      withCredentials: true,
    });
    return res.data;
  } catch (err) {
    const status = err.response?.status;
    let message = "Unexpected error occurred.";

    if (status === 404 || status === 500 || status === 400) {
      message = err.response.data.message;
    }

    const error = new Error(message);
    error.status = status ?? "";
    throw error;
  }
}
