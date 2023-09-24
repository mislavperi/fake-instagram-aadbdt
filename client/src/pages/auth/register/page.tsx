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
} from "@chakra-ui/react";
import { Formik, Form, Field, FieldProps, FormikProps } from "formik";
import { useUserContext } from "../../../context/userContext";
import User from "../../../types/user";
import { GoogleOAuthProvider, GoogleLogin } from "@react-oauth/google";

export default function Register() {
  const cookieJar = new Cookies();
  const accessToken = cookieJar.get("accessToken")
  const refreshToken = cookieJar.get("refreshToken")
  const navigate = useNavigate();

  const toast = useToast();
  const { dispatch } = useUserContext();

  const initialValues = {
    email: "",
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
  const validateEmail = (value: string) => {
    let error;
    if (!value || value.length < 2) {
      error = "Email must have at least 2 characters";
    }
    return error;
  };

  const onSubmit = (values: FormValues) => {
    fetch("/api/auth/register", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({
        firstName: "",
        lastName: "",
        email: values.email,
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
      <Formik initialValues={initialValues} onSubmit={onSubmit}>
        {(props: FormikProps<FormValues>) => (
          <Form>
            <Field name="email" validate={validateEmail}>
              {(fieldProps: FieldProps<string>) => (
                <FormControl
                  id="email"
                  isInvalid={!!props.errors.email && props.touched.email}
                >
                  <FormLabel>Email</FormLabel>
                  <Input {...fieldProps.field} type="text" />
                  <FormErrorMessage>{props.errors.email}</FormErrorMessage>
                </FormControl>
              )}
            </Field>
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
              Register
            </Button>
          </Form>
        )}
      </Formik>
  ) : undefined;
}
