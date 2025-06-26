import { useMutation, useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import ErrorPage from "../../globals/ErrorPage";
import { Box, Button, Typography } from "@mui/material";
import { useState, useRef } from "react";
import { verifyCode } from "../../services/verifyCode";
import { GlobalLoader } from "../../globals/GlobalLoader";
import { checkTokenExists } from "../../services/checkTokenExists";
import { resendCodeFn } from "../../services/resendCodeFn";

function VerifyEmail() {
  const [code, setCode] = useState(Array(6).fill(""));
  const [codeError, setCodeError] = useState("");
  const [postError, setPostError] = useState(null);

  const [cooldown, setCooldown] = useState(0);
  const intervalRef = useRef(null);

  const navigate = useNavigate();
  const inputRefs = useRef([]);
  const verifyButtonRef = useRef();

  const {
    isError,
    error,
    isPending: isPendingQuery,
  } = useQuery({
    queryKey: ["verify-email"],
    queryFn: checkTokenExists,
    retry: false,
    staleTime: Infinity,
  });

  const { mutate, isPending } = useMutation({
    mutationFn: verifyCode,
    onMutate: () => {
      verifyButtonRef.current.disabled = true;
    },
    onSuccess: (data) => {
      setPostError(null);
      if (data.status === 201) {
        navigate("/home", { replace: true });
      }
    },
    onError: (err) => {
      const status = err.response?.data.status;
      const message = err.response?.data.message;
      if (status === 400) {
        setCodeError(message);
      } else if (status === 404 || status === 500) {
        setPostError({ status, message });
      } else {
        setPostError("Unknown Error Occured!");
      }
    },
    onSettled: () => {
      verifyButtonRef.current.disabled = false;
    },
  });

  const { mutate: mutateResend, isPending: isPendingReset } = useMutation({
    mutationFn: resendCodeFn,
    onSuccess: () => {
      setCooldown(60);
      intervalRef.current = setInterval(() => {
        setCooldown((prev) => {
          if (prev <= 1) {
            clearInterval(intervalRef.current);
            return 0;
          }
          return prev - 1;
        });
      }, 1000);
    },
    onError: (err) => {
      const status = err.response?.data.status;
      const message = err.response?.data.message;
      if (status === 404 || status === 500) {
        setPostError({ status, message });
      } else if (status === 425) {
        setCodeError(message);
      } else {
        setPostError("Unknown Error Occured!");
      }
    },
  });

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

  if (isPendingQuery) return <div></div>;
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
          <div className="flex flex-col gap-5">
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

            <div className="flex w-full">
              {codeError && (
                <div className="w-1/2 text-red-500 text-md font-medium">
                  {codeError}
                </div>
              )}
              <div className="flex w-full justify-end">
                <button
                  className={`text-cyan-300 ${
                    cooldown > 0
                      ? "opacity-50 cursor-not-allowed"
                      : "cursor-pointer"
                  }`}
                  onClick={mutateResend}
                  disabled={cooldown > 0}
                >
                  {cooldown > 0 ? `Resend again in ${cooldown}s` : "Resend?"}
                </button>
              </div>
            </div>
          </div>
        </Box>
        <Button
          variant="contained"
          className="!mt-4 !bg-gradient-to-r !from-purple-700 !to-pink-600 hover:!from-pink-600 hover:!to-purple-700 !text-white !font-semibold !py-3 !rounded-xl !text-base !transition-all"
          onClick={() => mutate(code)}
          ref={verifyButtonRef}
        >
          Verify
        </Button>
        {(isPending || isPendingReset) && <GlobalLoader />}
      </Box>
    </Box>
  );
}

export default VerifyEmail;
