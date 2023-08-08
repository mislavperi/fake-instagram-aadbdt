"use client";

import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

import { useToast } from "@/components/ui/use-toast";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Toaster } from "@/components/ui/toaster";

import { Input } from "@/components/ui/input";
import { useEffect } from "react";

import { GoogleOAuthProvider, GoogleLogin } from "@react-oauth/google";

const formSchema = z.object({
  username: z.string().min(2, {
    message: "Username must be at least 2 characters long",
  }),
  password: z.string().min(8, {
    message: "Password has to be at least 8 characters long",
  }),
});

export default function LoginForm() {
  const { toast } = useToast();

  useEffect(() => {
    const hasCode = window.location.href.includes("?code=");

    if (hasCode) {
      const code = window.location.href.split("?code=")[1];

      fetch("http://localhost:8080/auth/gh_login", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          code: code,
        }),
      })
        .then((res) => res.json())
        .then((res) => {
          if (res.ok) {
            console.log("NIGGER")
            router.push("/");
          }
        });
    }
  }, []);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    const response = await fetch("http://localhost:8080/auth/login", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
      },
      body: JSON.stringify({
        username: values.username,
        password: values.password,
      }),
    });
    if (!response.ok) {
      toast({
        description: response.json(),
      });
    } else {
      toast({
        description: "Login was successful",
      });
      router.push("/");
    }
  }

  return (
    <GoogleOAuthProvider clientId="1004370425511-p05kgqo2ot4b90icjda6t1272cpuso0d.apps.googleusercontent.com">
      <main className="flex items-center bg-slate-50 h-screen justify-center">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <FormField
              control={form.control}
              name="username"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input placeholder="username" {...field} className="w-80" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input placeholder="password" {...field} className="w-80" />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <p className="text-xs font-light text-gray-500 p-0 m-0">
              New here? Click here to register
            </p>
            <Button type="submit">Login</Button>
            <Button asChild>
              <Link href="https://github.com/login/oauth/authorize?client_id=f4a7a59f8e527f183bcf">
                Login with Github
              </Link>
            </Button>
            <GoogleLogin
              onSuccess={(credentialResponse) => {
                fetch("http://localhost:8080/auth/g_login", {
                  method: "POST",
                  credentials: "include",
                  headers: {
                    "Accept": "application/json",
                    "Content-Type": "application/json",
                  },
                  body: JSON.stringify({
                    token: credentialResponse.credential,
                  }),
                })
                  .then((res) => res.json())
                  .then((res) => {
                    if (res.ok) {
                      router.push("/");
                    }
                  }
                    
                  );
              }}
              onError={() => {
                console.log("Login Failed");
              }}
            />
          </form>
        </Form>
        <Toaster />
      </main>
    </GoogleOAuthProvider>
  );
}
