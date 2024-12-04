import lexFridmanPhoto from "../../img/lex_fridman_deep_learning_course.jpg"; // Import the image

const Homepage = () => {
  return (
    <div className="mt-10 flex flex-col items-center">
      <h1 className="bg-gradient-to-r from-gray-300 to-stone-400 bg-clip-text text-7xl font-bold text-transparent">
        Lex Fridman
      </h1>

      <img
        className="mt-10 h-96 w-96 rounded-full"
        src={lexFridmanPhoto} // Use the imported image
        alt="Photo of Lex Fridman"
      />
    </div>
  );
};

export default Homepage;
