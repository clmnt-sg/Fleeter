import Fastify from 'fastify'
import database from './routes/database.js'
import user from './routes/user.js'

/**
* @type {import('fastify').FastifyInstance} Instance of Fastify
*/
const fastify = Fastify({
    logger: true
})
fastify.register(database)
fastify.register(user)

fastify.listen({host: "0.0.0.0", port: 3000 }, function (err, address) {
    if (err) {
        fastify.log.error(err)
        process.exit(1)
    }
    // Server is now listening on ${address}
})