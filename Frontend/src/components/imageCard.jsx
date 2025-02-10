const ImageCard = ({ children, imgSrc, ...props }) => {
  return (
    <div
      {...props}
      className="group relative h-80 max-w-xs overflow-hidden rounded-2xl shadow-lg"
    >
      <img
        src={imgSrc}
        alt=""
        className="h-full w-full object-cover transition-transform duration-200 group-hover:scale-110"
      />
      <div className="absolute inset-0 flex items-end bg-gradient-to-t from-black/60 to-transparent">
        <div className="p-4 text-white">{children}</div>
      </div>
    </div>
  );
};

export default ImageCard;
