import React, { useState } from "react";
import { FaAlignRight } from "react-icons/fa";

export default function Navbar() {
  const [toggle, setToggle] = useState(false);

  const Toggle = () => {
    setToggle(!toggle);
  };

  return (
    <>
      <div className="navBar">
        <button onClick={Toggle}>
          <FaAlignRight />
        </button>
        <ul className={toggle ? "nav-links show-nav" : "nav-links"}>
          <li ref="#">Home</li>
          <li ref="#">Playlists</li>
          <li ref="#">Account</li>
        </ul>
      </div>
    </>
  );
}
