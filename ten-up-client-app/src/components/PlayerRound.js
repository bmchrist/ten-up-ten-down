import React from 'react';

const PlayerRound = (props) => {
  const { playerRound } = props;
  return (
      <td>
        <span>{playerRound.player}, </span>
        <span>Bid: {playerRound.bid}, </span>
        <span>Score: {playerRound.score}</span>
      </td>
  );
};

export default PlayerRound;
