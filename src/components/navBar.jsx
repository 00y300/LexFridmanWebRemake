import { Link } from "react-router";

const NavBar = () => {
  return (
    <div className="flex justify-center items-center h-16 ">
      <ul className="flex space-x-4">
        <li className="p-2 border-2">
          <Link to="/">Home</Link>
        </li>
        <li className="p-2 border-2">
          <Link to="/lectures">Lectures</Link>
        </li>
        <li className="p-2 border-2">
          <Link to="/podcast">Podcast</Link>
        </li>
        <li className="p-2 border-2">
          <Link to="/youtube">YouTube</Link>
        </li>
        <li className="p-2 border-2">
          <Link to="/twitter">X Twitter</Link>
        </li>
        <li className="p-2 border-2">
          <Link to="/contact">Contact Lex</Link>
        </li>
        <li className="p-2 border-2">
          <Link to="/store">Store</Link>
        </li>
      </ul>
    </div>
  );
};

export default NavBar;
