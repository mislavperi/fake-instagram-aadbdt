import { useUserContext } from "../context/userContext"


const useAuth = () => {
    const { user } = useUserContext()

    return user.firstName !== ""
}

export default useAuth