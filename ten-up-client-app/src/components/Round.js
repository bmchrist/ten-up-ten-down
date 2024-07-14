import React from 'react';
import PlayerRound from './PlayerRound';

const Round = (props) => {
  const { round } = props;
  return (
    <div>
    <span>Round: {round.sequence}, Number of Cards: {round.numCards}</span>
    { round.playerRounds.map( playerRound =>
      <PlayerRound key={playerRound.id} playerRound={playerRound} />
    )}
    <br/>
    </div>
  );
};

export default Round;
