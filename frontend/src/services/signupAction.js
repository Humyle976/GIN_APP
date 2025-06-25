import { redirect } from "react-router-dom";
import axios from "axios";

export async function signupAction({ request }) {
  const formData = await request.formData();
  const data = Object.fromEntries(formData);

  if (
    data.firstName == "" ||
    data.lastName == "" ||
    data.dateOfBirth == "" ||
    data.gender == "" ||
    data.country == "" ||
    data.email == "" ||
    data.password == ""
  )
    return new Error("Please Fill out whole form");

  if (data.password.length < 6)
    return new Error("Password must have atleast 6 characters");

  data.firstName = data.firstName.trim();
  data.lastName = data.lastName.trim();
  data.dateOfBirth == data.dateOfBirth.trim();
  data.gender == data.gender.trim();
  data.country == data.country.trim();
  data.email == data.email.trim();
  data.password == data.password.trim();
  try {
    const res = await axios.post(
      "http://localhost:8000/auth/signup",
      {
        first_name: data.firstName,
        last_name: data.lastName,
        dob: data.dateOfBirth,
        gender: data.gender,
        country: data.country,
        email: data.email,
        password: data.password,
      },
      { withCredentials: true }
    );

    if (res.data.status === 200) {
      return redirect("/verify");
    }
  } catch (err) {
    console.log(err);
  }
}
