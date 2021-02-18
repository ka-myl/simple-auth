import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const DashboardView = () => {
  const [msg, setMsg] = useState('Nothing yet!');
  const [error, setError] = useState(null);

  useEffect(() => {
    axios
      .get('http://localhost:8000/secret', { withCredentials: true })
      .then((res) => setMsg(res.data.msg))
      .catch(setError)
  }, [])

  console.log('ERROR: ', error)

  return (
    <div>
      <nav>
        <Link to="/login">Log in</Link>
        <Link to="/register">Register</Link>
      </nav>
      <div>
        {
          error
            ? 'Ooops, something went wrong'
            : msg
        }
      </div>
    </div>
  );
};

export default DashboardView;