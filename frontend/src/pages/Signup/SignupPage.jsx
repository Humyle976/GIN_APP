import Box from "@mui/material/Box";
import RegistrationForm from "./components/SignUpForm";

function Signup() {
  return (
    <Box className="flex justify-center items-start pt-10 h-screen bg-gradient-to-br from-gray-950 via-black to-purple-950 overflow-y-auto">
      <Box className="flex flex-col gap-2 p-6 w-11/12 sm:w-3/5 md:w-3/5 lg:w-2/5 my-10 text-white rounded-3xl bg-opacity-20 bg-black backdrop-blur-md shadow-[0_0_40px_rgba(124,58,237,0.5)] border border-purple-700">
        <Box className="flex justify-center mb-2">
          <h1 className="text-xl md:text-3xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-purple-400 to-pink-500">
            Create A New Account
          </h1>
        </Box>
        <RegistrationForm />
      </Box>
    </Box>
  );
}

export default Signup;
