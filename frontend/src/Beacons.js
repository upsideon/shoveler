import axios from 'axios';
import React from 'react';
import { useState } from 'react';

import AddIcon from '@mui/icons-material/Add';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import Checkbox from '@mui/material/Checkbox';
import Fab from '@mui/material/Fab';
import IconButton from '@mui/material/IconButton';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import RemoveCircleOutlineIcon from '@mui/icons-material/RemoveCircleOutline';
import TextField from '@mui/material/TextField';

import {StatusCodes as HttpStatus} from 'http-status-codes';
import { getHeaders } from './SignUp';

class Beacon extends React.Component {
  constructor(props) {
    super(props);
    this.props = props;
  }

  render() {
    const props = this.props;
    return (
      <ListItem
        divider
        secondaryAction={
          props.removable ?
            <IconButton onClick={props.deleteOnClick}>
              <RemoveCircleOutlineIcon/>
            </IconButton> :
            <Checkbox />
        }
      >
        <ListItemButton>
          <ListItemAvatar>
            <Avatar />
          </ListItemAvatar>
          <ListItemText primary={props.address} />
        </ListItemButton>
      </ListItem>
    );
  }
}

class AddBeaconForm extends React.Component {
  constructor(props) {
    super(props);
    this.props = props;
    this.state = {};
    this.handleClose = this.props.handleClose;
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  async handleSubmit(event) {
    event.preventDefault();
    const addressTextField = document.getElementById('address');
    const address = addressTextField.value;
    const response = await axios.post(
      'http://localhost:8080/beacons',
      { address },
      getHeaders(),
    );
    this.handleClose();
  }

  render() {
    return (
       <Dialog
         open={this.props.open}
         onClose={this.handleClose}
       >
        <DialogTitle>Beacon Creation</DialogTitle>
        <DialogContent>
          <DialogContentText>
            To create a beacon to notify others that you would like help shoveling,
            please enter the following information.
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            id="address"
            label="Address"
            type="address"
            fullWidth
            variant="standard"
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={this.handleClose}>Cancel</Button>
          <Button onClick={this.handleSubmit}>Submit</Button>
        </DialogActions>
      </Dialog>
    );
  }
}

class AddBeacon extends React.Component {
  constructor(props) {
    super(props);
    this.props = props;
    this.state = {
      beacons: [],
      addModalOpen: false,
    };
    this.addOnClick = this.addOnClick.bind(this);
    this.deleteOnClick = this.deleteOnClick.bind(this);
    this.getBeacons = this.getBeacons.bind(this);
    this.handleClose = this.handleClose.bind(this);
  }

  async componentDidMount () {
    const beacons = await this.getBeacons();
    this.setState({ beacons });
  }

  async getBeacons() {
    const response = await axios.get(
      'http://localhost:8080/beacons',
      getHeaders(),
    );
    const beacons = JSON.parse(response.data);
    return beacons;
  }

  addOnClick() {
    this.setState({ addModalOpen: true });
  }

  deleteOnClick(id) {
    return async () => {
      const response = await axios.delete(
        `http://localhost:8080/beacons/${id}`,
        getHeaders(),
      );

      if (response.status === HttpStatus.OK) {
        this.setState({
          beacons: this.state.beacons.filter(b => b.id !== id),
        });
      }
    };
  }

  async handleClose() {
    this.setState({
      addModalOpen: false,
      beacons: await this.getBeacons(),
    });
  }

  render() {
    return (
      <div className="beacon-sections add-beacon">
        <h2>Add a Beacon</h2>
        <div className="add-beacon-subheader">
          <Fab
            className="add-button"
            onClick={this.addOnClick}
            color="primary"
            aria-label="add"
          >
            <AddIcon />
          </Fab>
          <p className="add-description">If you are looking for help, add a beacon here.</p>
        </div>
        {
          this.state.beacons.map(beacon => {
            const id = beacon.id;
            return <Beacon
              address={beacon.address}
              deleteOnClick={this.deleteOnClick(id)}
              id={id}
              key={id}
              removable
            />
          })
        }
        <AddBeaconForm open={this.state.addModalOpen} handleClose={this.handleClose}/>
      </div>
    );
  }
}

function FindBeacon() {
  return (
    <div className="beacon-sections find-beacon">
      <h2>Open Beacons</h2>
      <p>If you would like to help someone shovel, claim a beacon here.</p>
      <List>
        <Beacon address="30 Rockefeller Plaza, New York, NY"/>
        <Beacon address="1600 Pennsylvania Avenue NW, Washington, DC 20500"/>
      </List>
    </div>
  );
}

function Beacons() {
  return (
    <div>
      <AddBeacon />
      <FindBeacon />
    </div>
  );
}

export default Beacons;
