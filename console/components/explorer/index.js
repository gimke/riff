import React from 'react'
import GraphiQL from 'graphiql'
import 'graphiql/graphiql.css'
import './index.pcss'

function graphQLFetcher(graphQLParams) {
    return fetch(window.location.origin + '/api', {
        method: 'post',
        headers: {'Content-Type': 'application/json'},
        credentials: "same-origin",
        body: JSON.stringify(graphQLParams),
    }).then(response => response.json());
}

class Explorer extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return <GraphiQL fetcher={graphQLFetcher}>
        </GraphiQL>
    }
}

export default Explorer