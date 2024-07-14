import React from 'react';

const PlayerRound = (props) => {
  const { playerRound } = props;
  return (
    <div>
      <span>{playerRound.player},</span>
      <span>{playerRound.bid},</span>
      <span>{playerRound.score}</span>
    </div>
  );
};

export default PlayerRound;

