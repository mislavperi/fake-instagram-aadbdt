import React from "react";
import ReactDOM from "react-dom/client";
import Login from "./pages/auth/login/page.tsx";
import Register from "./pages/auth/register/page.tsx";
import Home from "./pages/home/page.tsx";
import Auth from "./components/auth.tsx";
import Profile from "./pages/profile/page.tsx";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Navigation from "./components/navigation.tsx";
import { ChakraProvider } from "@chakra-ui/react";

import { UserProvider } from "./context/userContext.tsx";
import UploadImage from "./pages/upload/page.tsx";
import RedirectComp from "./components/redirect.tsx";
import Edit from "./pages/profile/edit/page.tsx";
import UserStatistics from "./pages/profile/admin/statistics/page.tsx";
import UserPictures from "./pages/profile/admin/pictures/page.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <RedirectComp>
        <Navigation />
      </RedirectComp>
    ),
    children: [
      {
        path: "home",
        element: <Home />,
      },
      {
        path: "login",
        element: <Login />,
      },
      {
        path: "register",
        element: <Register />,
      },
      {
        path: "upload",
        element: (
          <Auth>
            <UploadImage />
          </Auth>
        ),
      },
      {
        path: "profile/",
        element: (
          <Auth>
            <Profile />
          </Auth>
        ),
      },
      {
        path: "edit/:id",
        element: (
          <Auth>
            <Edit />
          </Auth>
        ),
      },
      {
        path: "statistics/:id",
        element: (
          <UserStatistics />
        ),
      },
      {
        path: "userimages/:id",
        element: (
          <UserPictures />
        ),
      }
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ChakraProvider>
      <UserProvider>
        <RouterProvider router={router} />
      </UserProvider>
    </ChakraProvider>
  </React.StrictMode>
);
