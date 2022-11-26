import './Gallery.css'

function GalleryPanel(props) {
	return (
		<div className="gallery-panel">
		    <h3 className="description-font">{props.description}</h3>
		    <img src={props.background_image} className="gallery-image" alt={props.alt}/>
		</div>
	);
}

function Gallery() {
	return (
		<div>
			<GalleryPanel
		    alt="Man smiling while looking at phone."
				description="Create an account."
				background_image="/images/sign-up.jpg"
			/>
		  <GalleryPanel
		    alt="Woman looking out the window while holding a book."
		    description="Ask for help, if you need it."
		    background_image="/images/woman-window.jpg"
		  />
		  <GalleryPanel
		    alt="A snow shoveler pushing snow along a sidewalk."
		    description="Offer help, if you can."
		    background_image="/images/shovel.jpg"
		  />
		  <GalleryPanel
		    alt="A man and a woman smiling while holding snow shovels."
		    description="Build a community."
		    background_image="/images/community.jpg"
		  />
		</div>
	);
}

export { Gallery };
