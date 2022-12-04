import React from 'react';
import PropTypes from 'prop-types';

function PopUpModal(props) {
  const { children, closeOnClick, id } = props;
  return (
    <div className="modal" id={id}>
      <span className="close" onClick={closeOnClick} title="Close">
        &times;
      </span>
      {children}
    </div>
  );
}

PopUpModal.propTypes = {
  children: PropTypes.object.isRequired,
  closeOnClick: PropTypes.func.isRequired,
  id: PropTypes.string.isRequired,
};

export default PopUpModal;
