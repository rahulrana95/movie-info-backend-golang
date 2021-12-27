import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';

const Movies = () => {
  const [movies, setMovies] = React.useState([]);

  React.useEffect(() => {
    fetch('http://localhost:5000/v1/movies')
      .then((response) => response.json())
      .then((response) => {
        setMovies(response.movie);
      })
      .catch((err) => {
        console.error(err);
      });
  }, []);

  return (
    <div>
      Choose a Movie
      <ul className="list-group">
        {movies.map((movie) => {
          return (
            <li key={movie.id} className="list-group-item">
              <div>
                <img src={movie.poster_url} width="400px" alt="" />
              </div>
              <div className="">{movie.description}</div>
              {movie.url ? <div>
                <iframe width="-webkit-fill-available" title={movie.title} height="315" src={movie.url}></iframe>
              </div> : <div>Video not available.</div>}
              <Link to={`/movies/${movie.id}`}>
                {movie.title || 'Untitled'}
              </Link>
            </li>
          );
        })}
      </ul>
    </div>
  );
};

export default Movies;
