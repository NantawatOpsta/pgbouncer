from locust import HttpUser, task


class UserBehavior(HttpUser):
    @task
    def index(self):
        self.client.get("/")
