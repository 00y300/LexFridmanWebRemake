import Link from "next/link";

const ButtonsLink = ({ title, page }) => {
  return (
    <div>
      <Link href={page}>
        <button className="rounded bg-gray-500 px-4 py-2 font-bold text-white hover:bg-gray-700">
          {title}
        </button>
      </Link>
    </div>
  );
};

export default ButtonsLink;
