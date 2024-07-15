import React from 'react';
import PlayerRound from './PlayerRound';

const Round = (props) => {
  const { round } = props;
  return (
    <tr>
      <td>
        {round.numCards}
      </td>
      { round.playerRounds.map( playerRound =>
        <PlayerRound key={"" + playerRound.round + playerRound.player} playerRound={playerRound} />
      )}
    </tr>
  );
};

export default Round;
