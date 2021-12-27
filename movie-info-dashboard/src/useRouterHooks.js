import { useParams,useLocation } from 'react-router-dom';


const useRouterHooks = () => {
  const location = useLocation();
  console.log(location)
  return {
    params: useParams(),
    url:'',
    path:location.pathname
  }
}

export default useRouterHooks