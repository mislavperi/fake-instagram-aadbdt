import User from "./user";

type Picture = {
  id: number;
  title: string;
  description: string;
  pictureURI: string;
  uploadDateTime: string;
  hashtags: string[];
  user: User;
};

export default Picture