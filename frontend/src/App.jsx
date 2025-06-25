import {
  createBrowserRouter,
  RouterProvider,
  Navigate,
} from "react-router-dom";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { loginPageLoader, protectedRouteLoader } from "./services/authLoader";
import VerifyEmail from "./pages/VerifyEmail/VerifyEmail";
import Home from "./pages/Home/Home";
import Login from "./pages/Login/LoginPage";
import Signup from "./pages/Signup/SignupPage";
import AppLayout from "./main-ui/AppLayout";
import ErrorPage from "./globals/ErrorPage";

const queryClient = new QueryClient();

const Router = createBrowserRouter([
  {
    path: "",
    element: <AppLayout />,
    loader: protectedRouteLoader,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <Navigate to="/home" replace />,
      },
      {
        path: "home",
        element: <Home />,
      },
    ],
  },
  {
    path: "/login",
    element: <Login />,
    loader: loginPageLoader,
  },
  {
    path: "/signup",
    element: <Signup />,
    loader: loginPageLoader,
  },
  {
    path: "/verify",
    element: <VerifyEmail />,
  },
  {
    path: "*",
    element: <Error errorMessage={"Status 404: Page Not Found"} />,
  },
]);

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <ReactQueryDevtools />
      <RouterProvider router={Router}></RouterProvider>
    </QueryClientProvider>
  );
}

export default App;
