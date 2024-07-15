import React from 'react';

import { useMutation, gql } from '@apollo/client';

const UPDATE_TRICKS_MUTATION = gql`
  mutation updateTricks($roundId: Int!, $player: String!, $tricks: Int!) {
    updateTricks(roundId: $roundId, player: $player, tricks: $tricks){
      tricks
    }
  }
`
const UPDATE_BID_MUTATION = gql`
  mutation updateBid($roundId: Int!, $player: String!, $bid: Int!) {
    updateBid(roundId: $roundId, player: $player, bid: $bid){
      bid
    }
  }
`

const PlayerRound = (props) => {
  let bidInput;
  let tricksInput;

  const [updateBid] = useMutation(UPDATE_BID_MUTATION, {});
  const [updateTricks] = useMutation(UPDATE_TRICKS_MUTATION, {});
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
              bid: bidInput.value,
              player: playerRound.player,
              roundId: playerRound.round
            }})
        }}>
          <input
            ref={node => { bidInput = node}}
          />
          <button type="submit">Bid</button>
        </form>
        <form onSubmit={e => {
            e.preventDefault();
            updateTricks({ variables: {
              tricks: tricksInput.value,
              player: playerRound.player,
              roundId: playerRound.round
            }})
        }}>
          <input
            ref={node => { tricksInput = node}}
          />
          <button type="submit">Tricks</button>
        </form>
      </td>
  );
};

export default PlayerRound;
