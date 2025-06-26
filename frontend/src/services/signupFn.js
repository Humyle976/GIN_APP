import axios from "axios";

export async function signupFn(data) {
  if (
    data.firstName == "" ||
    data.lastName == "" ||
    data.dateOfBirth == "" ||
    data.gender == "" ||
    data.country == "" ||
    data.email == "" ||
    data.password == ""
  )
    data.firstName = data.firstName.trim();
  data.lastName = data.lastName.trim();
  data.dateOfBirth == data.dateOfBirth.trim();
  data.gender == data.gender.trim();
  data.country == data.country.trim();
  data.email == data.email.trim();
  data.password == data.password.trim();
  return axios.post(
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
}
