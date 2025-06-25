import axios from "axios";

export async function loginFn({ loginField, password }) {
  return axios.post(
    "http://localhost:8000/auth/login",
    { loginField, password },
    { withCredentials: true }
  );
}
