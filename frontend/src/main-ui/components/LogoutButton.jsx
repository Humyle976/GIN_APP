import { CgLogOff } from "react-icons/cg";
import { useNavigate } from "react-router-dom";
import { useMutation, useQueryClient } from "@tanstack/react-query";

import axios from "axios";
import { GlobalLoader } from "../../globals/GlobalLoader";

async function logoutFn() {
  return axios.post("http://localhost:8000/auth/logout", "", {
    withCredentials: true,
  });
}

const LogoutButton = ({ className }) => {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { mutate, isPending, isSuccess } = useMutation({
    mutationFn: logoutFn,
    onSuccess: () => {
      queryClient.invalidateQueries(["posts"]);
      navigate("/login");
    },
    onError: (err) => {
      console.log(err);
      console.error("Logout failed:", err.message);
    },
  });

  return (
    <>
      <button
        className={`flex self-start md:self-center bg-red-500 p-2 md:p-3 ml-2 mb-10 text-white text-xl uppercase rounded-lg cursor-pointer ${className}`}
        onClick={() => mutate()}
        disabled={isPending}
      >
        <CgLogOff className="block md:hidden" />
        <p className="hidden md:block">Log Out</p>
      </button>
      {(isPending || isSuccess) && <GlobalLoader />}
    </>
  );
};

export default LogoutButton;
