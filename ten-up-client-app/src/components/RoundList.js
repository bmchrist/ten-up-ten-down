import React from 'react';
import Round from './Round';
import { useQuery, gql } from '@apollo/client';

const ROUNDS_QUERY = gql`
  {
    allRounds {
      id
      numCards
      playerRounds {
        id
        player
        round
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
    <table className="rounds-container">
      <thead>
        <tr>
          <td>Round</td>
          <td>TODO Show names here</td>
        </tr>
      </thead>
      <tbody>
        { data && (
          <>
          { data.allRounds.map((round) => (
            <Round key={round.id} round={round} />
          ))}
          </>
        )}
      </tbody>
    </table>
  );
};

export default RoundList;
