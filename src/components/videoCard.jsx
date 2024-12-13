import { useState, useEffect, useCallback } from "react";

const VideoCard = ({ children, videoSrc, ...props }) => {
  const [isCentered, setIsCentered] = useState(false);
  const [cardStyle, setCardStyle] = useState({});
  const [currentRect, setCurrentRect] = useState(null);

  const centerCard = (rect) => {
    const viewportCenterX = window.innerWidth / 2;
    const viewportCenterY = window.innerHeight / 2;

    const offsetX = viewportCenterX - (rect.left + rect.width / 2);
    const offsetY = viewportCenterY - (rect.top + rect.height / 2);

    setCardStyle({
      position: "fixed",
      top: `${rect.top}px`,
      left: `${rect.left}px`,
      width: `${rect.width}px`,
      height: `${rect.height}px`,
      transform: `translate(${offsetX}px, ${offsetY}px) scale(1.5)`,
      zIndex: 50,
      transition: "transform 0.5s, width 0.5s, height 0.5s",
    });
  };

  const handleCardClick = (e) => {
    if (!isCentered) {
      const card = e.currentTarget;
      const rect = card.getBoundingClientRect();
      setCurrentRect(rect); // Store the card's original position and size
      centerCard(rect); // Center the card
    } else {
      setCardStyle({}); // Reset styles
      setCurrentRect(null); // Clear the stored rect
    }
    setIsCentered(!isCentered); // Toggle the centered state
  };

  const handleResize = useCallback(() => {
    if (isCentered && currentRect) {
      centerCard(currentRect); // Re-center the card using the stored rect
    }
  }, [isCentered, currentRect]); // Add dependencies here

  useEffect(() => {
    window.addEventListener("resize", handleResize); // Add resize listener
    return () => {
      window.removeEventListener("resize", handleResize); // Clean up listener
    };
  }, [handleResize]); // Add handleResize to the dependency array

  return (
    <>
      {isCentered && (
        <div
          onClick={handleCardClick} // Close the card when background is clicked
          className="fixed inset-0 z-40 bg-black/50 backdrop-blur-md"
        ></div>
      )}
      <div
        {...props}
        onClick={handleCardClick}
        style={cardStyle}
        className="relative h-80 max-w-xs transform overflow-hidden rounded-2xl shadow-lg"
      >
        <div className="h-full w-full">
          <video
            src={videoSrc}
            autoPlay
            muted
            loop
            className="h-full w-full object-cover"
          />
        </div>
        <div className="absolute inset-0 flex items-end bg-gradient-to-t from-black/60 to-transparent">
          <div className="p-4 text-white">{children}</div>
        </div>
      </div>
    </>
  );
};

export default VideoCard;
