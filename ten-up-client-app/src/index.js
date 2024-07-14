import React from 'react';
import ReactDOM from 'react-dom/client';
import './styles/index.css';
import App from './components/App';

import {
  ApolloProvider,
  ApolloClient,
  InMemoryCache,
} from '@apollo/client';

//import { loadErrorMessages, loadDevMessages } from "@apollo/client/dev";
//loadErrorMessages();

// This points to our gql sercer
const client = new ApolloClient({
  uri: '/graphql',
  cache: new InMemoryCache()
});


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <ApolloProvider client={client}>
    <App />
  </ApolloProvider>,
);
