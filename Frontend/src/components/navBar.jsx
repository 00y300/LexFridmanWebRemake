import Link from "next/link";

const NavBar = () => {
  return (
    <div className="flex h-16 items-center justify-center">
      <ul className="flex space-x-4">
        <li className="border-2 p-2">
          <Link href="/">Home</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/Lectures">Lectures</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/Podcast">Podcast</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/Youtube">YouTube</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/Twitter">X Twitter</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/Contact">Contact Lex</Link>
        </li>
        <li className="border-2 p-2">
          <Link href="/Store">Store</Link>
        </li>
      </ul>
    </div>
  );
};

export default NavBar;
