import { useReducer, useState } from "react";
import { Input, Box, FormLabel, Radio, Select, MenuItem } from "@mui/material";
import { Form, Link } from "react-router-dom";
import { countries } from "../../../data/countriesList";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import FormError from "../../../globals/FormError";
import axios from "axios";
import {
  initialState,
  Signupreducer,
} from "../../../reducers/signUpFormReducer";

export default function RegistrationForm() {
  const [state, dispatch] = useReducer(Signupreducer, initialState);
  const [showPassword, setShowPassword] = useState(false);

  const {
    firstName,
    lastName,
    gender,
    country,
    email,
    password,
    date,
    errors,
  } = state;

  const setField = (field, value) =>
    dispatch({ type: "SET_FIELD", field, value });

  const setError = (field, value) =>
    dispatch({ type: "SET_ERROR", field, value });

  const validateName = (name, field) => {
    setError(field, !name);
  };

  const validateDate = (inputDate) => {
    const now = new Date();
    const sixteenYearsAgo = new Date(
      now.getFullYear() - 16,
      now.getMonth(),
      now.getDate()
    );
    const dob = new Date(inputDate);

    const isValid = dob <= sixteenYearsAgo && dob <= new Date();
    setError("date", !isValid);
  };

  const validateEmail = async () => {
    if (!email) return setError("email", "Email is required");

    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!regex.test(email)) return setError("email", "Invalid email format");

    try {
      const res = await axios.get("http://localhost:8000/auth/email-exists", {
        params: { email },
      });
      setError("email", res.data.message || "");
    } catch {
      setError("email", "");
    }
  };

  const validatePassword = () => {
    const regex = /^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?!.* ).{7,}$/;
    setError("password", !regex.test(password));
  };

  const validateCountry = () => {
    setError("country", !country);
  };

  const canRegister =
    firstName &&
    lastName &&
    email &&
    password &&
    !Object.values(errors).some((err) => err);

  return (
    <Form className="p-5 h-full overflow-y-auto" method="POST">
      <Box className="flex flex-col gap-5">
        <Box className="flex gap-2">
          <Input
            name="firstName"
            placeholder="First Name"
            value={firstName}
            onChange={(e) => setField("firstName", e.target.value)}
            onBlur={() => validateName(firstName, "firstName")}
            error={errors.firstName}
            className="bg-gray-800 !text-white p-3 w-1/2"
            required
          />
          <Input
            name="lastName"
            placeholder="Last Name"
            value={lastName}
            onChange={(e) => setField("lastName", e.target.value)}
            onBlur={() => validateName(lastName, "lastName")}
            error={errors.lastName}
            className="bg-gray-800 !text-white p-3 w-1/2"
            required
          />
        </Box>

        <Box className="flex flex-col gap-1">
          <FormLabel className="!text-white">Date of Birth</FormLabel>
          <Input
            name="dateOfBirth"
            type="date"
            value={date}
            onChange={(e) => setField("date", e.target.value)}
            onBlur={() => validateDate(date)}
            error={errors.date}
            className="bg-gray-800 !text-white p-2"
            required
          />
          {errors.date && <FormError>Must be at least 16 years old.</FormError>}
        </Box>

        <Box className="flex flex-col gap-1">
          <Select
            name="country"
            value={country}
            onChange={(e) => setField("country", e.target.value)}
            onBlur={validateCountry}
            displayEmpty
            error={errors.country}
            className="!text-white bg-gray-800"
            required
          >
            <MenuItem disabled value="">
              Country
            </MenuItem>
            {countries.map((c) => (
              <MenuItem key={c.code} value={c.code}>
                {c.name}
              </MenuItem>
            ))}
          </Select>
          {errors.country && <FormError>Country is required</FormError>}
        </Box>

        <Box className="flex flex-col gap-1">
          <FormLabel className="!text-white">Gender</FormLabel>
          <Box className="flex gap-3">
            {["Male", "Female"].map((g) => (
              <Box
                key={g}
                className={`w-1/2 flex items-center gap-2 px-3 py-2 bg-gray-800 ${
                  gender === g
                    ? "border border-purple-500"
                    : "border border-gray-600"
                }`}
              >
                <Radio
                  name="gender"
                  value={g}
                  checked={gender === g}
                  onChange={(e) => setField("gender", e.target.value)}
                  sx={{ color: "white" }}
                />
                <FormLabel className="!text-white">{g}</FormLabel>
              </Box>
            ))}
          </Box>
        </Box>

        <Box>
          <Input
            name="email"
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setField("email", e.target.value)}
            onBlur={validateEmail}
            error={!!errors.email}
            className="bg-gray-800 !text-white p-2"
            required
            fullWidth
          />
          {errors.email && <FormError>{errors.email}</FormError>}
        </Box>

        <Box>
          <Input
            name="password"
            type={showPassword ? "text" : "password"}
            placeholder="Password"
            value={password}
            onChange={(e) => setField("password", e.target.value)}
            onBlur={validatePassword}
            error={errors.password}
            className="bg-gray-800 !text-white p-2"
            endAdornment={
              password &&
              (showPassword ? (
                <VisibilityOffIcon onClick={() => setShowPassword(false)} />
              ) : (
                <VisibilityIcon onClick={() => setShowPassword(true)} />
              ))
            }
            required
            fullWidth
          />
          {errors.password && (
            <FormError>
              Password must include uppercase, lowercase, a number, and be at
              least 7 characters.
            </FormError>
          )}
        </Box>

        <Box className="flex flex-col items-center gap-5 mt-3">
          <button
            disabled={!canRegister}
            type="submit"
            className={`w-full md:w-1/2 rounded-xl px-4 py-3 mt-5 font-semibold text-lg transition-all duration-300 ${
              canRegister
                ? "bg-gradient-to-r from-purple-700 to-pink-600 text-white shadow-md hover:shadow-lg"
                : "bg-gray-800 text-gray-400 cursor-not-allowed"
            }`}
          >
            Register
          </button>
          <p>
            Already have an account?{" "}
            <Link to="/login">
              <button className="text-blue-500 cursor-pointer">Login</button>
            </Link>
          </p>
        </Box>
      </Box>
    </Form>
  );
}
