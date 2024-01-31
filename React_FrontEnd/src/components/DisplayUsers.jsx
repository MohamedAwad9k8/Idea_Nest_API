import PropTypes from "prop-types";
import { useState } from "react";
import { Link } from "react-router-dom";

const DisplayUsers = ({ users, onDeleteUser }) => {
  const [query, setQuery] = useState("");

  const updateQuery = (query) => {
    setQuery(query.trim());
  };

  const clearQuery = () => {
    updateQuery("");
  };

  const showingUsers =
    query === ""
      ? users
      : users.filter(
          (user) =>
            user.name.toLowerCase().includes(query.toLowerCase()) ||
            user.email.toLowerCase().includes(query.toLowerCase())
        );

  return (
    <div className="list-contacts">
      <div className="list-contacts-top">
        <input
          className="search-contacts"
          type="text"
          placeholder="Search Users"
          value={query}
          onChange={(event) => updateQuery(event.target.value)}
        />
      </div>
      {showingUsers.length !== users.length && (
        <div className="showing-contacts">
          <span>
            Now Showing {showingUsers.length} of {users.length} Users
          </span>
          <button onClick={clearQuery}>Show All</button>
        </div>
      )}
      <ol className="contact-list">
        {showingUsers.map((user) => (
          <li key={user.id} className="contact-list-item">
            <div className="contact-details">
              <p>{user.name}</p>
              <p>{user.email}</p>
              <p>{user.password}</p>
              <p>{user.ref_token}</p>
            </div>
            <button
              className="contact-remove"
              onClick={() => onDeleteUser(user)}
            >
              Remove
            </button>
          </li>
        ))}
      </ol>
      <Link to="/" className="close-create-contact">
        Close
      </Link>
    </div>
  );
};

DisplayUsers.propTypes = {
  users: PropTypes.array.isRequired,
  onDeleteUser: PropTypes.func.isRequired,
};

export default DisplayUsers;


// import PropTypes from "prop-types";
// import { useState } from "react";
// import { Link } from "react-router-dom";

// const DisplayUsers = ({ contacts, onDeleteContact }) => {
//   const [query, setQuery] = useState("");

//   const updateQuery = (query) => {
//     setQuery(query.trim());
//   };

//   const clearQuery = () => {
//     updateQuery("");
//   };
//   const showingContacts =
//     query === ""
//       ? contacts
//       : contacts.filter((c) =>
//           c.name.toLowerCase().includes(query.toLowerCase())
//         );

//   return (
//     <div className="list-contacts">
//       <div className="list-contacts-top">
//         <input
//           className="search-contacts"
//           type="text"
//           placeholder="Search Contacts"
//           value={query}
//           onChange={(event) => updateQuery(event.target.value)}
//         />
//         <Link to="/create" className="add-contact">
//           Add Contact
//         </Link>
//       </div>
//       {showingContacts.length !== contacts.length && (
//         <div className="showing-contacts">
//           <span>
//             Now Showing {showingContacts.length} of {contacts.length} Contacts
//           </span>
//           <button onClick={clearQuery}>Show All</button>
//         </div>
//       )}
//       <ol className="contact-list">
//         {showingContacts.map((contact) => (
//           <li key={contact.id} className="contact-list-item">
//             <div
//               className="contact-avatar"
//               style={{
//                 backgroundImage: `url(${contact.avatarURL})`,
//               }}
//             ></div>
//             <div className="contact-details">
//               <p>{contact.name}</p>
//               <p>{contact.handle}</p>
//             </div>
//             <button
//               className="contact-remove"
//               onClick={() => onDeleteContact(contact)}
//             >
//               Remove
//             </button>
//           </li>
//         ))}
//       </ol>
//     </div>
//   );
// };

// DisplayUsers.propTypes = {
//   contacts: PropTypes.array.isRequired,
//   onDeleteContact: PropTypes.func.isRequired,
// };

// export default DisplayUsers;
