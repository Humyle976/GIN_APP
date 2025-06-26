import axios from "axios";
export async function resendCodeFn() {
  return axios.post("http://localhost:8000/auth/verify/resend", "", {
    withCredentials: true,
  });
}
