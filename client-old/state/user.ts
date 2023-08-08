'use client'

import { createContext } from "react";

type User = {
    name: string;
    email: string;
    pictureURI: string;
}

export const UserContext = createContext<User>({
    email: "",
    name: "",
    pictureURI: "",
})