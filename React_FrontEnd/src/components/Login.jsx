import { Link } from "react-router-dom";
import serializeForm from "form-serialize";
// import ImageInput from "./ImageInput";

const Login = ({ onLogin }) => {
  const handleSubmit = (e) => {
    e.preventDefault();

    const values = serializeForm(e.target, { hash: true });

    if (onLogin) {
      onLogin(values);
    }
  };
  return (
    <div>
      <Link to="/" className="close-create-contact">
        Close
      </Link>
      <form className="create-contact-form" onSubmit={handleSubmit}>
        {/* <ImageInput
          className="create-contact-avatar-input"
          name="avatarURL"
          maxHeight={64}
        /> */}
        <div className="create-contact-details">
          <input type="email" name="email" placeholder="john_doe@gmail.com" />
          <input type="password" name="password" placeholder="Password" />
          <button>Login</button>
        </div>
      </form>
    </div>
  );
};

export default Login;
