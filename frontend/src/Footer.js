import { React } from 'react';

function Footer() {
  return (
    <div>
      <hr className="footer-hr" />
      <h4>
        Shoveler
        {new Date().getFullYear()}
      </h4>
    </div>
  );
}

export default Footer;
