import React, { useState } from 'react';
import { useMutation, gql } from '@apollo/client';

const CREATE_ROUND_MUTATION = gql`
  mutation PostMutation(
    $round: String!
  ) {
    post(round: $round) {
      id
      round
    }
  }
`
// todo think about what this should actually do - create all the relevant playerrounds, create just one, etc?
//



const CreateRound = () => {
  const [formState, setFormState] = useState({
    round: ''
  });

  const [createRound] = useMutation(CREATE_ROUND_MUTATION, {
    variables: {
      round: formState.round,
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
            value={formState.round}
            onChange={(e) =>
              setFormState({
                ...formState,
                round: e.target.value
              })
            }
          type="text"
          placeholder="Round"
          />
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default CreateRound;
