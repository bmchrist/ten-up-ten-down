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
  const [formState, setFormState] = useState({
    numCards: ''
  });

  const [createRound] = useMutation(CREATE_ROUND_MUTATION, {
    variables: {
      numCards: formState.numCards,
    },
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
    <div>
      <form onSubmit={(e) =>{
        e.preventDefault();
        createRound();
      }}>
        <div className="flex flex-column mt3">
          <input
            className="mb2"
            value={formState.numCards}
            onChange={(e) =>
              setFormState({
                ...formState,
                numCards: e.target.value
              })
            }
          type="text"
          placeholder="Number of cards for next round"
          />
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default CreateRound;
