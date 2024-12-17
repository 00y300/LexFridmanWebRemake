import { useState } from "react";
import { createPortal } from "react-dom";
import ImageCard from "./imageCard";

const ModalCardVideo = ({ children, imgSrc, ...props }) => {
  const [isModalOpen, setModalOpen] = useState(false);

  const openModal = () => setModalOpen(true);
  const closeModal = () => setModalOpen(false);

  return (
    <>
      {/* Trigger Card */}
      <div onClick={openModal}>
        <ImageCard imgSrc={imgSrc} {...props}>
          {children}
        </ImageCard>
      </div>

      {/* Modal */}
      {isModalOpen &&
        createPortal(
          <div
            className="fixed inset-0 z-50 flex items-center justify-center bg-black/70 p-2 sm:p-4"
            onClick={closeModal} // Close modal on background click
          >
            <div
              {...props}
              onClick={(e) => e.stopPropagation()} // Prevent background click within the Modal
              className="relative w-full max-w-[90%] rounded-lg bg-white p-4 shadow-lg sm:max-w-md md:max-w-lg md:p-6 lg:max-w-2xl"
            >
              {/* Close Button */}
              <button
                className="absolute right-2 top-2 text-2xl text-black"
                onClick={closeModal}
              >
                ✕
              </button>

              {/* Modal Content */}
              <div className="flex flex-col items-center">
                <img
                  src={imgSrc}
                  alt=""
                  className="h-auto w-full rounded-lg object-cover"
                />
                <div className="mt-4 w-full text-left">{children}</div>
              </div>
            </div>
          </div>,
          document.body,
        )}
    </>
  );
};

export default ModalCardVideo;
