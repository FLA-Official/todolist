import { useState, useEffect, useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { getProfile } from "../api/auth";

const Profile = () => {
  const { token } = useContext(AuthContext);
  const [user, setUser] = useState(null);

  useEffect(() => {
    if (!token) return;
    getProfile(token).then(res => res.json()).then(setUser);
  }, [token]);

  if (!user) return null;

  return (
    <div className="profile">
      <h2>{user.fullname}</h2>
      <p>{user.gmail}</p>
    </div>
  );
};

export default Profile;