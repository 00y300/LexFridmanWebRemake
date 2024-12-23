import Image from "next/image";

const Hero = ({ title, description, imgSrc }) => {
  return (
    <div className="relative flex flex-col-reverse items-center justify-between space-y-4 px-6 py-12 md:flex-row md:space-x-4 md:space-y-0 lg:space-x-0 lg:px-44 lg:py-20">
      {/* Left Text Section */}
      <div className="max-w-lg md:text-left lg:text-left">
        <h1 className="text-3xl font-bold text-gray-800 lg:text-8xl">
          {title}
        </h1>
        <p className="mt-2 text-gray-600 lg:text-lg">{description}</p>
      </div>

      {/* Right Image Section */}
      <div className="w-full max-w-md md:max-w-lg">
        <Image
          src={imgSrc}
          alt="Hero Image"
          className="rounded-lg object-cover shadow-lg"
          width={800}
          height={600}
        />
      </div>
    </div>
  );
};

export default Hero;
