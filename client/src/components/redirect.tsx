import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const RedirectComp = ({children}: any) =>  {
    const navigate = useNavigate()

    useEffect(() => {
        navigate("/home")
    }, [])

    return children
}

export default RedirectComp