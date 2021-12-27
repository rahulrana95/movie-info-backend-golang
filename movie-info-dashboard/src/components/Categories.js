import React from 'react';
import { Link } from 'react-router-dom';
import useRouterHooks from '../useRouterHooks';

const Categories = (props) => {
  const { path, url } = useRouterHooks();
  return (
    <div>
      <div>
        {props.title}
      </div>
    </div>
  );
};

export default Categories;
