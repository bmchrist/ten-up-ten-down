import React from 'react';

import { useMutation, gql } from '@apollo/client';

const UPDATE_BID_MUTATION = gql`
  mutation updateBid($roundId: Int!, $player: String!, $bid: Int!) {
    updateBid(roundId: $roundId, player: $player, bid: $bid){
      id
      bid
    }
  }
`

const PlayerRound = (props) => {
  let input;

  const [updateBid] = useMutation(UPDATE_BID_MUTATION, {});
  // TODO why did this auto update the cache?
  // todo set initial value

  const { playerRound } = props;
  return (
      <td>
        <span>{playerRound.player}, </span>
        <span>Bid: {playerRound.bid}, </span>
        <span>Tricks: {playerRound.tricks}</span>
        <form onSubmit={e => {
            e.preventDefault();
            updateBid({ variables: {
              bid: input.value,
              player: playerRound.player,
              roundId: playerRound.round
            }})
        }}>
          <input
            ref={node => { input = node}}
          />
          <button type="submit">Bid</button>
        </form>
      </td>
  );
};

export default PlayerRound;
