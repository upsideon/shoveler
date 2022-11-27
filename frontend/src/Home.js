import { React } from 'react';

import Center from './Center';
import Footer from './Footer';
import Gallery from './Gallery';
import Navigation from './Navigation';

function Home() {
  return (
    <div className="app">
      <Navigation />
      <Center />
      <Gallery />
      <Footer />
    </div>
  );
}

export default Home;
