import React, { useState } from 'react';
import { Link } from 'react-router-dom';

const DashboardView = () => {
  const [msg, setMsg] = useState('Nothing yet!')

  return (
    <div>
      <nav>
        <Link to="/login">Log in</Link>
        <Link to="/register">Register</Link>
      </nav>
      <div>
        Message from the server: {msg}
      </div>
    </div>
  );
};

export default DashboardView;