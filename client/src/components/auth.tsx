// @ts-nocheck
import { useNavigate } from "react-router-dom";
import { useUserContext } from "../context/userContext";
import { useEffect } from "react";

const Auth = ({ children }: { children: any }) => {
  const { user } = useUserContext();
  const navigate = useNavigate();

  useEffect(() => {
    if (!user.username) {
      navigate("/home");
    }
  }, []);

  return children;
};

export default Auth;
