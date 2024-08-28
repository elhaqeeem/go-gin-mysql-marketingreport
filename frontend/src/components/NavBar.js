import React from 'react';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom

const NavBar = () => {
  return (
    <nav>
      <Link to="/">Home</Link>|
      <Link to="/komisi">Komisi</Link>
    </nav>
  );
};

export default NavBar;
