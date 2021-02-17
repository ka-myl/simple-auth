import React, { useState } from 'react';
import axios from 'axios';
import { useHistory } from 'react-router-dom'

const LoginView = () => {
  const history = useHistory();

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    setError(null);
    setLoading(true);

    try {
      await axios.post('http://localhost:8000/login', { username, password }, { withCredentials: true });
      history.push('/');
    } catch (err) {
      setError(err);
      setLoading(false);
    }
  };

  return (
    <div>
      <h1>Log In</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>
            Username
            <input value={username} onChange={(e) => setUsername(e.target.value)} />
          </label>
        </div>
        <div>
          <label>
            Password
            <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
          </label>
        </div>
        <div>
          <button type="submit">Log in</button>
        </div>
      </form>
    </div>
  );
};

export default LoginView;
