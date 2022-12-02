import React from 'react';
import PropTypes from 'prop-types';

import './Gallery.css';

function GalleryPanel(props) {
  const { alt, description, backgroundImage } = props;
  return (
    <div className="gallery-panel">
      <h3 className="description-font">{description}</h3>
      <img src={backgroundImage} className="gallery-image" alt={alt} />
    </div>
  );
}

GalleryPanel.propTypes = {
  alt: PropTypes.string.isRequired,
  description: PropTypes.string.isRequired,
  backgroundImage: PropTypes.string.isRequired,
};

function Gallery() {
  return (
    <div>
      <GalleryPanel
        alt="Man smiling while looking at phone."
        description="Create an account."
        backgroundImage="/images/sign-up.jpg"
      />
      <GalleryPanel
        alt="Woman looking out the window while holding a book."
        description="Ask for help, if you need it."
        backgroundImage="/images/woman-window.jpg"
      />
      <GalleryPanel
        alt="A snow shoveler pushing snow along a sidewalk."
        description="Offer help, if you can."
        backgroundImage="/images/shovel.jpg"
      />
      <GalleryPanel
        alt="A man and a woman smiling while holding snow shovels."
        description="Build a community."
        backgroundImage="/images/community.jpg"
      />
    </div>
  );
}

export default Gallery;
