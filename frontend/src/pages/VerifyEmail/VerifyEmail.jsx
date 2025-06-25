import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import ErrorPage from "../../globals/ErrorPage";
import { Box, Button, Typography } from "@mui/material";
import { useState, useEffect, useRef } from "react";

async function checkToken() {
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
function VerifyEmail() {
  const navigate = useNavigate();
  const { isError, error, isLoading } = useQuery({
    queryKey: ["verify-email"],
    queryFn: checkToken,
    retry: false,
    staleTime: Infinity,
  });

  const [code, setCode] = useState(Array(6).fill(""));
  const [codeError, setCodeError] = useState("");
  const [postError, setPostError] = useState(null);

  const inputRefs = useRef([]);
  const buttonRef = useRef();
  useEffect(() => {
    const timer = setTimeout(() => {
      inputRefs.current[0]?.focus();
    }, 0);

    return () => clearTimeout(timer);
  }, []);

  function handleChange(index, e) {
    const value = e.target.value;
    if (isNaN(value)) return;
    setCodeError("");
    const newCode = [...code];
    newCode[index] = value;
    setCode(newCode);

    if (value && index < 5) {
      inputRefs.current[index + 1].focus();
    }
  }

  function handleKeyDown(index, e) {
    if (e.key === "Backspace" || e.key === "Delete") {
      if (code[index] === "") {
        if (index > 0) {
          inputRefs.current[index - 1].focus();
        }
      } else {
        const newCode = [...code];
        newCode[index] = "";
        setCode(newCode);
      }
    }
  }

  function handlePaste(e) {
    e.preventDefault();
    const pastedText = e.clipboardData.getData("text/plain").slice(0, 6);
    if (!/^\d+$/.test(pastedText)) return;

    const newCode = pastedText.split("").slice(0, 6);
    setCode(newCode.concat(Array(6 - newCode.length).fill("")));

    const lastFilledIndex = newCode.length - 1;
    if (lastFilledIndex > -1 && lastFilledIndex < 6) {
      inputRefs.current[lastFilledIndex].focus();
    }
  }

  async function handleSubmit() {
    const codeInt = parseInt(code.join(""), 10);
    buttonRef.current.disabled = true;
    try {
      const res = await axios.post(
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
      setPostError(null);
      if (res.data.status === 201) {
        navigate("/home", { replace: true });
      }
    } catch (err) {
      const status = err.response?.data.status;
      const message = err.response?.data.message;
      if (status === 400) {
        setCodeError(message);
      } else if (status === 404 || status === 500) {
        setPostError({ status, message });
      } else {
        setPostError("Unknown Error Occured!");
      }
    }
    buttonRef.current.disabled = false;
  }

  if (isLoading) return <div></div>;
  if (postError) {
    return (
      <ErrorPage
        errorMessage={postError.message}
        errorStatus={postError.status}
      />
    );
  }
  if (isError) {
    return (
      <ErrorPage errorMessage={error.message} errorStatus={error.status} />
    );
  }

  return (
    <Box className="flex flex-col items-center justify-center min-h-screen bg-gradient-to-br from-gray-950 via-black to-purple-950 p-6">
      <Typography
        sx={{
          color: "white",
          fontSize: "36px",
          textAlign: "center",
          fontWeight: 600,
        }}
      >
        Check Your Inbox
      </Typography>

      <Box className="mt-10 flex flex-col gap-5 bg-white/10 text-white p-6 rounded-2xl w-full max-w-md shadow-[0_0_30px_rgba(124,58,237,0.5)] border border-purple-600 backdrop-blur-lg">
        <Typography className="text-base md:text-lg text-white text-center font-medium">
          Enter the 6-digit code we sent to your email.
        </Typography>
        <Box className="flex flex-col gap-3">
          <div className="flex justify-center gap-3">
            {code.map((digit, index) => (
              <input
                key={index}
                type="text"
                maxLength="1"
                value={digit}
                onChange={(e) => handleChange(index, e)}
                onKeyDown={(e) => handleKeyDown(index, e)}
                onPaste={handlePaste}
                ref={(el) => (inputRefs.current[index] = el)}
                className="w-14 h-14 text-center text-2xl text-black bg-white border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-500 shadow-md transition-all"
              />
            ))}
          </div>

          {codeError && (
            <div className="text-red-500 text-center font-medium">
              {codeError}
            </div>
          )}
        </Box>
        <Button
          variant="contained"
          className="!mt-4 !bg-gradient-to-r !from-purple-700 !to-pink-600 hover:!from-pink-600 hover:!to-purple-700 !text-white !font-semibold !py-3 !rounded-xl !text-base !transition-all"
          onClick={handleSubmit}
          ref={buttonRef}
        >
          Verify
        </Button>
      </Box>
    </Box>
  );
}

export default VerifyEmail;
