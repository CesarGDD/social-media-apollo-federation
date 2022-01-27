const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'users', url: 'http://users:8080/query' },
        { name: 'posts', url: 'http://posts:8081/query' },
        { name: 'comments', url: 'http://comments:8082/query' },
        { name: 'likes', url: 'http://likes:8083/query' },
        { name: 'hashtag', url: 'http://hashtag:8084/query' },
        { name: 'followers', url: 'http://followers:8085/query' }
    ],
});

const server = new ApolloServer({
    gateway,
    subscriptions: false,
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});