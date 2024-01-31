import { Link } from "react-router-dom";
import serializeForm from "form-serialize";
// import ImageInput from "./ImageInput";

const SignUp = ({ onSignUp }) => {
  const handleSubmit = (e) => {
    e.preventDefault();

    const values = serializeForm(e.target, { hash: true });

    if (onSignUp) {
      onSignUp(values);
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
          <input type="text" name="name" placeholder="John Doe" />
          <input type="email" name="email" placeholder="john_doe@gmail.com" />
          <input type="password" name="password" placeholder="Password" />
          <button>Sign Up</button>
        </div>
      </form>
    </div>
  );
};

export default SignUp;
