import React, { useState } from 'react';
import { useMutation, gql } from '@apollo/client';

const CREATE_ROUND_MUTATION = gql`
  mutation createRound($numCards: Int!) {
    createRound(numCards: $numCards){
      round
      player
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
    }
  });

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
