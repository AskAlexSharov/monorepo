import React from 'react';
import ReactDOM from 'react-dom';
import {ApolloProvider as ApolloHooksProvider} from 'react-apollo-hooks';
import ApolloClient from 'apollo-client';
import App from './App';
import {InMemoryCache} from 'apollo-cache-inmemory';
import {split} from 'apollo-link';
import {HttpLink} from 'apollo-link-http';
import {WebSocketLink} from 'apollo-link-ws';
import {getMainDefinition} from 'apollo-utilities';
import * as serviceWorker from './serviceWorker';

const wsLink = new WebSocketLink({
    uri: `ws://localhost:8085/query`,
    options: {
        reconnect: true
    }
});

const httpLink = new HttpLink({uri: 'http://localhost:8085/query'});


// depending on what kind of operation is being sent
interface Definintion {
    kind: string;
    operation?: string;
};

const link = split(
    ({ query }) => {
        const { kind, operation }: Definintion = getMainDefinition(query);
        return kind === 'OperationDefinition' && operation === 'subscription';
    },
    wsLink,
    httpLink,
);

const apolloClient = new ApolloClient({
    link: link,
    cache: new InMemoryCache(),
});

if (module['hot']) {
    module['hot'].accept('./App', () => {
        const NextApp = require('./App').default;
        render(<NextApp />);
    })
}

function render(component: any) {
    ReactDOM.render(
        <ApolloHooksProvider client={apolloClient}>
            {component}
        </ApolloHooksProvider>
        , document.getElementById('root'));
}

render(<App />);


// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
