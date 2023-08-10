import React from "react";
import ReactDOM from "react-dom/client";
import Login from "./pages/auth/login/page.tsx";
import Home from "./pages/home/page.tsx";
import Welcome from "./pages/welcome/page.tsx";
import Auth from "./components/auth.tsx";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Navigation from "./components/navigation.tsx";

import { UserProvider } from "./context/userContext.tsx";
import UploadImage from "./pages/upload/page.tsx";
import RedirectComp from "./components/redirect.tsx";

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
        path: "upload",
        element: (
          <Auth>
            <UploadImage />
          </Auth>
        ),
      },
    ],
  },
  {
    path: "/welcome",
    element: (
      <Auth>
        <Welcome />
      </Auth>
    ),
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <UserProvider>
      <RouterProvider router={router} />
    </UserProvider>
  </React.StrictMode>
);
