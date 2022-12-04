import axios from 'axios';
import React from 'react';

import AddIcon from '@mui/icons-material/Add';
import Avatar from '@mui/material/Avatar';
import Checkbox from '@mui/material/Checkbox';
import Fab from '@mui/material/Fab';
import IconButton from '@mui/material/IconButton';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import RemoveCircleOutlineIcon from '@mui/icons-material/RemoveCircleOutline';

import {jwtStore} from './SignUp';

function Beacon(props) {
  return (
    <ListItem
      divider
      secondaryAction={
        props.removable ? 
          <IconButton>
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

class AddBeacon extends React.Component {
  constructor(props) {
    super(props);
    this.props = props;
    this.state = {
      beacons: [],
    };
  }

  async componentDidMount () {
    const token = jwtStore.getState().token.payload;
    const bearer = `Bearer ${token}`;
    const response = await axios.get(
      'http://localhost:8080/beacons',
      { headers: { 'Authorization': bearer } },
    );
    const beacons = JSON.parse(response.data);
    this.setState({ beacons });
  }

  render() {
    return (
      <div className="beacon-sections add-beacon">
        <h2>Add a Beacon</h2>
        <div className="add-beacon-subheader">
          <Fab className="add-button" color="primary" aria-label="add">
            <AddIcon />
          </Fab>
          <p className="add-description">If you are looking for help, add a beacon here.</p>
        </div>
        {
          this.state.beacons.map(beacon => <Beacon address={beacon.address} removable />)
        }
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
