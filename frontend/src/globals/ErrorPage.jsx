import { IoMdSad } from "react-icons/io";

function ErrorPage({ errorStatus, errorMessage }) {
  return (
    <div className="w-full h-80 flex flex-col justify-center items-center ">
      <IoMdSad className="text-8xl" />
      <div className="flex gap-2 ">
        {errorStatus && (
          <h1 className="text-2xl font-bold">{errorStatus} Error: </h1>
        )}
        <h1 className="text-2xl">{errorMessage || "Unknown Error Occured"}</h1>
      </div>
    </div>
  );
}
export default ErrorPage;
