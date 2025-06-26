import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { countries } from "../../../data/countriesList";
import { Link, useNavigate } from "react-router-dom";
import { useState } from "react";
import { signupFn } from "../../../services/signupFn";

import axios from "axios";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import FormError from "../../../globals/FormError";
import { GlobalLoader } from "../../../globals/GlobalLoader";

export default function RegistrationForm() {
  const navigate = useNavigate();
  const [showPassword, setShowPassword] = useState(false);

  const { mutate, isPending } = useMutation({
    mutationFn: signupFn,
    onSuccess: (data) => {
      if (data.status === 200) {
        navigate("/verify");
      }
    },
    onError: (error) => console.log(error),
  });
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({ mode: "onBlur" });
  const alphaPattern = /^[A-Za-z\s]+$/;
  const emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  const passwordPattern = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[A-Za-z\d]{8,}$/;

  function validateDateOfBirth(value) {
    const today = new Date();
    const dob = new Date(value);
    const age = today.getFullYear() - dob.getFullYear();
    if (dob > today) return "Date of birth cannot be in the future.";
    if (age < 16) return "You must be at least 16 years old.";
    return true;
  }

  async function checkEmailExists(value) {
    try {
      const res = await axios.get(
        `http://localhost:8000/auth/email-exists?email=${value}`
      );
      if (res.status === 200) return "Email already exists";
      else if (res.status === 400) return "Invalid email";
      else return "Unknown error occured";
    } catch (err) {
      if (err.response?.status === 404) return true;
    }
  }

  return (
    <form onSubmit={handleSubmit(mutate)}>
      {isPending && <GlobalLoader />}
      <div className="flex flex-col gap-5">
        <div className="flex gap-3 w-full">
          <div className="flex flex-col w-1/2">
            <input
              autoComplete="firstName"
              type="text"
              className="bg-gray-900 p-4 !text-white focus:border-blue-500 focus:outline focus:border-b"
              placeholder="First Name"
              {...register("firstName", {
                required: "First name is required",
                pattern: {
                  value: alphaPattern,
                  message: "First name must contain only alphabets.",
                },
              })}
            />
            {errors.firstName && (
              <FormError>{errors.firstName.message}</FormError>
            )}
          </div>
          <div className="flex flex-col w-1/2">
            <input
              type="text"
              className="bg-gray-900 p-4 !text-white focus:border-blue-500 focus:outline focus:border-b"
              placeholder="Last Name"
              {...register("lastName", {
                required: "Last name is required",
                pattern: {
                  value: alphaPattern,
                  message: "Last name must contain only alphabets.",
                },
              })}
            />
            {errors.lastName && (
              <FormError>{errors.lastName.message}</FormError>
            )}
          </div>
        </div>
        <div className="flex flex-col gap-2 items-start">
          <label className="text-md text-red-500">*Date Of Birth</label>
          <input
            autoComplete="dateOfBirth"
            type="date"
            className="bg-gray-900 p-4 !text-white w-full focus:border-blue-500 focus:outline focus:border-b"
            {...register("dateOfBirth", {
              required: "Date of birth is required",
              validate: validateDateOfBirth,
            })}
          />
          {errors.dateOfBirth && (
            <FormError>{errors.dateOfBirth.message}</FormError>
          )}
        </div>
        <div className="flex flex-col gap-2 items-start">
          <label className="text-md text-red-500">*Gender</label>
          <div className="flex gap-3 w-full">
            <div className="flex gap-5 items-center bg-gray-900 p-4 !text-white w-1/2 lg:w-1/3 focus:border-blue-500 focus:outline focus:border-b">
              <input
                defaultChecked="true"
                type="radio"
                id="male"
                value="Male"
                {...register("gender", {
                  required: "Gender is required",
                })}
              ></input>
              <label
                htmlFor="male"
                className="text-xl font-medium tracking-wider"
              >
                Male
              </label>
            </div>
            <div className="flex gap-5 items-center bg-gray-900 p-4 !text-white w-1/2 lg:w-1/3 focus:border-blue-500 focus:outline focus:border-b">
              <input
                type="radio"
                id="female"
                value="Female"
                {...register("gender", {
                  required: "Gender is required",
                })}
              ></input>
              <label
                htmlFor="female"
                className="text-xl font-medium tracking-wider"
              >
                Female
              </label>
            </div>
          </div>
          {errors.gender && <FormError>{errors.gender.message}</FormError>}
        </div>

        <div>
          <select
            {...register("country", { required: "Select your country" })}
            className="bg-gray-900 p-4 !text-white w-full focus:border-blue-500 focus:outline focus:border-b"
          >
            <option value="" disabled>
              Country
            </option>
            {countries.map((country) => (
              <option key={country.code} value={country.code}>
                {country.name}
              </option>
            ))}
          </select>
          {errors.country && <FormError>{errors.country.message}</FormError>}
        </div>
        <div>
          <input
            type="email"
            className="bg-gray-900 p-4 !text-white w-full focus:border-blue-500 focus:outline focus:border-b"
            placeholder="Email"
            {...register("email", {
              required: "Email is required",
              pattern: {
                value: emailPattern,
                message: "Please enter a valid email.",
              },
              validate: checkEmailExists,
            })}
          />
          {errors.email && <FormError>{errors.email.message}</FormError>}
        </div>
        <div className="flex flex-col">
          <div className="relative">
            <input
              type={showPassword ? "text" : "password"}
              className="bg-gray-900 p-4 !text-white w-full focus:border-blue-500 focus:outline focus:border-b"
              placeholder="Password"
              {...register("password", {
                required: "Password is required",
                pattern: {
                  value: passwordPattern,
                  message:
                    "Password must be alphanumeric, contain at least 1 uppercase letter, 1 lowercase letter, and 1 number.",
                },
                minLength: {
                  value: 8,
                  message: "Password must be at least 8 characters long.",
                },
              })}
            />
            <span
              onClick={() => setShowPassword((prev) => !prev)}
              className="absolute right-3 top-1/2 transform -translate-y-1/2 cursor-pointer text-xl"
            >
              {showPassword ? <VisibilityIcon /> : <VisibilityOffIcon />}
            </span>
          </div>
          {errors.password && <FormError>{errors.password.message}</FormError>}
        </div>
        <div className="flex flex-col items-center gap-3 mt-5 text-lg">
          <button
            type="submit"
            className="w-full md:w-1/2 rounded-xl px-4 py-3 mt-5 font-semibold text-lg transition-all duration-300 
              bg-gradient-to-r from-purple-700 to-pink-600 text-white shadow-md hover:shadow-lg cursor-pointer"
          >
            Sign Up
          </button>
          <p>
            Already have an account?{" "}
            <Link to="/login">
              <span className="text-blue-400">Login</span>
            </Link>
          </p>
        </div>
      </div>
    </form>
  );
}
