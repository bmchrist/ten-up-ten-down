import React from 'react';
import ReactDOM from 'react-dom/client';
import './styles/index.css';
import App from './components/App';

import {
  ApolloProvider,
  ApolloClient,
  //createHttpLink,
  InMemoryCache,
  gql
} from '@apollo/client';

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
