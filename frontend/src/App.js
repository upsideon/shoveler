import { BrowserRouter } from 'react-router-dom';

import './App.css';
import { Center } from './Center';
import { Footer } from './Footer';
import { Gallery } from './Gallery';
import { Navigation } from './Navigation';

function App() {
  return (
		<BrowserRouter>
			<div className="app">
				<Navigation />
				<Center />
				<Gallery />
				<Footer />
			</div>
		</BrowserRouter>
  );
}

export default App;
