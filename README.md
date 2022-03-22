# GO Mongo
Simplistic template of basic Task Managing Application using [MongoDB](https://docs.mongodb.com/drivers/go/v1.7/) and [Fiber](https://docs.gofiber.io/). Please consider using `TestAppPostman.json` file for inspiration.

## Important configurations
- `PORT` - Port on which application will be listening of requests
- `DB_HOST` - URL to reach Mongo instance, e.g. mongodb://root:example@localhost:27017/?maxPoolSize=20&w=majority
  var

## Endpoints
- `POST /api/todo` - To create new TODO
- `GET /api/todo` - List all TODOs created so far
- `GET /api/todo/:id` - Get TODO by defined ID
- `PUT /api/todo/:id` - Updates TODO by defined ID