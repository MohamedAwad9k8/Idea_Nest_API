import { useState, useEffect } from "react";
import { Route, Routes, useNavigate } from "react-router-dom";
import "../css/App.css";
import DisplayUsers from "./DisplayUsers";
import SignUp from "./SignUp";
import * as ContactsAPI from "../utils/ContactsAPI";
import * as UsersAPI from "../utils/UsersAPI";
import MainPage from "./MainPage";
import Login from "./Login";


const App = () => {
  let navigate = useNavigate();
  const [users, setUsers] = useState([]);
  const [loggedUser, setLoggedUser] = useState([]);
  
  // const removeContact = (contact) => {
  //   ContactsAPI.remove(contact);
  //   setContacts(contacts.filter((c) => c.id !== contact.id));
  // };
  //onDeleteContact={removeContact}
  const handleSignUp = (user) => {
    const create = async () => {
      const res = await UsersAPI.signUpUser(user);
      setUsers([...users, res]);
    };

    create();
    navigate("/");
  };

  const handleLogin = (user) => {
    const login = async () => {
      try {
        const res = await UsersAPI.login(user);
        const { accessToken, refreshToken } = res;

        // Decode the base64-encoded payload of the access token
        const decodedAccessToken = atob(accessToken.split('.')[1]);
        const decodedAccessTokenObj = JSON.parse(decodedAccessToken);

        const decodedRefreshToken = atob(refreshToken.split('.')[1]);
        const decodedRefreshTokenObj = JSON.parse(decodedRefreshToken);

        // Save the decoded information to loggedUser
        setLoggedUser({
          id: decodedRefreshTokenObj.id, // Assuming id is present in the access token
          name: decodedAccessTokenObj.name, // Replace with the correct field name
          email: decodedAccessTokenObj.email, // Replace with the correct field name
          password: decodedAccessTokenObj.password, // Replace with the correct field name
          accessToken,
          refreshToken,
        });
      } catch (error) {
        console.error('Error during login:', error);
      }
    };

    login();
    navigate('/');
  };

  

  useEffect(() => {
    const getUsers = async () => {
      const res = await UsersAPI.getAllUsers();
      setUsers(res);
    };
    getUsers();
  }, []);

  return (
    <Routes>
      <Route
        exact
        path="/"
        element={
          <MainPage loggedUser={loggedUser} />
        }
      />
      <Route
        exact
        path="/users"
        element={
          <DisplayUsers users={users}  />
        }
      />
      <Route
        exact
        path="/signup"
        element={
          <SignUp
            onSignUp={(user) => handleSignUp(user)}
          />
        }
      />
      <Route
        exact
        path="/login"
        element={
          <Login
            onLogin={(user) => handleLogin(user)}
          />
        }
      />
    </Routes>
  );
};

export default App;
