import { useEffect, useState } from "react";

export default function VideoList() {
  const [videos, setVideos] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8000/yt")
      .then((res) => res.json())
      .then((data) => {
        const videoEntries = Object.entries(data.videos);
        setVideos(videoEntries);
      })
      .catch(console.error);
  }, []);

  return (
    <div>
      {videos.map(([id, title]) => (
        <iframe
          key={id}
          width="560"
          height="315"
          src={`https://www.youtube.com/embed/${id}`}
          title={title}
          frameBorder="0"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
          allowFullScreen
        />
      ))}
    </div>
  );
}
