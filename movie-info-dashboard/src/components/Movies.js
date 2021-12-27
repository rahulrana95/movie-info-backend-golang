import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';

const Movies = () => {
  const [movies, setMovies] = React.useState([]);

  React.useEffect(() => {
    setMovies([
      { id: 1, name: 'Spiderman: Homecoming' },
      { id: 2, name: 'Spiderman: No way home' },
    ]);
  }, []);

  return (
    <div>
      Choose a Movie
      <ul className="list-group">
        {movies.map((movie) => {
          return <li key={movie.id} className="list-group-item">
            <Link to={`/movies/${movie.id}`}>{movie.name || 'Untitled'}</Link></li>;
        })}
      </ul>
    </div>
  );
};

export default Movies;
