import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import { Link } from "react-router-dom";
import { loginFn } from "../../../services/loginFn";
import { GlobalLoader } from "../../../globals/GlobalLoader";

import Input from "@mui/material/Input";
import Box from "@mui/material/Box";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import ErrorPage from "../../../globals/ErrorPage";
import FormError from "../../../globals/FormError";

export function LoginForm() {
  const [loginField, setLoginField] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);

  const [pageError, setPageError] = useState(null);
  const [inlineError, setInlineError] = useState(null);
  const navigate = useNavigate();

  const { mutate, isPending, isSuccess } = useMutation({
    mutationFn: loginFn,
    onSuccess: () => {
      setInlineError(null);
      setPageError(null);
      navigate("/home");
    },
    onError: (error) => {
      if (error.response) {
        const { status } = error.response;

        if (status === 400 || status === 401) {
          setInlineError("Wrong Credentials");
        } else {
          setPageError({ status, message: error.message });
        }
      } else {
        setPageError({ message: "Network or server error" });
      }
    },
  });

  if (pageError) {
    return (
      <ErrorPage
        errorStatus={pageError.status}
        errorMessage={pageError.message}
      />
    );
  }
  return (
    <Box>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          mutate({ loginField, password });
        }}
      >
        <Box className="flex flex-col gap-5 max-h-full">
          <Input
            type="email"
            name="LoginField"
            placeholder="Email"
            value={loginField}
            onChange={(e) => setLoginField(e.target.value)}
            className="bg-gray-800 !text-white p-3"
            required
            fullWidth
          />
          <Input
            name="Password"
            type={showPassword ? "text" : "password"}
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="bg-gray-800 !text-white p-3"
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
          {inlineError && (
            <Box className="self-center">
              <FormError size="text-xl">{inlineError}</FormError>
            </Box>
          )}
        </Box>
        <Box className="flex flex-col items-center gap-5 mt-3">
          <button
            type="submit"
            className="w-full md:w-2/3 cursor-pointer rounded-xl px-4 py-3 mt-5 font-semibold text-lg transition-all duration-300 bg-gradient-to-r from-purple-700 to-pink-600 text-white shadow-md hover:shadow-lg"
          >
            Login
          </button>
          <p>
            Don't have an account?{" "}
            <Link to="/signup">
              <button className="text-blue-500 cursor-pointer">Sign Up</button>
            </Link>
          </p>
        </Box>
      </form>
      {(isPending || isSuccess) && <GlobalLoader />}
    </Box>
  );
}
