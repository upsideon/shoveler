import './App.css';

import { Center } from './Center';
import { Footer } from './Footer';
import { Gallery } from './Gallery';
import { Navigation } from './Navigation';

function App() {
  return (
    <div className="app">
			<Navigation />
			<Center />
		  <Gallery />
		  <Footer />
    </div>
  );
}

export default App;
