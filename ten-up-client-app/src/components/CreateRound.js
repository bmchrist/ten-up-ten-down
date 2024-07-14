import React, { useState } from 'react';
import { useMutation, gql } from '@apollo/client';

// TODO what does it change if I just use id vs get others?
const CREATE_ROUND_MUTATION = gql`
  mutation createRound($numCards: Int!) {
    createRound(numCards: $numCards){
      id
    }
  }
`
// on success do something
const CreateRound = () => {
  let input;

  const [createRound] = useMutation(CREATE_ROUND_MUTATION, {
    // TODO understand this better and the various options
    update(cache, { data: {createRound}}) {
      cache.modify({
        fields: {
          allRounds(existingRounds = []) {
            const newRoundRef = cache.writeFragment({
              data: createRound,
              fragment: gql`
                fragment NewRound on Round {
                  id
                }
              `
            });
            return [...existingRounds, newRoundRef]
          }
        }
      });
    }
  });

  // TODO simplify this form - feels like overkill
  return (
    <>
      <form onSubmit={e =>{
          e.preventDefault();
          createRound({ variables: { numCards: input.value }});
          input.value = '';
      }}>
        <input ref={node => {
          input = node
        }}/>
        <button type="submit">Add round</button>
      </form>
    </>
  );
};

export default CreateRound;
