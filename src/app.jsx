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
    <>
      <NavBar />
      <Routes>
        <Route path="/" element={<Homepage />}></Route>
        <Route path="/lectures" element={<Lectures />}></Route>
        <Route path="/podcast" element={<Podcast />}></Route>
        <Route path="/youtube" element={<Youtube />}></Route>
        <Route path="/twitter" element={<Twitter />}></Route>
        <Route path="/contact" element={<Contact />}></Route>
        <Route path="/store" element={<Store />}></Route>
      </Routes>
    </>
  );
};

export default App;
