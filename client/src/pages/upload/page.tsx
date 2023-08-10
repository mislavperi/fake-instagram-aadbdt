import Cookies from "universal-cookie";

import { useContext, useEffect, useState } from "react"
import User from "../../types/user";

import { Button } from "@chakra-ui/react";

import { UserContext } from "../../context/userContext";

export default function UploadImage() {
    const cookieJar = new Cookies();
    const accessToken = cookieJar.get("accessToken")
    const refreshToken = cookieJar.get("refreshToken")
  
    const user = useContext(UserContext)
  
    const [picture, setPicture] = useState()
    const [title, setTitle] = useState("")
    const [description, setDescription] = useState("")
    const [hashtags, setHashtags] = useState([])

    function handleChange(e) {
        e.preventDefault()
        setPicture(e.target.files[0])
    }

    function submitFile() {
        const formData = new FormData()
        if (picture != undefined) {
            formData.append("file", picture)
            formData.append("title", title) 
            formData.append("description", description)
            formData.append("hashtags", JSON.stringify(hashtags))
            formData.append("user", JSON.stringify(user))
        }

        fetch("http://localhost:8080/picture/upload", {
            method: "POST",
            headers: {
                Authorization: accessToken,
                Refresh: refreshToken
            },
            body: formData,
        })
    }

    return (
        <div>
            <input type="file" onChange={handleChange} />
            <Button onClick={submitFile}>Upload</Button>
        </div>
    )
}