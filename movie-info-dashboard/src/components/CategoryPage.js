import React from 'react';
import { Link } from 'react-router-dom';
import useRouterHooks from '../useRouterHooks';

const CategoryPage = () => {
  const { path, url } = useRouterHooks();
  return (
    <div>
      <h2>Categories</h2>

      <ul>
        <li>
          <Link to={`${path}/comedy`}>Comedy</Link>
        </li>
        <li>
          <Link to={`${path}/drama`}>Drama</Link>
        </li>
        <li>
          <Link to={`${path}/action`}>Action</Link>
        </li>
      </ul>
    </div>
  );
};

export default CategoryPage;
