// @ts-nocheck
import Cookies from "universal-cookie";

import { useEffect } from "react";

import FormValues from "../../../types/form";

import { useNavigate } from "react-router-dom";

import {
  FormControl,
  FormLabel,
  FormErrorMessage,
  Button,
  Input,
  useToast,
  Link,
  Box
} from "@chakra-ui/react";
import { Formik, Form, Field, FieldProps, FormikProps } from "formik";
import { useUserContext } from "../../../context/userContext";
import User from "../../../types/user";
import { GoogleOAuthProvider, GoogleLogin } from "@react-oauth/google";

export default function Login() {
  const cookieJar = new Cookies();
  const accessToken = cookieJar.get("accessToken");
  const refreshToken = cookieJar.get("refreshToken");
  const navigate = useNavigate();

  const toast = useToast();
  const { dispatch } = useUserContext();

  useEffect(() => {
    fetch("/api/user/whoami", {
      method: "GET",
      headers: {
        Accept: "application/json",
        Authorization: accessToken,
        Refresh: refreshToken,
      },
    }).then((res) => {
      if (res.ok) {
        res.json().then((res) => {
          const user: User = {
            firstName: res?.firstName,
            lastName: res?.lastName,
            email: res.email,
            username: res.username,
            pictureURI: res?.pictureURI,
            plan: res?.plan,
          };

          dispatch({
            type: "UPDATE_USER",
            payload: user,
          });
        });
        navigate("/home");
      }
    });
  }, []);

  useEffect(() => {
    const hasCode = window.location.href.includes("?code=");

    if (hasCode) {
      const code = window.location.href.split("?code=")[1];

      fetch("/api/auth/gh_login", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          code: code,
        }),
      }).then((res) => {
        if (res.ok) {
          res.json().then((_) => {
            toast({ description: "Sucess" });
          });
        }
      });
    }
  }, []);

  const initialValues = {
    username: "",
    password: "",
  };

  const validateUsername = (value: string) => {
    let error;
    if (!value || value.length < 2) {
      error = "Username must have at least 2 characters";
    }
    return error;
  };

  const onSubmit = (values: FormValues) => {
    fetch("/api/auth/login", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        username: values.username,
        password: values.password,
      }),
    })
      .then((res) => {
        if (!res.ok) {
          res.json().then(res => {
            toast({
              description: res,
            });
          })
        } else {
          res.json().then((_) => {
            toast({ description: "Success" });
          });
        }
      })
  };

  return 1 == 1 ? (
    <GoogleOAuthProvider clientId="1004370425511-p05kgqo2ot4b90icjda6t1272cpuso0d.apps.googleusercontent.com">
      <Formik initialValues={initialValues} onSubmit={onSubmit}>
        {(props: FormikProps<FormValues>) => (
          <Form>
            <Field name="username" validate={validateUsername}>
              {(fieldProps: FieldProps<string>) => (
                <FormControl
                  id="username"
                  isInvalid={!!props.errors.username && props.touched.username}
                >
                  <FormLabel>Username</FormLabel>
                  <Input {...fieldProps.field} type="text" />
                  <FormErrorMessage>{props.errors.username}</FormErrorMessage>
                </FormControl>
              )}
            </Field>

            <Field name="password">
              {(fieldProps: FieldProps<string>) => (
                <FormControl
                  id="password"
                  mt={4}
                  isInvalid={!!props.errors.password && props.touched.password}
                >
                  <FormLabel>Password</FormLabel>
                  <Input {...fieldProps.field} type="password" />
                  <FormErrorMessage>{props.errors.password}</FormErrorMessage>
                </FormControl>
              )}
            </Field>

            <Button mt={4} colorScheme="blue" type="submit">
              Log In
            </Button>
            <Button onClick={() => navigate("/register")}>
              Register
            </Button>
            <GoogleLogin
              onSuccess={(credentialResponse) => {
                fetch("/api/auth/g_login", {
                  method: "POST",
                  credentials: "include",
                  headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                  },
                  body: JSON.stringify({
                    token: credentialResponse.credential,
                  }),
                });
              }}
              onError={() => {
                console.log("Login Failed");
              }}
            />
            <Button>
              <a href="https://github.com/login/oauth/authorize?client_id=f4a7a59f8e527f183bcf">
                Login with Github
              </a>
            </Button>
          </Form>
        )}
      </Formik>
    </GoogleOAuthProvider>
  ) : undefined;
}
