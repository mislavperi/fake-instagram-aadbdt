import Plan from "./plan";

type User = {
    [key: string]: any;
    id: number;
    firstName: string;
    lastName: string;
    email: string;
    username: string;
    plan: Plan;
    pictureURI: string;
};

export default User