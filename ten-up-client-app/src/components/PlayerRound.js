import React from 'react';

const PlayerRound = (props) => {
  const { playerRound } = props;
  return (
    <div>
      <span>{playerRound.player}, </span>
      <span>Bid: {playerRound.bid}, </span>
      <span>Score: {playerRound.score}</span>
    </div>
  );
};

export default PlayerRound;
