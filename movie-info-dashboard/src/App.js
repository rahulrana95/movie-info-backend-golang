import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Home from './components/Home';
import ManageCatalogue from './components/ManageCatalogue';
import Movies from './components/Movies';
import SingleMovie from './components/SingleMovie';
import Categories from './components/Categories';
import CategoryPage from './components/CategoryPage';

function App() {
  return (
    <Router>
      <div className="App">
        <div className="container">
          <div className="row">
            <h1 className="mt-3">Go watch a movie</h1>
            <hr className="mb-3"></hr>
          </div>
          <div className="row">
            <div className="col-md-4">
              <nav>
                <ul className="list-group">
                  <li className="list-group-item">
                    <Link to="/home">Home</Link>{' '}
                  </li>
                  <li className="list-group-item">
                    <Link to="/movies">Movies</Link>{' '}
                  </li>
                  <li className="list-group-item">
                    <Link to="/by-category">Categories</Link>{' '}
                  </li>
                  <li className="list-group-item">
                    <Link to="/manage-catalogue"> Manage Catalogue</Link>
                  </li>
                </ul>
              </nav>
            </div>
            <div className="col-md-8">
              <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/home" element={<Home />} />
                <Route path="/movies" element={<Movies />} />
                <Route exact path="/by-category" element={<CategoryPage />} />

                <Route exact path="/by-category/drama" render={(props) =><Categories {...props} title={'Drama'}  /> }  />

                <Route path="/movies/:id" element={<SingleMovie />} />

                <Route
                  path="/manage-catalogue"
                  element={<ManageCatalogue />}
                ></Route>
              </Routes>
            </div>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;
