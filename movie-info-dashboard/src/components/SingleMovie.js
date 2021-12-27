import React from 'react';
import useRouterHooks from '../useRouterHooks';

const SingleMovie = () => {
  const { params } = useRouterHooks();
  const [movie, setMovie] = React.useState({
    name: 'Untitled',
    id: 1,
    runtime: '150 mins'
  });

  React.useEffect(() => { }, []);


  return <div>
    <h2>SingleMovie</h2>
    <table className="table table-compact table-strip">
      <thead></thead>
      <tbody>
        <tr>
          <td><strong>Title:</strong></td>
          <td>{movie.name} { movie.id}</td>
        </tr>
        <tr>
          <td><strong>Run Time:</strong></td>
          <td>{movie.runtime}</td>
        </tr>
      </tbody>
    </table>

  </div>
}

export default SingleMovie