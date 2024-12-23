import Link from "next/link";

const NavBar = () => {
  return (
    <div className="flex h-16 items-center justify-center">
      <ul className="flex space-x-4">
        <li className="border-2 p-2">
          <Link href="/">Home</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/lectures">Lectures</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/podcast">Podcast</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/youtube">YouTube</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/twitter">X Twitter</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/contact">Contact Lex</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/store">Store</Link>
        </li>
      </ul>
    </div>
  );
};

export default NavBar;
