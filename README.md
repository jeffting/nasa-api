# nasa-api
Steps to use this api <br>
1. In the root of this directory, run `./nasa-api` to start my api server.
2. In a separate terminal, run `python3 cli.py lastTenDays`. This gets the last 10 days. `Python3` must be installed on computer and the `requests` library needs to be installed too.
3. Or you can run `python3 cli.py specificDay {date}`. Date must be in the format `YYYY-MM-DD`. This gets all the images for a specific day.

Extensibility: I built a full blown api that would be easy to add more api endpoints, handlers, and clients. I used several interfaces which would make unit tests using gomocks easy to implement.