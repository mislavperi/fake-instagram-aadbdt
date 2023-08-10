import Plan from "./plan";

type User = {
    firstName: string;
    lastName: string;
    email: string;
    username: string;
    plan: Plan;
    pictureURI: string;
};

export default User