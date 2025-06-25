import axios from "axios";
import { redirect } from "react-router-dom";

export async function protectedRouteLoader() {
  try {
    const res = await axios.get("http://localhost:8000/auth/login", {
      withCredentials: true,
      validateStatus: () => true,
    });

    if (res.status === 200) {
      return null;
    } else {
      return redirect("/login");
    }
  } catch {
    return redirect("/login");
  }
}

export async function loginPageLoader() {
  try {
    const res = await axios.get("http://localhost:8000/auth/login", {
      withCredentials: true,
    });
    if (res.status === 200) {
      return redirect("/home");
    }
    return null;
  } catch {
    return null;
  }
}
