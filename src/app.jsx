import { Route, Routes } from "react-router";
import NavBar from "./components/navBar";
import Homepage from "./pages/lexFridmanHomePage";
import Lectures from "./pages/Lectures";
import Podcast from "./pages/Podcast";
import Youtube from "./pages/Youtube";
import Twitter from "./pages/Twitter";
import Contact from "./pages/Contact";
import Store from "./pages/Store";

const App = () => {
  return (
    <div className="h-screen bg-customBackground">
      {/* <div> */}
      <NavBar />
      <Routes>
        <Route path="/" element={<Homepage />} />
        <Route path="/lectures" element={<Lectures />} />
        <Route path="/podcast" element={<Podcast />} />
        <Route path="/youtube" element={<Youtube />} />
        <Route path="/twitter" element={<Twitter />} />
        <Route path="/contact" element={<Contact />} />
        <Route path="/store" element={<Store />} />
      </Routes>
    </div>
  );
};

export default App;
