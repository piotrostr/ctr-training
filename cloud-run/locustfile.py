from locust import HttpUser, task

class Req(HttpUser):

    @task
    def hello_world(self):
        self.client.get("/")
        self.client.get("/name")
