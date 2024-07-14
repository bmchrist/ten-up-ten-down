import React from 'react';
import Round from './Round';
import { useQuery, gql } from '@apollo/client';

const ROUNDS_QUERY = gql`
  {
    allRounds {
      id
      sequence
      playerRounds {
        id
        player
        bid
        tricks 
      }
    }
  }
`;

const RoundList = (props) => {
  const { data } = useQuery(ROUNDS_QUERY);

  // right now just wrapping return of a random round in an array, replace with list of rounds next

  return (
    <div className="rounds-container">
      { data && (
        <>
        { data.allRounds.map((round) => (
          <Round key={round.id} round={round} />
        ))}
        </>
      )}
    </div>
  );
};

export default RoundList;
