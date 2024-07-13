import React from 'react';

const Round = (props) => {
  const { round } = props;
  return (
    <div>
      <span>{round.player},</span>
      <span>{round.bid},</span>
      <span>{round.score}</span>
    </div>
  );
};

export default Round;
