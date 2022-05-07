# GO Todo Template
Simplistic template of basic Task Managing Application where you can configure your repository. To make work you need to implement 
`ITodoRepository` interface and assign to repository variable `mainRepository`. See `TodoRepository.go` file for inspiration. 
Created using [Fiber](https://docs.gofiber.io/). Please consider using `TestAppPostman.json` file for inspiration.

## Important configurations
- `PORT` - Port on which application will be listening of requests

## Endpoints
- `POST /api/todo` - To create new TODO
- `GET /api/todo` - List all TODOs created so far
- `GET /api/todo/:id` - Get TODO by defined ID
- `PUT /api/todo/:id` - Updates TODO by defined ID