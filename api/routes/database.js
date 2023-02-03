import fastifyPlugin from 'fastify-plugin'
import fastifyMongo from '@fastify/mongodb'

/**
 * @param {FastifyInstance} fastify
 * @param {Object} options
 */
async function database (fastify, options) {
    fastify.register(fastifyMongo, {
        url: 'mongodb://mongodb:27017/db',
        forceClose : true,
        auth:{
            username: "user",
            password: "password"
        }
    })
}

// Wrapping a plugin function with fastify-plugin exposes the decorators
// and hooks, declared inside the plugin to the parent scope.
export default fastifyPlugin(database)