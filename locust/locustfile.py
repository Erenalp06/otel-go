from locust import HttpUser, task, between

class UserBehavior(HttpUser):
    wait_time = between(1, 2)

    @task(2)
    def get_all_users(self):
        self.client.get("/api/v1/users")

    @task(1)
    def create_user(self):
        user_data = {
            "name": "John Doe",
            "email": "john.doe@example.com",
            "date": "2024-06-09",
            "city": "New York",
            "country": "USA"
        }
        self.client.post("/api/v1/users", json=user_data)

    @task(1)
    def get_user(self):
        self.client.get("/api/v1/users/1")

    @task(1)
    def update_user(self):
        user_data = {
            "id": 1,
            "name": "Jane Doe",
            "email": "jane.doe@example.com",
            "date": "2024-06-10",
            "city": "Los Angeles",
            "country": "USA"
        }
        self.client.put("/api/v1/users/1", json=user_data)

    @task(1)
    def delete_user(self):
        self.client.delete("/api/v1/users/1")
