import { Link } from "react-router";

const NavBar = () => {
  return (
    <div className="flex justify-between items-center">
      {/* Use flex-shrink to make sure the title doesn't push the UL to the next line */}
      <h1 className="text-3xl font-bold text-sky-300 flex-shrink-0">
        Lex Fridman Homepage
      </h1>
      <ul className="flex space-x-4">
        {/* <li className="p-2 border-2 border-orange-500 rounded-md"> */}

        <li className="p-2 border-2">
          <Link to="/">Home</Link>
        </li>
        {/* <li className="p-2 border-2 border-orange-500 rounded-md"> */}
        <li className="p-2 border-2">
          <Link to="/lectures">Lectures</Link>
        </li>
        {/* <li className="p-2 border-2 border-orange-500 rounded-md"> */}

        <li className="p-2 border-2">
          <Link to="/podcast">Podcast</Link>
        </li>
        {/* <li className="p-2 border-2 border-orange-500 rounded-md"> */}

        <li className="p-2 border-2">
          <Link to="/youtube">YouTube</Link>
        </li>
        {/* <li className="p-2 border-2 border-orange-500 rounded-md"> */}

        <li className="p-2 border-2">
          <Link to="/twitter">X Twitter</Link>
        </li>
        {/* <li className="p-2 border-2 border-orange-500 rounded-md"> */}

        <li className="p-2 border-2">
          <Link to="/contact">Contact Lex</Link>
        </li>
        {/* <li className="p-2 border-2 border-orange-500 rounded-md"> */}

        <li className="p-2 border-2">
          <Link to="/store">Store</Link>
        </li>
      </ul>
    </div>
  );
};

export default NavBar;
