const api = process.env.REACT_APP_CONTACTS_API_URL || "http://localhost:8080";

const headers = {
  Accept: "application/json",
  "Content-Type": "application/json",
};

export const getAllUsers = () =>
  fetch(`${api}/users`, { headers })
    .then((res) => res.json())
    .then((data) => data);
    
export const signUpUser = async (user) => {
  try {
    const response = await fetch(`${api}/user/signup`, {
      method: "POST",
      headers,
      body: JSON.stringify(user),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error("Error during signup:", error);
    throw error; // Rethrow the error for the component to handle
  }
};

export const login = async (user) => {
    try {
      const response = await fetch(`${api}/user/signin`, {
        method: "POST",
        headers,
        body: JSON.stringify(user),
      });
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      return await response.json();
    } catch (error) {
      console.error("Error during signup:", error);
      throw error; // Rethrow the error for the component to handle
    }
  };