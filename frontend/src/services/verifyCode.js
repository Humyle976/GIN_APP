import axios from "axios";
export async function verifyCode(code) {
  const codeInt = parseInt(code.join(""), 10);
  return await axios.post(
    "http://localhost:8000/auth/verify",
    {
      code: codeInt,
    },
    {
      headers: {
        "Content-Type": "application/json",
      },
      withCredentials: true,
    }
  );
}
